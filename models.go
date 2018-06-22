package main

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	//Token    string `json:"token"`
}

// func testDB() {
//
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
// 	//fmt.Println(id)
// }
