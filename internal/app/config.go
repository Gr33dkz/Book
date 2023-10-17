package app

import (
	"github.com/caarlos0/env"
	"github.com/mcuadros/go-defaults"
)

type Config struct {
	Database DatabaseConfig
	Swagger  SwaggerConfig
}

type DatabaseConfig struct {
	DriverName string `env:"DRIVER_NAME" default:"postgres"`
	HostName   string `env:"HOST_NAME" default:"127.0.0.1"`
	Port       string `env:"PORT" default:"5432"`
	UserName   string `env:"USER_NAME" default:"postgres"`
	DbName     string `env:"DB_NAME" default:"postgres"`
	SslMode    string `env:"SSL_MODE" default:"disable"`
	Password   string `env:"PASSWORD" default:"postgres"`
}

type SwaggerConfig struct {
	Url  string `env:"URL" default:"http://localhost:1323/swagger/doc.json"`
	Port string `env:"PORT" default:":1323"`
}

func NewConfig() (*Config, error) {
	cfg := new(Config)

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	defaults.SetDefaults(cfg)
	return cfg, nil
}
