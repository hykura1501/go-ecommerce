package handler

import (
	"BE_Ecommerce/db/repositories"
	"BE_Ecommerce/src/entity"
	"BE_Ecommerce/src/pkg"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (server *Server) getCart(c echo.Context) error {
	user := c.Get("user").(*jwt.Token).Claims.(*pkg.JwtCustomClaims)
	var paging entity.Paging
	if err := c.Bind(&paging); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}
	paging.ToPaging()
	cart, err := repositories.GetCartByUserId(server.dbInstance, user.UserID, &paging)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorGetData, err))
	}

	return c.JSON(http.StatusOK, pkg.ResponseSuccessWithData(pkg.InfoGetCartSuccess, echo.Map{
		"cart":   cart,
		"paging": paging,
	}))
}

func (server *Server) addToCart(c echo.Context) error {
	user := c.Get("user").(*jwt.Token).Claims.(*pkg.JwtCustomClaims)
	var req entity.AddToCartRequest
	if err := c.Bind(&req); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}
	err := repositories.AddToCart(server.dbInstance, user.UserID, req.ProductId)

	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorAddToCart, err))
	}

	return c.JSON(http.StatusOK, pkg.ResponseSuccess(pkg.InfoAddToCartSuccess))
}

func (server *Server) updateCart(c echo.Context) error {
	user := c.Get("user").(*jwt.Token).Claims.(*pkg.JwtCustomClaims)
	productIdStr := c.Param("product_id")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}

	var req entity.UpdateCartRequest
	if err := c.Bind(&req); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}

	err = repositories.UpdateCart(server.dbInstance, user.UserID, productId, req.Quantity)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorUpdateCart, err))
	}

	return c.JSON(http.StatusOK, pkg.ResponseSuccess(pkg.InfoUpdateCartSuccess))
}

func (server *Server) addItemsToCart(c echo.Context) error {
	user := c.Get("user").(*jwt.Token).Claims.(*pkg.JwtCustomClaims)
	var req entity.AddItemsToCartRequest
	if err := c.Bind(&req); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}

	err := repositories.AddItemsToCart(server.dbInstance, user.UserID, req.Items)

	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorAddToCart, err))
	}

	return c.JSON(http.StatusOK, pkg.ResponseSuccess(pkg.InfoAddToCartSuccess))
}

func (server *Server) deleteCartItem(c echo.Context) error {
	user := c.Get("user").(*jwt.Token).Claims.(*pkg.JwtCustomClaims)
	productIdStr := c.Param("product_id")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}

	err = repositories.DeleteCartItem(server.dbInstance, user.UserID, productId)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorDeleteCartItem, err))
	}

	return c.JSON(http.StatusOK, pkg.ResponseSuccess(pkg.InfoDeleteCartItemSuccess))
}
