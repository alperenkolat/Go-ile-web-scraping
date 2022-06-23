package read_json

//fonksiyona büyük harfle yaz
import (
	"encoding/json"
	"example/hello/main_project/catch_err"
	"fmt"
	"io/ioutil"
	"os"
)

/*
json dosyasını okur*/
func Read_f() (res map[string][]string) {

	fileContent, err := os.Open("C:\\Users\\alper\\OneDrive\\Belgeler\\GO PROJECTS\\GO project-1\\main_project\\read_json\\doipub.json")

	catch_err.Control(err)

	fmt.Println("The File is opened successfully...")

	byteResult, _ := ioutil.ReadAll(fileContent)

	json.Unmarshal([]byte(byteResult), &res)
	defer fileContent.Close()
	return res

}

/*
jsonFile, err := os.Open("C:\\Users\\alper\\OneDrive\\Belgeler\\GO PROJECTS\\GO project-1\\main_project\\Academicians.json")

if err != nil {
	fmt.Println(err)
}
fmt.Println("Successfully Academicians.json")

//var sdsdg []string
defer jsonFile.Close()
byteValue, _ := ioutil.ReadAll(jsonFile)
var academicians Academicians
json.Unmarshal(byteValue, &academicians)
for i := 0; i < 1; i++ {
	fmt.Println("akademisyen sicil no: " + academicians.RegistryNo)
	fmt.Println("akademisyen Abstract: " + academicians.Articles[0].Abstract)
	fmt.Println("akademisyen Doi: " + academicians.Articles[0].Doi)
	for i := 0; i < 2; i++ {

		fmt.Println("akademisyen anahtar kelimeler " + academicians.Articles[0].Keywords[i])
	}
}*/
