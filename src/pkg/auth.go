package pkg

import (
	"BE_Ecommerce/src/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func NewGoogleOAuthClient() *oauth2.Config {
	config := config.LoadEnv()
	return &oauth2.Config{
		ClientID:     config.GoogleClientID,
		ClientSecret: config.GoogleClientSecret,
		RedirectURL:  "postmessage",
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}
}
