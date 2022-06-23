package dowlaodpdf

import (
	"example/hello/main_project/catch_err"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

var (
	fileName    string
	fullURLFile string
)

/*pdfleri indirir*/
func Dow_pdf() {

	fullURLFile = ""

	if !strings.Contains(fullURLFile, ".pdf") {
		res, _ := http.Get(fullURLFile)
		fmt.Println("Getirilfi")
		if res.StatusCode != 200 {
			fmt.Println("Error:", res.StatusCode)
			return
		}
		bodyBytes, _ := io.ReadAll(res.Body)
		bodyString := string(bodyBytes)
		var re = regexp.MustCompile(`http(.*?).pdf`)

		for i, match := range re.FindAllStringSubmatch(bodyString, -1) {
			fmt.Println(match, "found at index", i)
			fmt.Println(match[0])
			fullURLFile = match[0]
		}

	}
	fileURL, err := url.Parse(fullURLFile)
	catch_err.Control(err)
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName = segments[len(segments)-1]
	file, err := os.Create(fileName)
	catch_err.Control(err)
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(fullURLFile)
	catch_err.Control(err)
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)
	catch_err.Control(err)

	file.Close()
	fmt.Printf("Downloaded a file %s with size %d \n", fileName, size)
	if !strings.Contains(fileName, ".pdf") {
		//err := os.Remove(fileName)
		//control(err)

		fmt.Println("pdf değil silindi")

	}

}
