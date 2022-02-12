package db

import (
	"github.com/d-vignesh/getpets/pkg/domain"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// inMemStore implements domain.PetDB with an memory storage
type inMemStore struct {
	Pets map[uuid.UUID]*domain.Pet
}

func NewInMem() domain.PetDB {
	return inMemStore{
		Pets: make(map[uuid.UUID]*domain.Pet),
	}
}

func (im inMemStore) Get(id uuid.UUID) (*domain.Pet, error) {
	pet, exists := im.Pets[id]
	if !exists {
		return nil, errors.Errorf("no pet found with id: %s", id)
	}
	return pet, nil
}

func (im inMemStore) List(category string) ([]*domain.Pet, error) {
	pets := []*domain.Pet{}
	for _, p := range im.Pets {
		if p.Category == category {
			pets = append(pets, p)
		}
	}
	return pets, nil
}

func (im inMemStore) Create(pet *domain.Pet) error {
	im.Pets[pet.ID] = pet
	return nil
}

func (im inMemStore) Delete(id uuid.UUID) error {
	delete(im.Pets, id)
	return nil
}
