package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func routers() {
	mrouter := mux.NewRouter() //.StrictSlash(true)
	mrouter.HandleFunc("/login", loginPOST).Methods("POST")
	//mrouter.HandleFunc("/login", loginGET).Methods("GET")
	mrouter.HandleFunc("/", homeGET)
	http.Handle("/", mrouter)
}

// func routers() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", loginGET)
// 	r.HandleFunc("/products", loginGET)
// 	r.HandleFunc("/articles", loginGET)
// 	http.Handle("/", r)
// }
