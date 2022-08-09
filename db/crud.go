package db

import (
	"errors"
	. "rest-api/db/models"
)

type PetFactory struct {
	pets []*Pet
}

func NewPetFactory() *PetFactory {
	return &PetFactory{}
}

// Create
func (this *PetFactory) NewPet(name, breed string) {
	newPet := &Pet{
		ID:    uint(len(this.pets)),
		Name:  name,
		Breed: breed,
	}

	this.pets = append(this.pets, newPet)
}

// Read

func (this *PetFactory) GetAllPets() []*Pet {
	return this.pets
}

func (this *PetFactory) GetPetById(id uint) (*Pet, error) {
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

// Update
func (this *PetFactory) UpdatePetsName(id uint, newName string) error {
	pet, err := this.GetPetById(id)
	if err != nil {
		return err
	}

	pet.Name = newName

	return nil
}

func (this *PetFactory) UpdatePetsBreed(id uint, newBreed string) error {
	pet, err := this.GetPetById(id)
	if err != nil {
		return err
	}

	pet.Breed = newBreed

	return nil
}
