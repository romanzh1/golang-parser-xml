package main

import (
	"fmt"
	"log"
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
	server.ProcessRequests(router) // TODO handle the error probably

	fmt.Println("Server start on port 4000")
	err := http.ListenAndServe(":4000", router)
	if err != nil {
		log.Fatalln(err)
	}

}
