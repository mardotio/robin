package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"net/http"
	"robin/api/handlers"
	"robin/api/utils"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/token", handlers.GetToken)

	http.ListenAndServe(fmt.Sprintf(":%d", utils.GetAppConfig().Port), r)
}
