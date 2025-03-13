package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	UserID     int32 `json:"user_id"`
	Permission int32 `json:"permission"`
	jwt.RegisteredClaims
}

var (
	JWT_SECRET = os.Getenv("JWT_SECRET")
)

func GenerateToken(userId int32, permission int32) (string, error) {
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
