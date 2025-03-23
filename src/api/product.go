package api

import (
	db "BE_Ecommerce/db/sqlc"
	"BE_Ecommerce/src/helpers"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (server *Server) getAllProducts(c echo.Context) error {
	page := c.QueryParam("page")
	perPage := c.QueryParam("per_page")

	paging := helpers.ToPaging(page, perPage)

	products, err := server.store.GetAllProducts(context.Background(), db.GetAllProductsParams{
		Offset: paging.PerPage * (paging.Page - 1),
		Limit:  paging.PerPage,
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": "Failed to get all products",
			"error":   err.Error(),
		})
	}

	counts, err := server.store.CountProducts(context.Background())

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": "Failed to get all products",
			"error":   err.Error(),
		})
	}

	paging.TotalItems = int32(counts)
	paging.TotalPages = (paging.TotalItems + paging.PerPage - 1) / paging.PerPage

	return c.JSON(http.StatusOK, echo.Map{
		"code":     http.StatusOK,
		"products": products,
		"paging":   paging,
	})
}

func (server *Server) getSpecialProducts(c echo.Context) error {
	newArrivalProducts, _ := server.store.GetNewArrivalProducts(context.Background())

	return c.JSON(200, echo.Map{
		"products": newArrivalProducts,
	})
}
