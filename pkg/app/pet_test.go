package app_test

import (
	"testing"

	"github.com/d-vignesh/getpets/pkg/app"
	"github.com/d-vignesh/getpets/pkg/db"
	"github.com/d-vignesh/getpets/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreatePet(t *testing.T) {
	// mocking the storage with inMemStore to test the service layer
	petDB := db.NewInMem()
	petSvc := app.NewPetSvc(petDB)
	testCases := []struct {
		desc string
		pet  *domain.Pet
		err  error
	}{
		{
			desc: "success",
			pet: &domain.Pet{
				Category: "dog",
				Breed:    "labour",
				Age:      2,
				Gender:   "male",
				Colors:   "black",
				Contact: domain.Contact{
					Owner: "bigshow",
					Phone: "123456789",
				},
				Price: 10000,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := petSvc.Create(tc.pet)
			assert.NoError(t, err)
		})
	}
}
