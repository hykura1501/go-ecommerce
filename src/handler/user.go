package handler

import (
	"BE_Ecommerce/db/repositories"
	"BE_Ecommerce/src/config"
	"BE_Ecommerce/src/entity"
	"BE_Ecommerce/src/pkg"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (server *Server) getAllUsers(c echo.Context) error {
	var query entity.UserQuery
	if err := c.Bind(&query); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}
	query.ToPaging()
	users, err := repositories.GetAllUsers(server.dbInstance, &query)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorGetData, err))
	}

	return c.JSON(http.StatusOK, pkg.ResponseSuccessWithData(pkg.InfoGetAllUsersSuccess, echo.Map{
		"users":  users,
		"paging": query.Paging,
	}))
}

func (server *Server) getUserMe(c echo.Context) error {
	tokenUser := c.Get("user").(*jwt.Token)
	claims := tokenUser.Claims.(*pkg.JwtCustomClaims)
	user, err := repositories.GetUserById(server.dbInstance, claims.UserID)

	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorGetData, err))
	}

	if user == nil {
		return c.JSON(http.StatusNotFound, pkg.ResponseError(pkg.ErrorUserNotFound, nil))
	}

	return c.JSON(http.StatusOK, pkg.ResponseSuccessWithData(pkg.InfoGetUserMeSuccess, echo.Map{
		"user": user,
	}))
}

func (server *Server) getUserById(c echo.Context) error {
	userIdStr := c.Param("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}
	user, err := repositories.GetUserById(server.dbInstance, userId)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorGetData, err))
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, pkg.ResponseError(pkg.ErrorUserNotFound, nil))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccessWithData(pkg.InfoGetUserByIdSuccess, echo.Map{
		"user": user,
	}))
}

func (server *Server) getNewUsersStatistics(c echo.Context) error {
	result, err := repositories.GetNewUsersStatistics(server.dbInstance)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorGetData, err))
	}

	return c.JSON(http.StatusOK, pkg.ResponseSuccessWithData(pkg.InfoGetNewUsersStatisticsSuccess, echo.Map{
		"statistics": result,
	}))
}

func (server *Server) resetPassword(c echo.Context) error {
	tokenUser := c.Get("user").(*jwt.Token)
	claims := tokenUser.Claims.(*pkg.JwtCustomClaims)
	user, err := repositories.GetUserById(server.dbInstance, claims.UserID)

	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorGetData, err))
	}

	config := config.LoadEnv()
	if *user.LoginProvider != config.ProviderLocal {
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorResetPasswordNotLocal, errors.New(pkg.ErrorUserNotFound)))
	}
	var resetPasswordRequest entity.ResetPasswordRequest
	if err := c.Bind(&resetPasswordRequest); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}
	hashedPassword := pkg.HashAndSalt([]byte(resetPasswordRequest.NewPassword))

	// update db
	err = repositories.ResetPassword(server.dbInstance, user.UserId, hashedPassword)

	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorUpdateData, err))
	}

	return c.JSON(http.StatusOK, pkg.ResponseSuccess(pkg.InfoResetPasswordSuccess))
}

func (server *Server) updateUser(c echo.Context) error {
	tokenUser := c.Get("user").(*jwt.Token)
	userId := tokenUser.Claims.(*pkg.JwtCustomClaims).UserID
	var updateUserRequest entity.UpdateUserRequest
	if err := c.Bind(&updateUserRequest); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}

	if updateUserRequest.Avatar != nil {
		avatar, err := pkg.UploadSingleImage(updateUserRequest.Avatar, pkg.AvatarFolder)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorUploadImage, err))
		}
		updateUserRequest.AvatarUrl = avatar
	}

	err := repositories.UpdateUser(server.dbInstance, userId, &updateUserRequest)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorUpdateData, err))
	}

	return c.JSON(http.StatusOK, pkg.ResponseSuccess(pkg.InfoUpdateUserSuccess))
}

func (server *Server) deleteUser(c echo.Context) error {
	userIdStr := c.Param("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}
	err = repositories.DeleteUser(server.dbInstance, userId)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorDeleteData, err))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccess(pkg.InfoDeleteUserSuccess))
}
