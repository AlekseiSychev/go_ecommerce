package main

import (
	"go-ecom/db"
	_ "go-ecom/docs"
	"go-ecom/server"
	"log"

	"github.com/joho/godotenv"
)

// @title Ecom mini app
func main() {

	// envInit() для локальной разработки без докера
	db.ConnectDatabase()
	server.Run()
}

func envInit() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
