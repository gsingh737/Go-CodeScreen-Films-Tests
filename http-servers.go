package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Film struct {
	Name         string  `json:'name'`
	Length       int32   `json:'length'`
	Rating       float64 `json:'rating'`
	ReleaseDate  string  `json:'releaseDate'`
	DirectorName string  `json:'directorName'`
}

func main() {
	http.HandleFunc("/films", getData)
	port := 8080 // Change this to the desired port number
	fmt.Printf("Server is running on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func getData(w http.ResponseWriter, r *http.Request) {
	// Open and read the JSON file
	file, err := os.Open("data.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Decode the JSON data from the file
	var films []Film
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&films)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")

	// Encode and send the JSON response
	if err := json.NewEncoder(w).Encode(films); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
