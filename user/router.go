package user

import (
	"go-rest/core"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

func MountUserRouter(r *chi.Mux, db *sqlx.DB) {
	userService := UserService{
		UserDAO:        UserDAO{DB: db},
		UserSerializer: UserSerializer{},
	}

	r.Mount("/users", core.NewRouter(userService))
}
