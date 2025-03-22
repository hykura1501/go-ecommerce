package api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (server *Server) getAllProducts(c echo.Context) error {
	products, err := server.store.GetAllProducts(context.Background())

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": "Failed to get all products",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"code":     http.StatusOK,
		"products": products,
	})
}

func (server *Server) getSpecialProducts(c echo.Context) error {
	newArrivalProducts, _ := server.store.GetNewArrivalProducts(context.Background())

	return c.JSON(200, echo.Map{
		"products": newArrivalProducts,
	})
}
