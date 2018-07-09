package main

import (
	"container/list"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
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
	return true
}

func addUserToDB(user User) bool {
	user.Password = encrypt(user.Password)
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

func getFromDB(collection_name string, filter *bson.Document) *list.List {
	db := getDB()
	collection := db.Collection(collection_name)
	results := list.New()
	//filter := bson.NewDocument(bson.EC.String("username", "aabcsd"))
	cur, err := collection.Find(context.Background(), filter)
	if err == nil {
		defer cur.Close(context.Background())
		for cur.Next(context.Background()) {
			elem := bson.NewDocument()
			err := cur.Decode(elem)
			if err != nil {
				log.Fatal(err)
			} else {
				results.PushBack(elem)
			}
		}
	} else {
		log.Fatal(err)
	}
	return results
}

func getUserFromDB(username string) *User {
	log.Println("get user from db")
	// db := getDB()
	// collection := db.Collection("user")
	// results := list.New()
	filter := bson.NewDocument(bson.EC.String("username", username))
	results := getFromDB("user", filter)
	var user *User
	user = new(User)
	if results.Len() == 1 {
		element := results.Front().Value.(*bson.Document)
		log.Println(element.Lookup("username").StringValue())
		// ele := bson.Document(*element)
		// log.Println(reflect.TypeOf(ele))
		user.Username = element.Lookup("username").StringValue()
		user.Password = element.Lookup("password").StringValue()

		// user.Username = element.Value.Lookup("username").StringValue()
		// user.Password = element.Value.Lookup("password").StringValue()
		return user
	} else {
		log.Printf("getUserFromDB> number of users: %d", results.Len())
		return nil
	}
	// cur, err := collection.Find(context.Background(), filter)
	// var data_m map[string]interface{}
	// if err == nil {
	// 	log.Println("found: ")
	// 	defer cur.Close(context.Background())
	// 	for cur.Next(context.Background()) {
	// 		elem := bson.NewDocument()
	// 		err := cur.Decode(elem)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		} else {
	// 			log.Println("element: ")
	// 			log.Println(elem)
	// 			log.Println(reflect.TypeOf(elem))
	// 			log.Println("element username: ")
	// 			log.Println(elem.Lookup("username").StringValue())
	// 			log.Println("printed username")
	// 			results.PushBack(elem)
	// 			b, err2 := elem.MarshalBSON()
	// 			if err2 == nil {
	// 				log.Println(b)
	// 				log.Println("ele marshaled ok")
	// 				err = bson.Unmarshal(b, &data_m)
	// 				if err == nil {
	// 					log.Println("b unmarshaled")
	// 					log.Println(data_m)
	// 				} else {
	// 					log.Println("error in unmarshaled")
	// 					log.Println(err)
	// 				}
	// 			} else {
	// 				log.Println("err2")
	// 				log.Println(err2)
	// 			}
	// 		}
	// 	}
	// 	log.Println("list: ")
	// 	//log.Println(results.Front())
	//
	// 	// for temp := results.Front(); temp != nil; temp = temp.Next() {
	// 	// 	err = temp.Value.UnmarshalBSON(&data_m)
	// 	// 	if err == nil {
	// 	// 		log.Println("data_m: ")
	// 	// 		log.Println(data_m)
	// 	// 	} else {
	// 	// 		log.Println("error unmarshalling")
	// 	// 	}
	// 	//
	// 	// 	//log.Println(temp.Value.Lookup())
	// 	// 	//log.Println(reflect.TypeOf(temp.Value))
	// 	// }
	// 	log.Println("end the for")
	// } else {
	// 	log.Println("error: ")
	// 	log.Println(err)
	// }

}
