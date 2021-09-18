package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/romanzh1/golang-parser-xml/pkg/parse"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I'm alive"))
	})

	router.Get("/file", func(w http.ResponseWriter, r *http.Request) {
		xmlFile, err := os.Open("cmd/parser-xml/export_yandex_leningradka_msk.xml")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully Opened users.xml")
		defer xmlFile.Close()

		projects := parse.FromXML(xmlFile)
		projectJSON, err := json.MarshalIndent(projects, "	", "  ")
		if err != nil {
			fmt.Println(err)
		}
		w.Write([]byte(projectJSON))
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

	http.ListenAndServe(":4000", router)
}
