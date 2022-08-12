package api

import (
	"encoding/json"
	"log"
	"net/http"
	db "rest-api/db"
	"rest-api/db/models"
)

type ResponseContent interface {
	*models.Pet | []*models.Pet | string
}

type Response[T ResponseContent] struct {
	Status  string `json:"status"` // ok, error
	Content T      `json:"content"`
}

var petFactory = db.NewPetFactory()

func petsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		// Get pets
		pets := petFactory.GetAllPets()
		resp := Response[[]*models.Pet]{
			Status:  "ok",
			Content: pets,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	case "POST":
		r.ParseForm()

		name := r.FormValue("name")
		breed := r.FormValue("breed")

		if err := petFactory.NewPet(name, breed); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response[string]{Status: "error", Content: err.Error()})
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(Response[string]{Status: "ok", Content: "pet added successfully"})
		}

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
