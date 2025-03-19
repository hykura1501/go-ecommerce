package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (server *Server) getAllProducts(c echo.Context) error {

	return c.JSON(http.StatusOK, echo.Map{})
}
