package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func HandelRequest() {
	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/location/{id}", DeleteHistory).Methods("DELETE")
	myRouter.HandleFunc("/location/{id}", AppendHistory).Methods("PUT")
	myRouter.HandleFunc("/location/{id}", GetHistory).Methods("GET")

	url := os.Getenv("HISTORY_SERVER_LISTEN_ADDR") + ":8080"
	log.Fatal(http.ListenAndServe(url, myRouter))

}

func main() {
	fmt.Println("Main")
	mu.Lock()
	OrderMap["abc123"] = &Order{Order_id: "abc123", History: []Location{{Lat: 12.34, Lng: 56.75}, {Lat: 13.34, Lng: 78.74}}}
	mu.Unlock()
	HandelRequest()
}
