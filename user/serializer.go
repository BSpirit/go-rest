package user

import (
	"encoding/json"
	"net/http"
)

type UserSerializer struct{}

func (s UserSerializer) Decode(r *http.Request) (interface{}, error) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil

}

func (s UserSerializer) Encode(w http.ResponseWriter, statusCode int, ressource interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(ressource)
}
