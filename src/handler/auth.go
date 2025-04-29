package handler

import (
	"BE_Ecommerce/db/repositories"
	"BE_Ecommerce/src/config"
	"BE_Ecommerce/src/entity"
	"BE_Ecommerce/src/pkg"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"cloud.google.com/go/auth/credentials/idtoken"
	"github.com/labstack/echo/v4"
)

func (server *Server) loginLocal(c echo.Context) error {
	var req entity.LoginLocalRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}

	user, err := repositories.GetUserByUsername(server.dbInstance, req.Username)

	if err != nil {
		return c.JSON(http.StatusNotFound, pkg.ResponseError(pkg.ErrorUserNotExisted, err))
	}

	isValidPassword := pkg.ComparePasswords(*user.Password, []byte(req.Password))

	if !isValidPassword {
		return c.JSON(http.StatusNotFound, pkg.ResponseError(pkg.ErrorPasswordIncorrect, errors.New(pkg.ErrorPasswordIncorrect)))
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

func (server *Server) loginGoogle(c echo.Context) error {
	var req entity.LoginGoogleRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}

	googleOAuthClient := pkg.NewGoogleOAuthClient()

	tokenByCode, err := googleOAuthClient.Exchange(c.Request().Context(), req.Code)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorExchangeToken, err))
	}
	idToken := tokenByCode.Extra("id_token").(string)

	config := config.LoadEnv()

	payload, err := idtoken.Validate(c.Request().Context(), idToken, config.GoogleClientID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorInvalidIdToken, err))
	}
	loginProvider := config.ProviderGoogle
	providerId := payload.Claims["sub"].(string)

	// check existed user
	existedUser, err := repositories.GetUserByProvider(server.dbInstance, loginProvider, providerId)
	if err != nil {
		// not existed user, create new user
		fullname := payload.Claims["name"].(string)
		avatar := payload.Claims["picture"].(string)
		user := &entity.User{
			LoginProvider: &loginProvider,
			ProviderId:    &providerId,
			Fullname:      &fullname,
			Avatar:        &avatar,
			Permission:    0,
		}
		err := repositories.CreateUserByProvider(server.dbInstance, user)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorCreateUser, err))
		}

		// Generate token
		token, err := pkg.GenerateToken(user.UserId, user.Permission)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorGenerateToken, err))
		}
		return c.JSON(http.StatusOK, pkg.ResponseSuccessWithData(pkg.InfoLoginSuccess, echo.Map{
			"token": token,
		}))
	}
	// Generate token
	token, err := pkg.GenerateToken(existedUser.UserId, existedUser.Permission)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorGenerateToken, err))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccessWithData(pkg.InfoLoginSuccess, echo.Map{
		"token": token,
	}))
}

func (server *Server) loginFacebook(c echo.Context) error {
	var req entity.LoginFacebookRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}
	userInfoURL := fmt.Sprintf("https://graph.facebook.com/me?fields=id,name,email,picture.width(400).height(400)&access_token=%s", req.AccessToken)

	res, err := http.Get(userInfoURL)
	if err != nil || res.StatusCode != 200 {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorGetUserInfo, err))
	}

	defer res.Body.Close()
	var fbUser entity.FacebookUser
	json.NewDecoder(res.Body).Decode(&fbUser)

	config := config.LoadEnv()
	loginProvider := config.ProviderFacebook
	providerId := fbUser.ID

	// check existed user
	existedUser, err := repositories.GetUserByProvider(server.dbInstance, loginProvider, providerId)
	if err != nil {
		// not existed user, create new user
		user := &entity.User{
			LoginProvider: &loginProvider,
			ProviderId:    &providerId,
			Fullname:      &fbUser.Name,
			Avatar:        &fbUser.Picture.Data.URL,
			Permission:    0,
		}
		err := repositories.CreateUserByProvider(server.dbInstance, user)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorCreateUser, err))
		}

		// Generate token
		token, err := pkg.GenerateToken(user.UserId, user.Permission)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorGenerateToken, err))
		}
		return c.JSON(http.StatusOK, pkg.ResponseSuccessWithData(pkg.InfoLoginSuccess, echo.Map{
			"token": token,
		}))
	}
	// Generate token
	token, err := pkg.GenerateToken(existedUser.UserId, existedUser.Permission)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorGenerateToken, err))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccessWithData(pkg.InfoLoginSuccess, echo.Map{
		"token": token,
	}))
}
