package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/romanzh1/golang-parser-xml/internal/server"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// include for parsing xml file into structures and printing to console in json
	// pathXMLFile := "cmd/parser-xml/export_yandex_leningradka_msk.xml"
	// file.PrintJSONData(pathXMLFile)

	// request processing function
	server.ProcessRequests(router)
	http.ListenAndServe(":4000", router)
}
