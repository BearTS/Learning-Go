package main

import (
	"Learning-Go/routes"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running....")
	})
	router.HandleFunc("/posts", routes.GetPosts).Methods("GET")
	router.HandleFunc("/posts", routes.AddPost).Methods("POST")
	log.Println("Server is running on port ", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
