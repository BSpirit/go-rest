package core

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func NewRouter(service Service) http.Handler {
	r := chi.NewRouter()

	r.Method("GET", "/{id}", Handler(func(w http.ResponseWriter, r *http.Request) *StatusError {
		s := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return &StatusError{Code: 400, Err: Trace(err)}
		}

		ressource, err := service.DAO().Get(id)
		if err != nil {
			return &StatusError{Code: 404, Err: Trace(err)}
		}

		err = service.Serializer().Encode(w, http.StatusOK, ressource)
		if err != nil {
			return &StatusError{Code: 500, Err: Trace(err)}
		}

		return nil
	}))

	r.Method("GET", "/", Handler(func(w http.ResponseWriter, r *http.Request) *StatusError {
		ressource, err := service.DAO().GetAll()
		if err != nil {
			return &StatusError{Code: 500, Err: Trace(err)}
		}

		err = service.Serializer().Encode(w, http.StatusOK, ressource)
		if err != nil {
			return &StatusError{Code: 500, Err: Trace(err)}
		}

		return nil
	}))

	r.Method("POST", "/", Handler(func(w http.ResponseWriter, r *http.Request) *StatusError {
		ressource, err := service.Serializer().Decode(r)
		if err != nil {
			return &StatusError{Code: 400, Err: Trace(err)}
		}

		ressource, err = service.DAO().Create(ressource)
		if err != nil {
			return &StatusError{Code: 500, Err: Trace(err)}
		}

		err = service.Serializer().Encode(w, http.StatusCreated, ressource)
		if err != nil {
			return &StatusError{Code: 500, Err: Trace(err)}
		}

		return nil
	}))

	r.Method("PUT", "/{id}", Handler(func(w http.ResponseWriter, r *http.Request) *StatusError {
		s := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return &StatusError{Code: 400, Err: Trace(err)}
		}

		_, err = service.DAO().Get(id)
		if err != nil {
			return &StatusError{Code: 404, Err: Trace(err)}
		}

		ressource, err := service.Serializer().Decode(r)
		if err != nil {
			return &StatusError{Code: 400, Err: Trace(err)}
		}

		err = service.DAO().Update(id, ressource)
		if err != nil {
			return &StatusError{Code: 500, Err: Trace(err)}
		}

		err = service.Serializer().Encode(w, http.StatusOK, ressource)
		if err != nil {
			return &StatusError{Code: 500, Err: Trace(err)}
		}

		return nil
	}))

	r.Method("DELETE", "/{id}", Handler(func(w http.ResponseWriter, r *http.Request) *StatusError {
		s := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return &StatusError{Code: 400, Err: Trace(err)}
		}

		_, err = service.DAO().Get(id)
		if err != nil {
			return &StatusError{Code: 404, Err: Trace(err)}
		}

		err = service.DAO().Delete(id)
		if err != nil {
			return &StatusError{Code: 500, Err: Trace(err)}
		}

		return nil
	}))

	return r
}
