package user

import (
	"go-rest/core"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

func MountUserRouter(r *chi.Mux, db *sqlx.DB) {
	userDAO := UserDAO{DB: db}
	r.Mount("/users", core.NewRouter[User](userDAO))
}
