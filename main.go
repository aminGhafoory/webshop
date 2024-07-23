package main

import (
	"context"
	"net/http"

	"github.com/aminGhafoory/webshop/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		c := views.Hello("amin")
		c.Render(context.Background(), w)
	})

	r.Get("/public/*", func(w http.ResponseWriter, r *http.Request) {
		http.FileServerFS(FS).ServeHTTP(w, r)

	})
	http.ListenAndServe(":3000", r)
}
