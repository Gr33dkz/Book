all: runall

startApp:
	go run cmd/app/main.go

swaggerGen:
	swag init -g cmd/app/swagui.go

swagSrvUp:
	go run cmd/app/swagui.go