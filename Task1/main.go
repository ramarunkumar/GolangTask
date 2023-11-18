package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	models "task/Models"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var event models.Event
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&event)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	fmt.Println("event:", event)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", handleRequest)
	port := 8080
	fmt.Printf("Server listening on :%d\n", port)
	err := http.ListenAndServe("8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
