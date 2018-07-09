package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func registerPOST(w http.ResponseWriter, r *http.Request) {
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
	added := addUserToDB(user)
	log.Println(added)
	fmt.Fprintf(w, "register post")
}

func loginPOST(w http.ResponseWriter, r *http.Request) {
	var ok bool
	var username string
	var password string
	responseJ := make(map[string]string)
	r.ParseForm()
	log.Printf("loginPOST> in Login Post")
	if _, ok = r.Form["username"]; ok {
		username = r.Form["username"][0]
		log.Printf("loginPOST> username is %s", username)
		if _, ok = r.Form["password"]; ok {
			password = r.Form["password"][0]
			log.Printf("loginPOST> password passed")
			password = encrypt(password)
			user := getUserFromDB(username)
			if user != nil {
				if user.Password == password {
					log.Println("loginPOST> password is correct")
					responseJ["message"] = "logged in successfully"
					w.WriteHeader(http.StatusCreated)
				} else {
					log.Println("loginPOST> password does not match")
					responseJ["error"] = "password does not match"
					w.WriteHeader(http.StatusUnauthorized)
				}
			} else {
				log.Println("loginPOST> username does not exists")
				responseJ["error"] = "username does not exists"
				w.WriteHeader(http.StatusUnauthorized)
			}
		} else {
			log.Println("loginPOST> password is not passed")
			responseJ["error"] = "password is not passed"
			w.WriteHeader(http.StatusUnauthorized)
		}
	} else {
		log.Println("loginPOST> username is not passed")
		responseJ["error"] = "username is not passed"
		w.WriteHeader(http.StatusUnauthorized)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseJ)
	//fmt.Fprintf(w, "login post ")
}

func homeGET(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "usermanagement home page")
}
