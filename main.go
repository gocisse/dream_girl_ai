package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	router := chi.NewMux()

	// router.Get("/", http.HandlerFunc)

	log.Fatal(http.ListenAndServe(os.Getenv("HTTP_LISTEN_ADDR"), router))
}

func initEveryThing() error {
	return  godotenv.Load()
}