package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/romanzh1/golang-parser-xml/pkg/parse"
)

func main() {
	projects := parse.FromXML()
	fmt.Println(len(projects))
	fmt.Println(projects)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	projectJSON, err := json.MarshalIndent(projects, "	", "  ")
	if err != nil {
		fmt.Println(err)
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I'm alive"))
		w.Write([]byte(projectJSON))
	})
	http.ListenAndServe(":4000", router)
}
