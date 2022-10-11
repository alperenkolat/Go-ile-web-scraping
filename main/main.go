package main

import (
	"example/hello/main_project/pull_abstract"
	"example/hello/main_project/pull_keywords"
	"example/hello/main_project/read_json"
	"fmt"
)

type Academicians struct {
	RegistryNo string     `json:"registry no"`
	Articles   []Articles `json:"Articles"`
}

type Articles struct {
	Doi      string   `json:"doi"`
	Abstract string   `json:"abstract"`
	Keywords []string `json:"keywords"`
}

func main() {
	res := read_json.Read_f()
	keywords_count := 0
	abstract_count := 0
	count_index := 0
	for registry, doi := range res {
		for i := 0; i < len(doi); i++ {

			abstract_slice := pull_abstract.Get_abstract(doi[i])
			/*	if abstract_slice == nil {
				log.Fatal("abstract is nil")

			}*/

			keywords := pull_keywords.Get_key(doi[i])
			/*	if keywords == nil {
				log.Fatal("keywords is nil")

			}*/

			//data_Access.Data_Add_database("", "", abstract_slice[0], keywords)
			if keywords == nil {

				keywords_count++

			}
			if abstract_slice == nil {
				abstract_count++

			}

			fmt.Println(registry, doi[i], i)
			fmt.Println(abstract_slice, keywords)
			fmt.Println("----------------------------")
			count_index++
			fmt.Println(count_index)
			fmt.Println(keywords_count, abstract_count)

		}
	}
}
