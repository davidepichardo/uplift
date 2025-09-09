package main 

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request method:", r.Method)
	fmt.Println("Request path:", r.URL.Path)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) 
	w.Write([]byte(`{"message": "hello"}`))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Request Method:", r.Method)
    fmt.Println("Request Path:", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "pong"}`))
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}