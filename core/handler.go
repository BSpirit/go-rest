package core

import (
	"log"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) *StatusError

// ServeHTTP allows our Handler type to satisfy http.Handler.
func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	if err != nil {
		log.Printf("%s", err.Unwrap())
		http.Error(w, http.StatusText(err.Code), err.Code)
	}
}
