package main

import (
	"fmt"
	"log"
	"net/http"
)

// This is going to be our get request to query the database.
func getServerData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getServerData")
}

func handleRequests() {
	http.HandleFunc("/getServerData", getServerData)
	log.Println("Server up and running on port:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	handleRequests()
}
