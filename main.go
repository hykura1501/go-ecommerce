package main

import (
	"BE_Ecommerce/db"
	"BE_Ecommerce/src/config"
	"BE_Ecommerce/src/handler"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := config.LoadEnv()

	dbInstance, err := db.Connect()

	if err != nil {
		log.Fatal("cannot connect to db")
	}

	defer db.Close(dbInstance)

	goth.UseProviders(
		google.New(config.GoogleClientID, config.GoogleClientSecret, config.GoogleRedirectURI, "email", "profile"),
	)

	server, err := handler.NewServer(dbInstance)
	if err != nil {
		log.Fatal("cannot create echo server")
	}
	server.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
