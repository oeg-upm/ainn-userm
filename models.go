package main

import (
	"container/list"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type User struct {
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	TokenCode   string    `json:"tokencode"`
	TokenExpiry time.Time `json:"tokenexpiry"`
}

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

func addStructToDB(collection_name string, doc interface{}) bool {
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

func updateWithStructDB(collection_name string, filter *bson.Document, doc interface{}) bool {
	db := getDB()
	if db == nil {
		return false
	}
	collection := db.Collection(collection_name)
	_, err := collection.UpdateOne(context.Background(), filter, doc)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
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
	user.TokenCode = getRandomSequence(10)
	user.TokenExpiry = time.Now().AddDate(0, 0, -1)
	return addStructToDB("user", user)
}

func getFromDB(collection_name string, filter *bson.Document) *list.List {
	db := getDB()
	collection := db.Collection(collection_name)
	results := list.New()
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
	filter := bson.NewDocument(bson.EC.String("username", username))
	results := getFromDB("user", filter)
	var user *User
	user = new(User)
	if results.Len() == 1 {
		element := results.Front().Value.(*bson.Document)
		log.Println(element.Lookup("username").StringValue())
		user.Username = element.Lookup("username").StringValue()
		user.Password = element.Lookup("password").StringValue()
		return user
	} else {
		log.Printf("getUserFromDB> number of users: %d", results.Len())
		return nil
	}
}

func renewUserToken(user *User) {
	user.TokenCode = getRandomSequence(10)
	user.TokenExpiry = time.Now().AddDate(0, 0, 1)
	//add update user
}
