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

	// auth api
	router.POST("/api/auth/login/local", server.localLogin)
	router.POST("/api/users", server.register) // Register
	router.GET("/api/auth/login/google", server.googleLogin)
	router.GET("/api/auth/login/google/callback", server.googleLoginCallback)

	// product api
	router.GET("/api/products", server.getAllProducts)
	router.GET("/api/products/special-products", server.getSpecialProducts)
	server.router = router
}

func (server *Server) Start(address string) {
	server.router.Logger.Fatal(server.router.Start(address))
}
