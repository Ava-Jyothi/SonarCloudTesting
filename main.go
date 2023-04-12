package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router instance
	r := CreateRouter()

	// Define the "/hello/{name}" endpoint
	r.HandleFunc("/hello/{name}", helloHandler).Methods("GET")

	// Start the HTTP server
	log.Printf("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func CreateRouter() *mux.Router {
	// Create a new router instance
	r := mux.NewRouter()

	// Define the "/hello/{name}" endpoint
	r.HandleFunc("/hello/{name}", helloHandler).Methods("GET")

	return r
}

type ValidateResponse struct {
	Code    string
	Message string
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the "name" path parameter from the request
	vars := mux.Vars(r)
	name := vars["name"]

	validationResult := validate(name)

	if validationResult != nil {
		// If validation fails, return an HTTP 400 Bad Request error
		http.Error(w, validationResult.Message, http.StatusBadRequest)
		return
	}

	// Write the response back to the client
	fmt.Fprintf(w, "Hi, %s! How you doing?", name)
}

func validate(req string) *ValidateResponse {

	alphabetsRegex := regexp.MustCompile(`^[a-zA-Z]*$`)
	if strings.TrimSpace(req) == "" {
		return &ValidateResponse{Code: "400", Message: "Name cannot be empty"}
	} else if !alphabetsRegex.MatchString(req) {
		return &ValidateResponse{Code: "400", Message: "Name should not contain numbers"}
	} else {
		return nil
	}

}
