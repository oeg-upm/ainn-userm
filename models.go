package main

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	//Token    string `json:"token"`
}

// func testDB() {
// 	client, err := mongo.NewClient("mongodb://localhost:27017")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = client.Connect(context.TODO())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	collection := client.Database("ainnauth").Collection("test")
// 	res, err := collection.InsertOne(context.Background(), map[string]string{"hello": "world"})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	id := res.InsertedID
// 	fmt.Println(id)
// }

func getDB() *mongo.Database {
	client, err := mongo.NewClient("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return client.Database("ainnauth")
}

func addToDB(collection_name string, doc map[string]interface{}) bool {
	db := getDB()
	if db == nil {
		return false
	}
	collection := db.Collection(collection_name)
	res, err := collection.InsertOne(context.Background(), doc)
	if err != nil {
		log.Fatal(err)
		return false
	}
	id := res.InsertedID
	fmt.Println(id)
	h := md5.New()
	io.WriteString(h, "The fog is getting thicker!")
	//io.WriteString(h, "And Leon's getting laaarger!")
	fmt.Printf("%x", h.Sum(nil))
	return true
}

func addUserToDB(user User) bool {
	var data_m map[string]interface{}
	log.Println("user: ")
	log.Println(user)
	b, err := json.Marshal(user)
	if err == nil {
		log.Println("Marshal success")
		log.Println(b)
		err = json.Unmarshal(b, &data_m)
		//data = json.Unmarshal
		//json.Unmarshal(data, v)
		if err == nil {
			log.Println("unMarshal success")
			log.Println(data_m)
			return addToDB("user", data_m)
		} else {
			log.Println("unMarshal fail")
			log.Println(err)
		}

	} else {
		log.Println("Marshall error")
		log.Println(err)
	}
	return false
}
