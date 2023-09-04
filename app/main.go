package main

import (
	"book/data"
	d "book/deliveries/http"
	"book/repository"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	defaultBook := data.Book{
		Id:          "111",
		Author:      "Tolkien",
		Quantity:    4,
		Price:       111.11,
		ReleaseDate: "1990",
		Description: "Classic",
	}

	repository.InitBook(defaultBook)

	prefix := "/book"
	mux.Handle(prefix+"/", http.StripPrefix(prefix+"/", http.HandlerFunc(d.HandleBooksWithId)))
	mux.Handle(prefix, http.HandlerFunc(d.HandleBooks))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("SERVER INIT ERROR")
	}

}
