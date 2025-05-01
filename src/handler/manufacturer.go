package handler

import (
	"BE_Ecommerce/db/repositories"
	"BE_Ecommerce/src/pkg"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (server *Server) getAllManufacturers(c echo.Context) error {
	manufacturers, err := repositories.GetAllManufacturers(server.dbInstance)
	if err != nil {
		return c.JSON(http.StatusBadGateway, pkg.ResponseError(pkg.ErrorGetManufacturers, err))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccessWithData(pkg.InfoGetManufacturersSuccess, echo.Map{
		"manufacturers": manufacturers,
	}))
}
