package user

import (
	"go-rest/core"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

func MountUserRouter(r *chi.Mux, db *sqlx.DB) {
	var userDAO core.DAO[User] = UserDAO{DB: db}
	r.Mount("/users", core.NewRouter(userDAO))
}
