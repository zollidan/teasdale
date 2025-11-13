package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/zollidan/teasdale/config"
	"github.com/zollidan/teasdale/database"
)

func main() {
	cfg := config.New()
	_ = database.Init(cfg)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	
	fmt.Printf("Starting server on %s\n", cfg.ServerAddress)
	err := http.ListenAndServe(cfg.ServerAddress, r)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
