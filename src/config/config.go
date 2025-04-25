package config

import "os"

type Env struct {
	//  DB
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string

	// Server
	Port string

	// JWT
	JWTSecret string

	// Google OAuth
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURI  string
}

var env *Env

func LoadEnv() *Env {
	if env != nil {
		return env
	}

	env = &Env{
		DBHost:             os.Getenv("DB_HOST"),
		DBUser:             os.Getenv("DB_USER"),
		DBPassword:         os.Getenv("DB_PASSWORD"),
		DBName:             os.Getenv("DB_NAME"),
		DBPort:             os.Getenv("DB_PORT"),
		Port:               os.Getenv("PORT"),
		JWTSecret:          os.Getenv("JWT_SECRET"),
		GoogleClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		GoogleRedirectURI:  os.Getenv("GOOGLE_CLIENT_REDIRECT_URI"),
	}

	return env
}
