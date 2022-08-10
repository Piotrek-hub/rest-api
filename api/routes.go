package api

import (
	"encoding/json"
	"log"
	"net/http"
	db "rest-api/db"
	models "rest-api/db/models"
)

type GetPetsResponse struct {
	Status string
	Pets   []*models.Pet
}

var petFactory = db.NewPetFactory()

func petsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		// Get pets
		pets := petFactory.GetAllPets()
		resp := GetPetsResponse{
			Status: "ok",
			Pets:   pets,
		}

		json.NewEncoder(w).Encode(resp)
	case "POST":
		r.ParseForm()

		name := r.FormValue("name")
		breed := r.FormValue("breed")

		if err := petFactory.NewPet(name, breed); err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte("Pet created succesfully"))
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
