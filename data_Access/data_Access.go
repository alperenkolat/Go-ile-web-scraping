package data_Access

import (
	"context"
	"example/hello/main_project/catch_err"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

//json dosyasını database yazar
/*func Write_j_database() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	catch_err.Control(err)

	Collection := client.Database("First_Database").Collection("First COllection")

	jsonFile, err := os.Open("C:\\Users\\alper\\OneDrive\\Belgeler\\GO PROJECTS\\GO project-1\\main_project\\read_json\\doipub.json")
	catch_err.Control(err)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string][]string

	json.Unmarshal([]byte(byteValue), &result)

	_, insertErr := Collection.InsertOne(context.TODO(), result)
	catch_err.Control(insertErr)

}*/

func Data_Add_database(registry_no string, doi string, abstract string, keywords []string) {

	data := Academicians{

		RegistryNo: registry_no,
		Articles: []Articles{
			{
				Doi:      doi,
				Abstract: abstract,
				Keywords: keywords,
			},
		},
	}
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	catch_err.Control(err)

	Collection := client.Database("First_Database").Collection("Collection_one")
	_, insertErr := Collection.InsertOne(context.TODO(), data)
	catch_err.Control(insertErr)
}

func Pull_data() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	catch_err.Control(err)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	Collection := client.Database("First_Database").Collection("Collection_one")

	cur, err := Collection.Find(ctx, bson.D{})
	catch_err.Control(err)
	defer cur.Close(ctx)
	var list []Academicians
	for cur.Next(ctx) {
		var result Academicians
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal("Hata : " + err.Error())
		}
		list = append(list, result)
		//burdan
	}
	if err := cur.Err(); err != nil {
		log.Fatal("Hata : " + err.Error())
	}
	fmt.Println(list)

}
