package handler

import (
	"BE_Ecommerce/src/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type Server struct {
	dbInstance *gorm.DB
	router     *echo.Echo
}

func NewServer(dbInstance *gorm.DB) (*Server, error) {
	server := &Server{
		dbInstance: dbInstance,
	}
	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := echo.New()
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:*", "http://localhost:5173"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))
	api := router.Group("/api")
	// auth api
	api.POST("/auth/login/local", server.loginLocal)
	api.POST("/auth/register", server.register)
	// api.POST("/auth/login/refresh", server.loginRefresh)
	api.POST("/auth/login/google", server.loginGoogle)
	api.POST("/auth/login/facebook", server.loginFacebook)

	// product api
	api.GET("/products", server.getAllProducts)
	api.POST("/products/create", server.createProduct, middlewares.Authenticate(), middlewares.IsAdmin())
	api.GET("/products/detail/:product_id", server.getProductDetail)
	api.GET("/products/special-products", server.getSpecialProducts)
	api.PUT("/products/:product_id", server.updateProduct, middlewares.Authenticate(), middlewares.IsAdmin())
	api.DELETE("/products/:product_id", server.deleteProduct, middlewares.Authenticate(), middlewares.IsAdmin())
	api.GET("/products/statistic/category", server.getStatisticByCategory, middlewares.Authenticate(), middlewares.IsAdmin())
	api.GET("/products/statistic/manufacturer", server.getStatisticByManufacturer, middlewares.Authenticate(), middlewares.IsAdmin())

	// category api
	api.GET("/categories", server.getAllCategories)
	api.GET("/categories/:category_id", server.getCategoryById, middlewares.Authenticate(), middlewares.IsAdmin())
	api.POST("/categories", server.createCategory, middlewares.Authenticate(), middlewares.IsAdmin())
	api.PUT("/categories/:category_id", server.updateCategory, middlewares.Authenticate(), middlewares.IsAdmin())
	api.DELETE("/categories/:category_id", server.deleteCategory, middlewares.Authenticate(), middlewares.IsAdmin())
	api.GET("/categories/products/:id", server.getProductsByCategoryId)

	// manufacturer api
	api.GET("/manufacturers", server.getAllManufacturers)

	// user api
	api.GET("/users", server.getAllUsers, middlewares.Authenticate(), middlewares.IsAdmin())
	api.GET("/users/me", server.getUserMe, middlewares.Authenticate())
	api.GET("/users/:user_id", server.getUserById, middlewares.Authenticate(), middlewares.IsAdmin())
	api.GET("/users/statistics/new-users", server.getNewUsersStatistics, middlewares.Authenticate(), middlewares.IsAdmin())
	api.PUT("/users/reset-password", server.resetPassword, middlewares.Authenticate())
	api.PUT("/users", server.updateUser, middlewares.Authenticate())
	api.DELETE("/users/:user_id", server.deleteUser, middlewares.Authenticate(), middlewares.IsAdmin())

	// cart API
	api.GET("/carts", server.getCart, middlewares.Authenticate())
	api.POST("/carts", server.addToCart, middlewares.Authenticate())
	api.PUT("/carts/:product_id", server.updateCart, middlewares.Authenticate())
	api.POST("/carts/items", server.addItemsToCart, middlewares.Authenticate())
	api.DELETE("/carts/:product_id", server.deleteCartItem, middlewares.Authenticate())

	// order API
	api.POST("/orders", server.createOrder, middlewares.Authenticate())
	api.GET("/orders/history", server.getOrderHistoryOfUser, middlewares.Authenticate())
	api.GET("/orders/history/:user_id", server.getOrderHistoryOfUser, middlewares.Authenticate(), middlewares.IsAdmin())
	api.GET("/orders", server.getOrders, middlewares.Authenticate(), middlewares.IsAdmin())
	api.GET("/orders/statistics/revenue", server.getRevenueStatistics, middlewares.Authenticate(), middlewares.IsAdmin())
	api.GET("/orders/statistics/best-sellers", server.getBestSellersStatistics, middlewares.Authenticate(), middlewares.IsAdmin())
	api.GET("/orders/statistics/top-customers", server.getTopCustomersStatistics, middlewares.Authenticate(), middlewares.IsAdmin())
	api.GET("/orders/detail/:id", server.getOrderDetail, middlewares.Authenticate())
	server.router = router
}

func (server *Server) Start(address string) {
	server.router.Logger.Fatal(server.router.Start(address))
}
