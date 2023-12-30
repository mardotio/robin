package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"robin/api/src/handlers"
	"robin/api/src/utils"
)

func main() {
	log.Println("Starting API")
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/token", handlers.GetToken)

	serverErr := http.ListenAndServe(fmt.Sprintf(":%d", utils.GetAppConfig().Port), r)

	if nil != serverErr {
		log.Fatal(serverErr)
	}
}
