package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/d-vignesh/getpets/docs"
	"github.com/d-vignesh/getpets/mocks"
	"github.com/d-vignesh/getpets/pkg/app"
	"github.com/d-vignesh/getpets/pkg/db"
	ihttp "github.com/d-vignesh/getpets/pkg/http"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

func main() {
	args := os.Args
	op := "server"
	if len(args) > 1 {
		op = args[0]
	}
	if err := run(op); err != nil {
		fmt.Println(fmt.Errorf("error - server failed to start. err: %v", err))
	}
}

func run(op string) error {
	if op == "http_test" {
		// mocking the pet svc to test our http routes
		svc := mocks.PetSvc{}
		h := ihttp.NewHandler(svc)
		r := chi.NewRouter()
		ihttp.Routes(r, h)
		return http.ListenAndServe(":6000", r)
	}

	// tying up all the components together and running the server
	db, err := db.NewMongoStore()
	if err != nil {
		return errors.Wrap(err, "unable to intialize db")
	}
	svc := app.NewPetSvc(db)
	h := ihttp.NewHandler(svc)
	r := chi.NewRouter()
	ihttp.Routes(r, h)
	fmt.Println("starting server")
	return http.ListenAndServe(":9000", r)
}
