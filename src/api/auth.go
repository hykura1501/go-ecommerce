package api

import (
	db "BE_Ecommerce/db/sqlc"
	"BE_Ecommerce/src/helpers"
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

type localLoginRequest struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type registerRequest struct {
	Fullname *string `json:"fullname"`
	Username *string `json:"username"`
	Password *string `json:"password"`
}

func (server *Server) localLogin(c echo.Context) error {
	var req localLoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": "Error binding",
		})
	}

	fmt.Println(req)

	user, err := server.store.GetUserByUsername(context.Background(), req.Username)

	fmt.Println(user)

	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"code":    http.StatusNotFound,
			"message": "User is not existed!",
		})
	}
	isTheSamePassword := helpers.ComparePasswords(*user.Password, []byte(*req.Password))
	if !isTheSamePassword {
		return c.JSON(http.StatusNotFound, echo.Map{
			"code":    http.StatusNotFound,
			"message": "Password is incorrect!",
		})
	}
	// return token
	token, err := helpers.GenerateToken(user.UserID, user.Permission)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"code":    http.StatusInternalServerError,
			"message": "Failed to generate token!",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"code":    http.StatusOK,
		"token":   token,
		"message": "Login successfully.",
	})

}

func (server *Server) register(c echo.Context) error {
	var req registerRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": "Error binding",
		})
	}

	hashedPassword := helpers.HashAndSalt([]byte(*req.Password))
	fmt.Println(hashedPassword)

	value, _ := strconv.Atoi(os.Getenv("PROVIDER_LOCAL"))
	PROVIDER_LOCAL := int32(value)

	user, err := server.store.CreateUser(context.Background(), db.CreateUserParams{
		Fullname:      req.Fullname,
		Username:      req.Username,
		Password:      &hashedPassword,
		Permission:    0,
		LoginProvider: &PROVIDER_LOCAL,
	})

	if err != nil {
		return c.JSON(http.StatusConflict, echo.Map{
			"code":    http.StatusConflict,
			"message": "Email already existed",
		})
	}
	token, err := helpers.GenerateToken(user.UserID, user.Permission)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"code":    http.StatusInternalServerError,
			"message": "Failed to generate token!",
		})
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"code":    http.StatusCreated,
		"message": "Register successfully.",
		"token":   token,
	})
}
