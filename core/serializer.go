package core

import "net/http"

type HTTPSerializer interface {
	Decode(r *http.Request) (interface{}, error)
	Encode(w http.ResponseWriter, statusCode int, ressource interface{}) error
}
