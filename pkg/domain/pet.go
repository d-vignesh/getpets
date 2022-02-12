package domain

import (
	"github.com/google/uuid"
)

type Pet struct {
	ID       uuid.UUID `json:"id" bson:"id,omitempty"`
	Category string    `json:"category" bson:"category,omitempty"`
	Breed    string    `json:"breed" bson:"breed,omitempty"`
	Age      int       `json:"age" bson:"age,omitempty"`
	Gender   string    `json:"gender" bson:"gender,omitempty"`
	Colors   string    `json:"colors" bson:"colors,omitempty"`
	Contact  Contact   `json:"contact" bson:"contact,omitempty"`
	Price    float64   `json:"price" bson:"price,omitempty"`
}

type Contact struct {
	Owner string `json:"owner" bson:"owner,omitempty"`
	Phone string `json:"phone" bson:"phone,omitempty"`
	City  string `json:"city" bson:"city,omitempty"`
	State string `json:"state" bson:"state,omitempty"`
}

type PetSvc interface {
	Get(id uuid.UUID) (*Pet, error)
	List(category string) ([]*Pet, error)
	Create(p *Pet) error
	Delete(id uuid.UUID) error
}

type PetDB interface {
	Get(id uuid.UUID) (*Pet, error)
	List(category string) ([]*Pet, error)
	Create(p *Pet) error
	Delete(id uuid.UUID) error
}
