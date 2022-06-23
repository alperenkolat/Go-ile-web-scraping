//kelime çekme paketi
package pull_keywords

//kelime çekme paketi
import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

/*anahtarlar kelimeleri web sitelerinden kazır
 */
func Get_key(url string) (keyword_slice []string) {
	re_strings := [...]string{`keywords" content="(.*?)>`, `(?m)"Keywords":"(.*?)",`, `(?m)<b>Keywords:<\/b> (.*?)<`, `keywords" xml:lang="en" content="(.*?)>`, `Keywords: </span>(.*?)<`, `kwd":\[(.*?)]`, `"keyword"><span>(.*?)<`}

	url = "https://doi.org/" + url
	res, _ := http.Get(url)
	if res.StatusCode != 200 {
		fmt.Println("hata", res.StatusCode)
		return
	}
	bodyBytes, _ := io.ReadAll(res.Body)
	bodyString := string(bodyBytes)
	for i := 0; i < len(re_strings); i++ {
		var re = regexp.MustCompile(re_strings[i])
		var keyword_string = ""
		var count_1 = len(re.FindAllStringSubmatch(bodyString, -1))

		for i, match := range re.FindAllStringSubmatch(bodyString, -1) {
			if count_1 == 2 && i == 1 {
				keyword_string = match[i]
				keyword_slice = strings.Split(keyword_string, ",")

			} else if count_1 < 2 && i == 0 {
				keyword_string = match[1]
				keyword_slice = strings.Split(keyword_string, ",")

			} else if count_1 > 2 {
				keyword_string = match[1]
				keyword_slice = append(keyword_slice, keyword_string)

			}
		}
		if keyword_slice != nil {
			fmt.Println(re_strings[i])
			break

		}
	}

	return
}
