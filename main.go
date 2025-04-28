package main

import (
	"BE_Ecommerce/db"
	"BE_Ecommerce/src/handler"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbInstance, err := db.Connect()

	if err != nil {
		log.Fatal("cannot connect to db")
	}

	defer db.Close(dbInstance)

	server, err := handler.NewServer(dbInstance)
	if err != nil {
		log.Fatal("cannot create echo server")
	}
	server.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
