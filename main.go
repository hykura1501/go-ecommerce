package main

import (
	db "BE_Ecommerce/db/sqlc"
	"BE_Ecommerce/src/api"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cnnString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	connPool, err := pgxpool.New(context.Background(), cnnString)
	if err != nil {
		log.Fatal("cannot connect to postgres!")
	}

	store := db.NewStore(connPool)

	server, err := api.NewServer(store)
	if err != nil {
		log.Fatal("cannot create echo server")
	}
	server.Start(":8000")
}
