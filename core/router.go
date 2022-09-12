package core

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func decode[T any](r *http.Request) (*T, error) {
	var resource T
	if err := json.NewDecoder(r.Body).Decode(&resource); err != nil {
		return nil, err
	}
	return &resource, nil

}

func encode[T any](w http.ResponseWriter, statusCode int, resource T) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(resource)
}

func NewRouter[T any](dao DAO[T]) http.Handler {
	r := chi.NewRouter()

	r.Method("GET", "/{id}", Handler(func(w http.ResponseWriter, r *http.Request) *StatusError {
		s := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return &StatusError{Code: 400, Err: Trace(err)}
		}

		resource, err := dao.Get(id)
		if err != nil {
			return &StatusError{Code: 404, Err: Trace(err)}
		}

		err = encode(w, http.StatusOK, resource)
		if err != nil {
			return &StatusError{Code: 500, Err: Trace(err)}
		}

		return nil
	}))

	r.Method("GET", "/", Handler(func(w http.ResponseWriter, r *http.Request) *StatusError {
		resource, err := dao.GetAll()
		if err != nil {
			return &StatusError{Code: 500, Err: Trace(err)}
		}

		err = encode(w, http.StatusOK, resource)
		if err != nil {
			return &StatusError{Code: 500, Err: Trace(err)}
		}

		return nil
	}))

	r.Method("POST", "/", Handler(func(w http.ResponseWriter, r *http.Request) *StatusError {
		resource, err := decode[T](r)
		if err != nil {
			return &StatusError{Code: 400, Err: Trace(err)}
		}

		resource, err = dao.Create(*resource)
		if err != nil {
			return &StatusError{Code: 500, Err: Trace(err)}
		}

		err = encode(w, http.StatusCreated, resource)
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

		_, err = dao.Get(id)
		if err != nil {
			return &StatusError{Code: 404, Err: Trace(err)}
		}

		resource, err := decode[T](r)
		if err != nil {
			return &StatusError{Code: 400, Err: Trace(err)}
		}

		err = dao.Update(id, *resource)
		if err != nil {
			return &StatusError{Code: 500, Err: Trace(err)}
		}

		err = encode(w, http.StatusOK, resource)
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

		_, err = dao.Get(id)
		if err != nil {
			return &StatusError{Code: 404, Err: Trace(err)}
		}

		err = dao.Delete(id)
		if err != nil {
			return &StatusError{Code: 500, Err: Trace(err)}
		}

		return nil
	}))

	return r
}
