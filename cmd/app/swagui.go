package main

import (
	_ "book/docs"
	"book/internal/app"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"log"
	"net/http"
)

// @title Book
// @version 1.0
// @description This is a simple Books CRUD service
// @termsOfService https://LinkToTerms.com

// @contact.name Book Support
// @contact.url http://contacturl.com
// @contact.email books@swagger.io

// @license.name Apache 2.0
// @license.url http://www.book.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {

	cfg, err := app.NewConfig()
	if err != nil {
		panic(err)
	}

	swaggerRoute := chi.NewRouter()

	swaggerRoute.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(cfg.Swagger.Url), //The url pointing to API definition
	))
	err = http.ListenAndServe(cfg.Swagger.Port, swaggerRoute)
	if err != nil {
		log.Fatal("SERVER INIT ERROR") // TODO ADD LOGGER
	}
}
