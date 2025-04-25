package middlewares

import (
	"BE_Ecommerce/src/pkg"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	echojwt "github.com/labstack/echo-jwt/v4"
)

var (
	JWT_SECRET = os.Getenv("JWT_SECRET")
)

func Authenticate() echo.MiddlewareFunc {

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(pkg.JwtCustomClaims)
		},
		SigningKey: []byte(JWT_SECRET),
	}

	return echojwt.WithConfig(config)
}
