package main

import (
	"cockroach-test/config"
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
	conf := config.SetConfig()
	fmt.Println("Connected to port " + conf.HttpPort)
	log.Fatal(http.ListenAndServe(":"+conf.HttpPort, router))
}
