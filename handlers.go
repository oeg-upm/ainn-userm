package main

import (
	"fmt"
	"log"
	"net/http"
)

func loginGET(w http.ResponseWriter, r *http.Request) {
	//testDB()
	fmt.Fprintf(w, "I am the login ")
}

func loginPOST(w http.ResponseWriter, r *http.Request) {
	//decoder := json.NewDecoder(r.Body)
	//var errorMsg string
	var ok bool
	var username string
	var password string
	var user User
	log.Println("loginPOST")
	r.ParseForm()
	if _, ok = r.Form["username"]; ok {
		username = r.Form["username"][0]
		log.Printf("username is %s", username)
		if _, ok = r.Form["password"]; ok {
			password = r.Form["password"][0]
			log.Printf("password passed")
			user = User{Username: username, Password: password}
			log.Printf("user obj is created %s", user.Username)
		} else {
			log.Println("password is not passed")
		}
	} else {
		log.Println("username is not passed")
	}

	fmt.Fprintf(w, "login post ")
}

func homeGET(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "usermanagement home page")
}
