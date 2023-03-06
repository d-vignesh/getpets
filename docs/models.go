package docs

import "github.com/d-vignesh/getpets/pkg/domain"

// model for PetID param
// swagger:parameters GetPet DeletePet
type PetIDParam struct {
	// The id of the pet
	//
	// in:path
	// required:true
	ID int `json:"id"`
}

// model for error response
// swagger:response ErrorResponse
type ErrorResponse struct {
	// in:body
	Body struct {
		Msg string `json:"message"`
	} `json:"body"`
}

// model for get pet response
// swagger:response GetPetResponse
type GetPetResponse struct {
	// in:body
	Body struct {
		Msg  string     `json:"message"`
		Data domain.Pet `json:"data"`
	} `json:"body"`
}

// model for pet category query param
// swagger:parameters ListPets
type PetCategoryQueryParam struct {
	// pet category to filter
	//
	// in:query
	Category string `json:"category"`
}

// model for list pets response
// swagger:response ListPetsResponse
type ListPetsResponse struct {
	// in:body
	Body struct {
		Msg  string       `json:"message"`
		Data []domain.Pet `json:"data"`
	}
}

// model for add pet request
// swagger:parameters AddPet
type AddPetRequest struct {
	// in: body
	// required: true
	Body domain.Pet `json:"body"`
}

// model for add success response without data
// swagger:response SuccessRespWithoutData
type SuccessRespWithoutData struct {
	// in:body
	Body struct {
		Msg string `json:"message"`
	}
}
