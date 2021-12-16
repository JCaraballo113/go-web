package main

import (
	"fmt"
	"go-web/pkg/config"
	"go-web/pkg/handlers"
	"go-web/pkg/render"
	"log"
	"net/http"
)

const PORT = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	fmt.Printf("Listening on port %s", PORT)

	srv := &http.Server{
		Addr: PORT,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}