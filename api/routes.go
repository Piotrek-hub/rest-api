package api

import (
	"encoding/json"
	"log"
	"net/http"
	"rest-api/db/models"
)

func petsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		pets, err := GetPetsHandler(r)
		if err != nil {
			Error(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response[[]*models.Pet]{Status: "ok", Content: pets})

	case "POST":
		err := AddPetHandler(r)
		if err != nil {
			Error(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response[string]{Status: "ok", Content: "Pet added successfully"})

	case "PUT":
		err := UpdatePetHandler(r)
		if err != nil {
			Error(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response[string]{Status: "ok", Content: "Pet updated successfully"})

	case "DELETE":
		err := DeletePetHandler(r)
		if err != nil {
			Error(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response[string]{Status: "ok", Content: "Pet deleted successfully"})
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func Serve() {
	http.HandleFunc("/pets/", petsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
