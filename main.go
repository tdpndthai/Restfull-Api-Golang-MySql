package main

import (
	"fmt"
	"net/http"
	"restfull-api/apis/proc_api"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/proc/findall", proc_api.FindAll).Methods("GET")
	router.HandleFunc("/api/proc/search/{keyword}", proc_api.Search).Methods(http.MethodGet)
	router.HandleFunc("/api/proc/searchprices/{min}-{max}", proc_api.SearchPrices).Methods(http.MethodGet)
	router.HandleFunc("/api/proc/create", proc_api.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/proc/update", proc_api.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/proc/delete/{id}", proc_api.Delete).Methods(http.MethodDelete)

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	} else {

	}
}
