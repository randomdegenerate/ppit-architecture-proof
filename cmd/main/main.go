package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ppit-architecture-proof/mongodb"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /item/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Println(id)
		w.Write(mongodb.FindById(id))
	})

	router.HandleFunc("PUT /item/add", func(w http.ResponseWriter, r *http.Request) {
		
		var item mongodb.Item
		err := json.NewDecoder(r.Body).Decode(&item)
		if err != nil {
			w.WriteHeader(400)
			fmt.Println(err)
		}
		fmt.Println(item)

		err = mongodb.InsertItem(item)
		if err != nil {
			w.WriteHeader(400)
			fmt.Println(err)
		} else {
			w.Write([]byte("item inserted sucessfully"))
		}
	})
	mongodb.ConnectDatabase()
	server := http.Server{
		Addr: ":8080",
		Handler: router,
	}

	log.Println("Starting server on port :8080")
	server.ListenAndServe()
}
