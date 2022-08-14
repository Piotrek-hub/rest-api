package db

import (
	"errors"
	. "rest-api/db/models"
	"strconv"
)

type PetFactory struct {
	pets []*Pet
}

func NewPetFactory() *PetFactory {
	return &PetFactory{}
}

func (this *PetFactory) NewPet(name, breed string) error {
	if name == "" || breed == "" {
		return errors.New("Provide all information about pet")
	}
	newPet := &Pet{
		ID:    strconv.Itoa(len(this.pets)),
		Name:  name,
		Breed: breed,
	}

	this.pets = append(this.pets, newPet)

	return nil
}

func (this *PetFactory) GetAllPets() []*Pet {
	return this.pets
}

func (this *PetFactory) GetPetById(id string) (*Pet, error) {
	for _, pet := range this.pets {
		if pet.ID == id {
			return pet, nil
		}
	}

	return nil, errors.New("Pet not found, probably bad ID")
}

func (this *PetFactory) GetPetByName(name string) (*Pet, error) {
	for _, pet := range this.pets {
		if pet.Name == name {
			return pet, nil
		}
	}

	return nil, errors.New("Pet not found, probably bad name")
}

func (this *PetFactory) GetPetsByBreed(breed string) ([]*Pet, error) {
	var pets []*Pet
	for _, pet := range this.pets {
		if pet.Breed == breed {
			pets = append(pets, pet)
		}
	}

	if len(pets) == 0 {
		return nil, errors.New("No pets with this breed")
	} else {
		return pets, nil
	}
}

// Update
func (this *PetFactory) UpdatePetsName(id string, newName string) error {
	pet, err := this.GetPetById(id)
	if err != nil {
		return err
	}

	pet.Name = newName

	return nil
}

func (this *PetFactory) UpdatePetsBreed(id string, newBreed string) error {
	pet, err := this.GetPetById(id)
	if err != nil {
		return err
	}

	pet.Breed = newBreed

	return nil
}
