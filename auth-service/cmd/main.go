package main

import (
	"net/http"

	"github.com/AbhiramiRajeev/task-orchestrator/auth-service/internal/handler"
	chi "github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
	

	r.Get("/user/{userID}", func(w http.ResponseWriter, r *http.Request) {

		userID := chi.URLParam(r, "userID")
		w.Write([]byte("Hello user " + userID))
	})

	r.Post("/register",handler.Registerhandler)
	http.ListenAndServe(":8080", r)

}
