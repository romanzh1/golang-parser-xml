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
		// var url struct {
		// 	l string			// TODO change the map to struct or some other way
		// }
		url := map[string]string{}
		err := json.NewDecoder(r.Body).Decode(&url)
		if err != nil {
			fmt.Println(err)
		}

		projects := parse.FromURL(url["url"])
		projectJSON, err := json.MarshalIndent(projects, "	", "  ")
		if err != nil {
			fmt.Println(err)
		}
		w.Write([]byte(projectJSON))
	})
}
