package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type animal struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

var animals = []animal{
	{ID: 1, Name: "lion", Icon: "🦁"},
	{ID: 2, Name: "cat", Icon: "🐈"},
	{ID: 3, Name: "dog", Icon: "🐕"},
}

func main() {
	http.HandleFunc("/", getAnimals)
	fmt.Println("server starting port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getAnimals(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("access-Control-Allow-Origin", "http://localhost:3000")
	json.NewEncoder(w).Encode(animals)
}
