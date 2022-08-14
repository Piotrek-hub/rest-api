package api

import (
	"errors"
	"net/http"
	"rest-api/db"
	"rest-api/db/models"
)

var pf = db.NewPetFactory()

type ResponseContent interface {
	[]*models.Pet | string
}

type Response[T ResponseContent] struct {
	Status  string `json:"status"` // ok, error
	Content T      `json:"content"`
}

func GetPetsHandler(r *http.Request) ([]*models.Pet, error) {
	query := r.URL.Query()

	if name, ok := query["name"]; ok {
		pets, err := pf.GetPetsByName(name[0])
		if err != nil {
			return nil, err
		}
		return pets, nil
	}

	if id, ok := query["id"]; ok {
		pets, err := pf.GetPetById(id[0])
		if err != nil {
			return nil, err
		}
		return []*models.Pet{pets}, nil
	}

	if breed, ok := query["breed"]; ok {
		pets, err := pf.GetPetsByBreed(breed[0])
		if err != nil {
			return nil, err
		}
		return pets, nil
	}

	pets := pf.GetAllPets()
	return pets, nil
}

func AddPetHandler(r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	name := r.FormValue("name")
	breed := r.FormValue("breed")

	if err := pf.NewPet(name, breed); err != nil {
		return err
	}
	return nil
}

func UpdatePetHandler(r *http.Request) error {
	query := r.URL.Query()
	id, ok := query["id"]
	if !ok {
		return errors.New("Pet ID missing")
	}

	if name, ok := query["name"]; ok {
		err := pf.UpdatePetsName(id[0], name[0])
		return err
	}

	if breed, ok := query["breed"]; ok {
		err := pf.UpdatePetsBreed(id[0], breed[0])
		return err
	}

	return errors.New("Please provide more data")
}

func DeletePetHandler(r *http.Request) error {
	query := r.URL.Query()
	id, ok := query["id"]
	if !ok {
		return errors.New("Pet ID missing")
	}

	err := pf.DeletePet(id[0])
	return err
}
