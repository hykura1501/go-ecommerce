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

func (server *Server) getProductDetail(c echo.Context) error {
	productIdStr := c.Param("product_id")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}

	product, err := repositories.GetProductDetail(server.dbInstance, productId)

	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorGetData, err))
	}

	return c.JSON(http.StatusOK, pkg.ResponseSuccessWithData(pkg.InfoGetProductDetailSuccess, product))
}

func (server *Server) getSpecialProducts(c echo.Context) error {
	sizeStr := c.QueryParam("size")
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 20
	}
	results, err := repositories.GetSpecialProducts(server.dbInstance, size)

	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorGetData, err))
	}

	return c.JSON(http.StatusOK, pkg.ResponseSuccessWithData(pkg.InfoGetSpecialProductsSuccess, results))
}

func (server *Server) createProduct(c echo.Context) error {
	var req entity.NewProductRequest
	if err := c.Bind(&req); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}
	// handle uploading images
	if req.Images != nil {
		urls, err := pkg.UploadMultipleImages(req.Images, pkg.ProductImageFolder)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorUploadImage, err))
		}
		req.ImageUrls = urls
	}

	err := repositories.CreateProduct(server.dbInstance, &req)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorCreateData, err))
	}

	return c.JSON(http.StatusOK, pkg.ResponseSuccess(pkg.InfoCreateProductSuccess))
}
