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

	router.HandleFunc("GET /item/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		fmt.Println(name)
		w.Write(mongodb.GetByName(name))
	})
		
	router.HandleFunc("PUT /item/add", func(w http.ResponseWriter, r *http.Request) {
		
		var item mongodb.Item
		err := json.NewDecoder(r.Body).Decode(&item)
		if err != nil {
			//error declaration stuff
			fmt.Printf("decoding error:%v",err)
			return
		}
		
		fmt.Printf("item: \n%+v\n", item)


		err = mongodb.InsertItem(item)
		if err != nil {
			//error declaration
			fmt.Println("insertion error")
			return
		}	
	})
	 
	router.HandleFunc("GET /items", func(w http.ResponseWriter, r *http.Request) {
		jsonData,err := mongodb.GetItems()
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.Write(jsonData)
	})
	

	mongodb.ConnectDatabase()
	server := http.Server{
		Addr: ":8080",
		Handler: router,
	}

	log.Println("Starting server on port :8080")
	server.ListenAndServe()
}
