package handler

import (
	"net/http"

	"gitbub.com/gocisse/dream_girl_ai/view/home"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	return home.Index().Render(r.Context(), w)
}