package main

import (
	d "book/deliveries/http"
	"book/repository"
	"book/service"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres sslmode=disable password=postgres")
	if err != nil {
		fmt.Println("CONNECTION TO DB ERROR", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("DB ERROR", err)
	}

	st := repository.NewRepo(db)
	srv := service.New(st)
	handler := d.NewHandler(srv)
	mux := handler.Register()

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("SERVER INIT ERROR")
	}

}
