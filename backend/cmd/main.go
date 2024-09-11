package main

import (
	"backend/internal/config"
	"backend/internal/handler"
	db "backend/internal/mongo"
	"backend/internal/server"
	"backend/internal/storer"
	"os"
)

func main() {
	config.LoadEnv()

	secretKey := os.Getenv("SECRET_KEY")

	mongo := db.InitDB()

	st := storer.NewMongoStorer(mongo)

	srv := server.NewServer(st)

	hdl := handler.NewHandler(srv, secretKey)
	
	handler.RegisterRoutes(hdl)
	handler.Start(":3000")
}