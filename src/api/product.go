package api

import (
	"BE_Ecommerce/db/repositories"
	"BE_Ecommerce/src/entity"
	"BE_Ecommerce/src/pkg"
	"log"
	"net/http"

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
	// newArrivalProducts, _ := server.store.GetNewArrivalProducts(context.Background())

	return c.JSON(200, echo.Map{
		"products": nil,
	})
}
