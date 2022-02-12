package mocks

import (
	"github.com/d-vignesh/getpets/pkg/domain"
	"github.com/google/uuid"
)

// mock for pet service
type PetSvc struct {
	GetPetResp   domain.Pet
	GetPetErr    error
	ListPetResp  []*domain.Pet
	ListPetErr   error
	CreatePetErr error
	DeletePetErr error
}

func (ps PetSvc) Get(uuid.UUID) (*domain.Pet, error) {
	return &ps.GetPetResp, ps.GetPetErr
}

func (ps PetSvc) List(string) ([]*domain.Pet, error) {
	return ps.ListPetResp, ps.ListPetErr
}

func (ps PetSvc) Create(*domain.Pet) error {
	return ps.CreatePetErr
}

func (ps PetSvc) Delete(uuid.UUID) error {
	return ps.DeletePetErr
}
