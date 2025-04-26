package api

import (
	"github.com/labstack/echo/v4"
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
	api := router.Group("/api")
	// auth api
	api.POST("/auth/login/local", server.localLogin)
	api.POST("/users", server.register) // Register
	api.GET("/auth/login/google", server.googleLogin)
	api.GET("/auth/login/google/callback", server.googleLoginCallback)

	// product api
	api.GET("/products", server.getAllProducts)
	api.GET("/products/detail/:product_id", server.getProductDetail)
	api.GET("/products/special-products", server.getSpecialProducts)
	server.router = router
}

func (server *Server) Start(address string) {
	server.router.Logger.Fatal(server.router.Start(address))
}
