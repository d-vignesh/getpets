package http

import (
	"fmt"
	"net/http"

	"github.com/d-vignesh/getpets/pkg/domain"
	"github.com/google/uuid"
)

type Handler struct {
	Svc domain.PetSvc
}

func NewHandler(svc domain.PetSvc) *Handler {
	return &Handler{
		Svc: svc,
	}
}

func (h *Handler) GetPet(w http.ResponseWriter, r *http.Request) {
	ID := r.Context().Value("id").(string)
	fmt.Println("pet id : ", ID)
	petID, err := uuid.Parse(ID)
	if err != nil {
		resp := Resp{
			Code: http.StatusBadRequest,
			Msg:  "invalid pet_id provided in url param",
		}
		respond(w, r, &resp)
		return
	}
	if petID == uuid.Nil {
		resp := Resp{
			Code: http.StatusBadRequest,
			Msg:  "please provide the pet id to retrieve",
		}
		respond(w, r, &resp)
		return
	}
	pet, err := h.Svc.Get(petID)
	if err != nil {
		fmt.Println(fmt.Errorf("error - fetching pet detail from db failed, err : %v", err))
		resp := Resp{
			Code: http.StatusInternalServerError,
			Msg:  "fetching pet detail failed. please try again later",
		}
		respond(w, r, &resp)
		return
	}
	resp := Resp{
		Code: http.StatusOK,
		Msg:  "success",
		Data: pet,
	}
	respond(w, r, &resp)
}

func (h *Handler) ListPets(w http.ResponseWriter, r *http.Request) {
	query := r.Context().Value(QUERY).(*ListPetsQuery)
	pets, err := h.Svc.List(query.Category)
	if err != nil {
		fmt.Println(fmt.Errorf("error - fetching pet details for given category from db failed, err : %v", err))
		resp := Resp{
			Code: http.StatusInternalServerError,
			Msg:  "fetching all pets detail failed. please try again later",
		}
		respond(w, r, &resp)
		return
	}
	resp := Resp{
		Code: http.StatusOK,
		Msg:  "success",
		Data: pets,
	}
	respond(w, r, &resp)
}

func (h *Handler) AddPet(w http.ResponseWriter, r *http.Request) {
	req := r.Context().Value(BODY).(*domain.Pet)
	err := h.Svc.Create(req)
	if err != nil {
		fmt.Println(fmt.Errorf("error - adding new pet detail to db failed, err : %v", err))
		resp := Resp{
			Code: http.StatusInternalServerError,
			Msg:  "adding new pet failed. please try again later",
		}
		respond(w, r, &resp)
		return
	}
	resp := Resp{
		Code: http.StatusOK,
		Msg:  "success",
	}
	respond(w, r, &resp)
}

func (h *Handler) DeletePet(w http.ResponseWriter, r *http.Request) {
	ID := r.Context().Value("id").(string)
	fmt.Println("pet id : ", ID)
	petID, err := uuid.Parse(ID)
	if err != nil {
		resp := Resp{
			Code: http.StatusBadRequest,
			Msg:  "invalid pet_id provided in url param",
		}
		respond(w, r, &resp)
		return
	}
	if petID == uuid.Nil {
		resp := Resp{
			Code: http.StatusBadRequest,
			Msg:  "please provide the pet id to retrieve",
		}
		respond(w, r, &resp)
		return
	}
	err = h.Svc.Delete(petID)
	if err != nil {
		fmt.Println(fmt.Errorf("error - deleting pet record from db failed, err: %v", err))
		resp := Resp{
			Code: http.StatusInternalServerError,
			Msg:  "unable to delete the pet details. please try again later",
		}
		respond(w, r, &resp)
		return
	}
	resp := Resp{
		Code: http.StatusOK,
		Msg:  "success",
	}
	respond(w, r, &resp)
}
