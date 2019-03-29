package main

import (
	"cockroach-test/queries"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/jabatan", queries.GetAllJabatan).Methods("GET")
	router.HandleFunc("/jabatan/update/{id}", queries.UpdateJabatan).Methods("POST", "GET")

	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}
