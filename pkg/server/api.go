package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	a := &api{}

	r := mux.NewRouter()
	//r.HandleFunc("/home", a.homeLink).Methods(http.MethodGet)
	//r.HandleFunc("/event", a.checkMutant).Methods(http.MethodPost)
	r.HandleFunc("/mutant", a.checkMutant).Methods(http.MethodPost)
	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
