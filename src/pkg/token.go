package pkg

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	UserID     int `json:"user_id"`
	Permission int `json:"permission"`
	jwt.RegisteredClaims
}

var (
	JWT_SECRET = os.Getenv("JWT_SECRET")
)

func GenerateToken(userId int, permission int) (string, error) {
	claims := &JwtCustomClaims{
		userId,
		permission,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
