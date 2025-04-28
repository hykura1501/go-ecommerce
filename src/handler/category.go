package handler

import (
	"BE_Ecommerce/db/repositories"
	"BE_Ecommerce/src/pkg"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (server *Server) getAllCategories(c echo.Context) error {
	categories, err := repositories.GetAllCategories(server.dbInstance)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorGetData, err))
	}

	treeCategories := pkg.GetCategoryTree(categories, 0)
	return c.JSON(http.StatusOK, pkg.ResponseSuccessWithData(pkg.InfoGetCategoriesSuccess, echo.Map{
		"categories":      categories,
		"tree_categories": treeCategories,
	}))
}
