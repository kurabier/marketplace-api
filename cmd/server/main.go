package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/kurabier/marketplace-api/internal/gen"
	"github.com/kurabier/marketplace-api/internal/handlers"
)

func main() {
	h := handlers.New()

	r := chi.NewRouter()
	gen.HandlerFromMux(gen.NewStrictHandler(h, nil), r)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
