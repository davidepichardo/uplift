package main 

import (
	// "database/sql"
	"os"
	"fmt"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// code to create a new user, sending a POST request

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	// code to get all users and return as JSON
}

func getUser(w http.ResponseWriter, r *http.Request) {
	// code to get a specific user by id and return as JSON
	
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	// code to update a specific user by id
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	// code to delete a specific user by id
}