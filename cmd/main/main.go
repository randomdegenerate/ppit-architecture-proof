package main

import (
	"net/http"
	"log"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/item/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("this is a response for item get"))
	})

	server := http.Server{
		Addr: ":8080",
		Handler: router,
	}

	log.Println("Starting server on port :8080")
	server.ListenAndServe()
}
