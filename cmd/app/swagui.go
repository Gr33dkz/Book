package main

import (
	_ "book/docs"
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

	swaggerRoute := chi.NewRouter()

	swaggerRoute.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:1323/swagger/doc.json"), //The url pointing to API definition
	))
	err := http.ListenAndServe(":1323", swaggerRoute)
	if err != nil {
		log.Fatal("SERVER INIT ERROR") // TODO ADD LOGGER
	}
}
