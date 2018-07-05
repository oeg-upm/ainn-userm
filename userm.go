package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there!")
}

func main() {
	routers()
	//http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8001", nil))
}

// func main() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", hello)
// 	// r.HandleFunc("/products", ProductsHandler)
// 	// r.HandleFunc("/articles", ArticlesHandler)
// 	http.Handle("/", r)
//
// 	log.Fatal(http.ListenAndServe(":8001", nil))
// }
