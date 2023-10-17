package main

import (
	"book/internal/app"
	d "book/internal/deliveries/http/v1"
	"book/internal/repository"
	"book/internal/service"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	cfg, err := app.NewConfig()
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(
		cfg.Database.DriverName, makeDsn(&cfg.Database),
	) // TODO MOVE TO CONNECTIONS DIR
	if err != nil {
		fmt.Println("CONNECTION TO DB ERROR", err) // TODO ADD LOGGER
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("DB ERROR", err) // TODO ADD LOGGER
	}

	st := repository.NewRepo(db)
	srv := service.New(st)
	handler := d.NewHandler(srv)
	mux := handler.Register()
	hh := cors.AllowAll().Handler(mux)

	err = http.ListenAndServe(":8080", hh) // TODO ADD LOGGER
	if err != nil {
		log.Fatal("SERVER INIT ERROR", err) // TODO ADD LOGGER
	}

}

func makeDsn(dbCfg *app.DatabaseConfig) string {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		dbCfg.HostName, dbCfg.Port, dbCfg.UserName, dbCfg.DbName, dbCfg.SslMode, dbCfg.Password)
	return dsn
}

// TODO ADD ALL DEFINIONS IN RUN FILE
