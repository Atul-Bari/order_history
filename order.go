package main

import "sync"

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Order struct {
	Order_id string     `json:"order_id"`
	History  []Location `json:"history"`
}

// key order_id
var OrderMap = make(map[string]*Order)

// lock
var mu sync.RWMutex
