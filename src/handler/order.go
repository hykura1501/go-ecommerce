package handler

import (
	"BE_Ecommerce/db/repositories"
	"BE_Ecommerce/src/entity"
	"BE_Ecommerce/src/pkg"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (server *Server) createOrder(c echo.Context) error {
	var req entity.CreateOrderRequest
	user := c.Get("user").(*jwt.Token).Claims.(*pkg.JwtCustomClaims)
	if err := c.Bind(&req); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}

	err := repositories.CreateOrder(server.dbInstance, user.UserID, &req)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorCreateOrder, err))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccess(pkg.InfoCreateOrderSuccess))
}

func (server *Server) getOrderHistoryOfUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token).Claims.(*pkg.JwtCustomClaims)
	userId := user.UserID

	userIdStr := c.Param("user_id")
	if userIdStr != "" {
		id, err := strconv.Atoi(userIdStr)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
		}
		userId = id
	}

	var orderQuery entity.OrderQuery

	if err := c.Bind(&orderQuery); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}

	orderQuery.ToPaging()

	sort, isValid := entity.OrderSortOptions[orderQuery.Order]

	if isValid {
		orderQuery.OrderBy = sort
	} else {
		orderQuery.OrderBy = entity.OrderSortOptions["id_asc"]
	}

	orders, err := repositories.GetOrderHistoryOfUser(server.dbInstance, userId, &orderQuery)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorGetOrderHistory, err))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccessWithData(pkg.InfoGetOrderHistorySuccess, echo.Map{
		"orders": orders,
		"paging": orderQuery.Paging,
	}))
}

func (server *Server) getOrders(c echo.Context) error {
	var orderQuery entity.OrderQuery
	if err := c.Bind(&orderQuery); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}

	orderQuery.ToPaging()

	sort, isValid := entity.OrderSortOptions[orderQuery.Order]

	if isValid {
		orderQuery.OrderBy = sort
	} else {
		orderQuery.OrderBy = entity.OrderSortOptions["id_asc"]
	}

	orders, err := repositories.GetOrders(server.dbInstance, &orderQuery)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorGetOrders, err))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccessWithData(pkg.InfoGetOrdersSuccess, echo.Map{
		"orders": orders,
		"paging": orderQuery.Paging,
	}))
}

func (server *Server) getRevenueStatistics(c echo.Context) error {
	revenues, err := repositories.GetRevenueStatistics(server.dbInstance)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorGetRevenueStatistics, err))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccessWithData(pkg.InfoGetRevenueStatisticsSuccess, echo.Map{
		"revenues": revenues,
	}))
}

func (server *Server) getBestSellersStatistics(c echo.Context) error {
	bestSellers, err := repositories.GetBestSellersStatistics(server.dbInstance)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorGetBestSellersStatistics, err))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccessWithData(pkg.InfoGetBestSellersStatisticsSuccess, echo.Map{
		"best_sellers": bestSellers,
	}))
}

func (server *Server) getTopCustomersStatistics(c echo.Context) error {
	topCustomers, err := repositories.GetTopCustomersStatistics(server.dbInstance)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorGetTopCustomersStatistics, err))
	}
	return c.JSON(http.StatusOK, pkg.ResponseSuccessWithData(pkg.InfoGetTopCustomersStatisticsSuccess, echo.Map{
		"top_customers": topCustomers,
	}))
}

func (server *Server) getOrderDetail(c echo.Context) error {
	user := c.Get("user").(*jwt.Token).Claims.(*pkg.JwtCustomClaims)
	orderIdStr := c.Param("id")
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, pkg.ResponseError(pkg.ErrorBindingData, err))
	}
	order, err := repositories.GetOrderDetail(server.dbInstance, user.UserID, orderId, user.Permission)

	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, pkg.ResponseError(pkg.ErrorGetOrderDetail, err))
	}

	return c.JSON(http.StatusOK, pkg.ResponseSuccessWithData(pkg.InfoGetOrderDetailSuccess, echo.Map{
		"order": order,
	}))
}
