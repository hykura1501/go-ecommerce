package api

import (
	db "BE_Ecommerce/db/sqlc"
	"BE_Ecommerce/src/api/entity"
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func (server *Server) getAllProducts(c echo.Context) error {
	var req entity.ProductRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": "Failed to bind request",
			"error":   err.Error(),
		})
	}

	paging := entity.ToPaging(req.Page, req.PerPage)

	sort, isValid := entity.SortProductsOptions[req.Sort]

	if !isValid {
		req.SortBy = "p.product_id"
		req.SortValue = "asc"
	} else {
		req.SortBy, req.SortValue = strings.Split(sort, " ")[0], strings.Split(sort, " ")[1]
	}

	getParams := db.GetAllProductsParams{
		Offset:     paging.Offset,
		Limit:      paging.Limit,
		Search:     &req.Search,
		SortBy:     &req.SortBy,
		SortValue:  &req.SortValue,
		CategoryID: &req.CategoryId,
		Tag:        &req.Tag,
		PriceMin:   &req.PriceMin,
		PriceMax:   &req.PriceMax,
	}

	products, err := server.store.GetAllProducts(context.Background(), getParams)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": "Failed to get all products",
			"error":   err.Error(),
		})
	}

	countParams := db.CountProductsParams{
		Search:     &req.Search,
		CategoryID: &req.CategoryId,
		Tag:        &req.Tag,
		PriceMin:   &req.PriceMin,
		PriceMax:   &req.PriceMax,
	}

	counts, err := server.store.CountProducts(context.Background(), countParams)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": "Failed to get all products",
			"error":   err.Error(),
		})
	}

	paging.SetTotalPages(counts)

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
