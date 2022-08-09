package api

import (
	"log"
	"net/http"

	"rest-api/db"
)

func petsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		// Get pets
		pets := db.GetAllPets()
		log.Println(pets)
	case "POST":
		log.Println("POST")
	case "UPDATE":
		log.Println("UPDATE")
	case "DELETE":
		log.Println("DELETE")
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func Serve() {
	http.HandleFunc("/pets/", petsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
