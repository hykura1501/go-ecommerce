package config

import (
	"os"
	"strconv"
)

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

	// Facebook OAuth
	FacebookClientID     string
	FacebookClientSecret string
	FacebookRedirectURI  string

	// Cloudinary (Upload Images)
	CloudName      string
	CloudApiKey    string
	CloudApiSecret string

	//Provider
	ProviderLocal    int
	ProviderGoogle   int
	ProviderFacebook int
}

var env *Env

func LoadEnv() *Env {
	if env != nil {
		return env
	}

	providerLocal := os.Getenv("PROVIDER_LOCAL")
	providerGoogle := os.Getenv("PROVIDER_GOOGLE")
	providerFacebook := os.Getenv("PROVIDER_FACEBOOK")

	// parse to int
	providerLocalInt, _ := strconv.Atoi(providerLocal)
	providerGoogleInt, _ := strconv.Atoi(providerGoogle)
	providerFacebookInt, _ := strconv.Atoi(providerFacebook)

	env = &Env{
		DBHost:               os.Getenv("DB_HOST"),
		DBUser:               os.Getenv("DB_USER"),
		DBPassword:           os.Getenv("DB_PASSWORD"),
		DBName:               os.Getenv("DB_NAME"),
		DBPort:               os.Getenv("DB_PORT"),
		Port:                 os.Getenv("PORT"),
		JWTSecret:            os.Getenv("JWT_SECRET"),
		GoogleClientID:       os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleClientSecret:   os.Getenv("GOOGLE_CLIENT_SECRET"),
		GoogleRedirectURI:    os.Getenv("GOOGLE_CLIENT_REDIRECT_URI"),
		FacebookClientID:     os.Getenv("FACEBOOK_CLIENT_ID"),
		FacebookClientSecret: os.Getenv("FACEBOOK_CLIENT_SECRET"),
		FacebookRedirectURI:  os.Getenv("FACEBOOK_CLIENT_REDIRECT_URI"),
		CloudName:            os.Getenv("CLOUDINARY_CLOUD_NAME"),
		CloudApiKey:          os.Getenv("CLOUDINARY_API_KEY"),
		CloudApiSecret:       os.Getenv("CLOUDINARY_API_SECRET"),
		ProviderLocal:        providerLocalInt,
		ProviderGoogle:       providerGoogleInt,
		ProviderFacebook:     providerFacebookInt,
	}

	return env
}
