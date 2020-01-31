package main

import (
	"go-rest/core"
	"go-rest/user"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	db, err := core.InitDB("sqlite3", "db/test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	user.MountUserRouter(r, db)

	log.Fatal(http.ListenAndServe(":8080", r))
}
