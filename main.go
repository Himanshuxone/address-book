// Package used to Parse CSV content for creating an address book
package main

import (
	"app/parsecsv"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	f, err := os.OpenFile("logs", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	parsecsv.CheckError(err)
	//set output of logs to f
	log.SetOutput(f)
	r := mux.NewRouter()
	parsecsv.ProcessCSV()

	// pass a string name to find the contact information using Search handler
	r.HandleFunc("/{firstname}", parsecsv.Search).Methods("GET")
	http.Handle("/", r)
	fmt.Println("Listen to server on port 8080")
	http.ListenAndServe(":8080", nil)
}
