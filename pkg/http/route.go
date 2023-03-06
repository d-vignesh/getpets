package http

import (
	"io/ioutil"
	"net/http"

	"github.com/d-vignesh/getpets/pkg/domain"
	"github.com/go-chi/chi"
	redocmiddleware "github.com/go-openapi/runtime/middleware"
)

func Routes(r chi.Router, h *Handler) {
	opts := redocmiddleware.RedocOpts{Path: "docs", SpecURL: "swagger.json"}
	docsHandler := redocmiddleware.Redoc(opts, nil)
	r.Handle("/docs", docsHandler)
	r.Get("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		spec, err := ioutil.ReadFile("docs/swagger.json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(spec)
	})
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
