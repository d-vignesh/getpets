package http

import (
	"github.com/d-vignesh/getpets/pkg/domain"
	"github.com/go-chi/chi"
)

func Routes(r chi.Router, h *Handler) {
	r.Route("/pets", func(r chi.Router) {
		r.With(ValidateQueryParam(ListPetsQuery{})).Get("/", h.ListPets)
		r.With(ValidateBody(domain.Pet{})).Post("/", h.AddPet)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(ValidateURLParam("id"))
			r.Get("/", h.GetPet)
			r.Delete("/", h.DeletePet)
		})
	})
}
