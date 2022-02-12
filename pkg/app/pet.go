package app

import (
	"github.com/d-vignesh/getpets/pkg/domain"
	"github.com/google/uuid"
)

// petSvc implements domain.PetSvc
type petSvc struct {
	DB domain.PetDB
}

func NewPetSvc(db domain.PetDB) domain.PetSvc {
	return petSvc{
		DB: db,
	}
}

func (ps petSvc) Get(id uuid.UUID) (*domain.Pet, error) {
	return ps.DB.Get(id)

}

func (ps petSvc) List(category string) ([]*domain.Pet, error) {
	return ps.DB.List(category)
}

func (ps petSvc) Create(pet *domain.Pet) error {
	pet.ID = uuid.New()
	return ps.DB.Create(pet)
}

func (ps petSvc) Delete(id uuid.UUID) error {
	return ps.DB.Delete(id)
}
