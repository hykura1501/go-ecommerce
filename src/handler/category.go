package handler

import (
	"BE_Ecommerce/db/repositories"
	"BE_Ecommerce/src/entity"
	"BE_Ecommerce/src/pkg"
	"log"
	"net/http"
	"strconv"

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

func (server *Server) getCategoryById(c echo.Context) error {
	categoryIdStr := c.Param("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}
	category, err := repositories.GetCategoryById(server.dbInstance, categoryId)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorGetData, err))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccessWithData(pkg.InfoGetCategorySuccess, category))
}

func (server *Server) createCategory(c echo.Context) error {
	var req entity.CategoryRequest
	if err := c.Bind(&req); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}
	if req.Thumbnail != nil {
		url, err := pkg.UploadSingleImage(req.Thumbnail, pkg.CategoryImageFolder)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorUploadImage, err))
		}
		req.ThumbnailUrl = url
	}
	err := repositories.CreateCategory(server.dbInstance, &req)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorCreateData, err))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccess(pkg.InfoCreateCategorySuccess))
}

func (server *Server) updateCategory(c echo.Context) error {

	categoryIdStr := c.Param("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}
	var req entity.CategoryRequest
	if err := c.Bind(&req); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}

	if req.Thumbnail != nil {
		url, err := pkg.UploadSingleImage(req.Thumbnail, pkg.CategoryImageFolder)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorUploadImage, err))
		}
		req.ThumbnailUrl = url
	}
	err = repositories.UpdateCategory(server.dbInstance, categoryId, &req)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorUpdateData, err))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccess(pkg.InfoUpdateCategorySuccess))
}

func (server *Server) deleteCategory(c echo.Context) error {
	categoryIdStr := c.Param("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}
	err = repositories.DeleteCategory(server.dbInstance, categoryId)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorDeleteData, err))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccess(pkg.InfoDeleteCategorySuccess))
}

func (server *Server) getProductsByCategoryId(c echo.Context) error {
	var paging entity.Paging
	if err := c.Bind(&paging); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}
	categoryIdStr := c.Param("id")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}

	paging.ToPaging()

	products, err := repositories.GetProductsByCategoryId(server.dbInstance, categoryId, &paging)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorGetData, err))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccessWithData(pkg.InfoGetProductsByCategoryIdSuccess, echo.Map{
		"products": products,
		"paging":   paging,
	}))
}
