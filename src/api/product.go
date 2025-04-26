package api

import (
	"BE_Ecommerce/db/repositories"
	"BE_Ecommerce/src/entity"
	"BE_Ecommerce/src/pkg"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (server *Server) getAllProducts(c echo.Context) error {
	var req entity.ProductQuery

	if err := c.Bind(&req); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}

	sort, isValid := entity.SortProductsOptions[req.Sort]

	if isValid {
		req.OrderBy = sort
	} else {
		req.OrderBy = entity.SortProductsOptions["product_id_asc"]
	}

	req.ToPaging()
	products, err := repositories.GetAllProducts(server.dbInstance, &req)

	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorGetData, err))
	}

	return c.JSON(http.StatusOK, pkg.ResponseSuccessWithData(pkg.InfoGetProductsSuccess, echo.Map{
		"products": products,
		"paging":   req.Paging,
	}))
}

func (server *Server) getSpecialProducts(c echo.Context) error {
	sizeStr := c.QueryParam("size")
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 20
	}
	fmt.Println("size", size)
	// results := repositories.GetSpecialProducts(server.dbInstance, size)

	return c.JSON(200, echo.Map{
		"products": nil,
	})
}
