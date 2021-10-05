package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/romanzh1/golang-parser-xml/pkg/parse"
)

func ProcessRequests(router *chi.Mux) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I'm alive"))
	})

	router.Post("/api/parse", func(w http.ResponseWriter, r *http.Request) {
		type link struct {
			URL string `json:"url"`
		}
		var ref link
		err := json.NewDecoder(r.Body).Decode(&ref)
		if err != nil {
			fmt.Println(err)
		}

		projects := parse.FromURL(ref.URL)
		projectJSON, err := json.MarshalIndent(projects, "	", "  ")
		if err != nil {
			fmt.Println(err)
		}
		w.Write([]byte(projectJSON))
	})
}
