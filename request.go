package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteHistory(w http.ResponseWriter, r *http.Request) {
	log.Println("In DeleteHistory")
	vars := mux.Vars(r)
	key := vars["id"]

	mu.Lock()
	defer mu.Unlock()
	for i, v := range OrderMap {
		if key == i {
			v.History = nil
		}
	}

}

func AppendHistory(w http.ResponseWriter, r *http.Request) {
	log.Println("In AppendHistory")
	vars := mux.Vars(r)
	key := vars["id"]

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Err: ", err)
		return
	}

	var location Location
	err = json.Unmarshal(reqBody, &location)
	if err != nil {
		log.Println("Location Unmarshal err: ", err)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	for i, v := range OrderMap {
		if key == i {
			v.History = append(v.History, location)
		}
	}

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
			var res Order
			for i := len(v.History) - 1; i >= 0; i-- {
				res.History = append(res.History, v.History[i])
			}
			json.NewEncoder(w).Encode(res)
		}
	} else {
		intMax, err := strconv.Atoi(max)
		if err != nil || intMax == 0 {
			log.Println("Max is invalid")
			return
		}
		var res Order
		if v, ok := OrderMap[key]; ok {
			if len(v.History) > intMax {
				res.Order_id = v.Order_id
				cnt := 0
				for i := len(v.History) - 1; i >= 0; i-- {
					if cnt >= intMax {
						break
					}
					res.History = append(res.History, v.History[i])
					cnt += 1
				}
				json.NewEncoder(w).Encode(res)
			} else {
				for i := len(v.History) - 1; i >= 0; i-- {
					res.History = append(res.History, v.History[i])
				}
				json.NewEncoder(w).Encode(res)
			}

		}
	}

}
