package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := initEveryThing(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()

	// router.Get("/", http.HandlerFunc)
	port := os.Getenv("HTTP_LISTEN_ADDR")
	slog.Info("application running on", "port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func initEveryThing() error {
	return  godotenv.Load()
}