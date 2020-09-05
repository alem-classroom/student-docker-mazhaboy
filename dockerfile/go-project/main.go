package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	PORT = os.Getenv("PORT")
)

func main() {
	if PORT == "" {
		log.Printf("port not specified!")
		return
	}
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		resp := map[string]string{
			"Hello": os.Getenv("APP_TYPE"),
		}
		body, _ := json.Marshal(resp)
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}).Methods("GET")

	router.HandleFunc("/items/{item_id}", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		resp := map[string]string{
			"item_id": vars["item_id"],
			"page":    "items",
		}
		body, _ := json.Marshal(resp)
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}).Methods("GET")

	log.Printf("running on port: %s\n", PORT)
	http.ListenAndServe(":"+PORT, router)
}
