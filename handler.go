package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ErrorMsg struct {
	Msg string `json:"msg"`
}

func DeleteHistory(w http.ResponseWriter, r *http.Request) {
	log.Println("In DeleteHistory")
	vars := mux.Vars(r)
	key := vars["id"]

	mu.Lock()
	defer mu.Unlock()
	if v, ok := OrderMap[key]; ok {
		v.History = []Location{}
		json.NewEncoder(w).Encode(v)
	} else {
		json.NewEncoder(w).Encode(ErrorMsg{Msg: "Invalid order_id"})
	}
}

func AppendHistory(w http.ResponseWriter, r *http.Request) {
	log.Println("In AppendHistory")
	vars := mux.Vars(r)
	key := vars["id"]

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("i/o Read err: ", err)
		json.NewEncoder(w).Encode(ErrorMsg{Msg: "i/o read error"})
		return
	}

	var location Location
	err = json.Unmarshal(reqBody, &location)
	if err != nil {
		log.Println("Location Unmarshal err: ", err)
		json.NewEncoder(w).Encode(ErrorMsg{Msg: "Invalid input"})
		return
	}
	mu.Lock()
	defer mu.Unlock()
	if v, ok := OrderMap[key]; ok {
		v.History = append([]Location{location}, v.History...)
	} else {
		OrderMap[key] = &Order{Order_id: key, History: []Location{location}}
	}
	json.NewEncoder(w).Encode(location)

}

func GetHistory(w http.ResponseWriter, r *http.Request) {
	log.Println("In GetHistory")
	vars := mux.Vars(r)
	key := vars["id"]

	max := r.URL.Query().Get("max")
	mu.Lock()
	defer mu.Unlock()
	if max == "" {
		if v, ok := OrderMap[key]; ok {
			json.NewEncoder(w).Encode(v)
		} else {
			json.NewEncoder(w).Encode(ErrorMsg{Msg: "Invalid order_id"})
		}
	} else {
		intMax, err := strconv.Atoi(max)
		if err != nil || intMax == 0 {
			log.Println("Max is invalid")
			json.NewEncoder(w).Encode(ErrorMsg{Msg: "Invalid input"})
			return
		}
		var res Order
		if v, ok := OrderMap[key]; ok {
			if len(v.History) > intMax {
				res.Order_id = v.Order_id
				res.History = v.History[:intMax]
				json.NewEncoder(w).Encode(res)
			} else {
				json.NewEncoder(w).Encode(v)
			}
		} else {
			json.NewEncoder(w).Encode(ErrorMsg{Msg: "Invalid order_id"})
		}
	}

}
