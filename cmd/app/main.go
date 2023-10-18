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
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	cfg, err := app.NewConfig()
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(
		cfg.Database.DriverName, makeDsn(&cfg.Database),
	) // TODO MOVE TO CONNECTIONS DIR
	if err != nil {
		log.Fatal("connection open error", err)
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("connection to db error ", err)
		panic(err)
	}

	st := repository.NewRepo(db)
	srv := service.New(st)
	handler := d.NewHandler(srv)
	mux := handler.Register()
	hh := cors.AllowAll().Handler(mux)

	log.WithField("port", cfg.Service.Port).Info("Server started at")
	err = http.ListenAndServe(cfg.Service.Port, hh)
	if err != nil {
		log.Fatal("service init error", err)
		panic(err)
	}
}

func makeDsn(dbCfg *app.DatabaseConfig) string {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		dbCfg.HostName, dbCfg.Port, dbCfg.UserName, dbCfg.DbName, dbCfg.SslMode, dbCfg.Password)
	return dsn
}
