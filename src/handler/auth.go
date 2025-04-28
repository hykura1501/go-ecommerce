package handler

import (
	"BE_Ecommerce/db/repositories"
	"BE_Ecommerce/src/entity"
	"BE_Ecommerce/src/pkg"
	"context"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

func (server *Server) localLogin(c echo.Context) error {
	var req entity.LocalLoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}

	user, err := repositories.GetUserByUsername(server.dbInstance, req.Username)

	if err != nil {
		return c.JSON(http.StatusNotFound, pkg.ResponseError(pkg.ErrorUserNotExisted, err))
	}

	isValidPassword := pkg.ComparePasswords(*user.Password, []byte(req.Password))

	if !isValidPassword {
		return c.JSON(http.StatusNotFound, pkg.ResponseError(pkg.ErrorPasswordIncorrect, err))
	}

	// return token
	token, err := pkg.GenerateToken(user.UserId, user.Permission)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorGenerateToken, err))
	}

	return c.JSON(http.StatusOK, pkg.ResponseSuccessWithData(pkg.InfoLoginSuccess, echo.Map{
		"token": token,
	}))

}

func (server *Server) register(c echo.Context) error {
	var req entity.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}

	hashedPassword := pkg.HashAndSalt([]byte(req.Password))

	PROVIDER_LOCAL, _ := strconv.Atoi(os.Getenv("PROVIDER_LOCAL"))

	user, err := repositories.CreateUser(server.dbInstance, &entity.User{
		Username:      &req.Username,
		Password:      &hashedPassword,
		Fullname:      &req.Fullname,
		LoginProvider: &PROVIDER_LOCAL,
		Permission:    0,
	})

	if err != nil {
		return c.JSON(http.StatusConflict, pkg.ResponseError(pkg.ErrorUserExisted, err))
	}

	token, err := pkg.GenerateToken(user.UserId, user.Permission)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorGenerateToken, err))
	}
	return c.JSON(http.StatusCreated, pkg.ResponseSuccessWithData(pkg.InfoRegisterSuccess, echo.Map{
		"token": token,
	}))
}

func (server *Server) googleLogin(c echo.Context) error {
	req := c.Request()
	res := c.Response()
	req = req.WithContext(context.WithValue(req.Context(), "provider", "google"))
	gothic.BeginAuthHandler(res, req)
	return nil
}

func (server *Server) googleLoginCallback(c echo.Context) error {
	return c.Redirect(http.StatusOK, os.Getenv("CLIENT_URL"))
}
