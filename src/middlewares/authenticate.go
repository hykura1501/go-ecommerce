package middlewares

import (
	"BE_Ecommerce/src/pkg"
	"errors"
	"net/http"
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

func IsAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*pkg.JwtCustomClaims)
			if claims.Permission != pkg.RoleAdmin {
				return c.JSON(http.StatusForbidden, pkg.ResponseError(pkg.ErrorForbidden, errors.New(pkg.ErrorForbidden)))
			}
			return next(c)
		}
	}
}
