package pull_abstract

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

/*
abstract'ı web sitelerinden kazır
*/
func Get_abstract(url string) (abstract_slice []string) {
	re_strings := [...]string{`description" content="(.*?)>`, `Description" content="(.*?)>`, `abstract" content="(.*?)>`, `Description" xml:lang="en" content="(.*?)>`, `description content="(.*?)>`, `abstractInFull">(.*?)<\/p`}
	url = "https://doi.org/" + url
	res, _ := http.Get(url) //!!
	if res.StatusCode != 200 {
		fmt.Println("hata", res.StatusCode)
		return
	}
	bodyBytes, _ := io.ReadAll(res.Body)
	bodyString := string(bodyBytes)

	for i := 0; i < 6; i++ {

		var re = regexp.MustCompile(re_strings[i])
		var abstract_string = ""

		var count_1 = len(re.FindAllStringSubmatch(bodyString, -1))

		for i, match := range re.FindAllStringSubmatch(bodyString, -1) {
			if count_1 == 2 && i == 1 {
				abstract_string = match[i]
				abstract_slice = strings.Split(abstract_string, ",")

			} else if count_1 < 2 && i == 0 {
				abstract_string = match[1]
				if len(abstract_string) > 40 {

					abstract_slice = append(abstract_slice, abstract_string)
				}

			} else if count_1 > 2 {
				abstract_string = match[1]
				if len(abstract_string) > 40 {

					abstract_slice = append(abstract_slice, abstract_string)
				}

			}
		}
		if abstract_slice != nil {
			fmt.Println(re_strings[i])
			break

		}
	}
	return
}
