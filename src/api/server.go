package api

import (
	"net/http"

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
	router.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"code":    http.StatusOK,
			"message": "Hello World",
		})
	})
	router.POST("/api/auth/login/local", server.localLogin)
	router.GET("/api/auth/login/google", server.googleLogin)
	router.GET("/api/auth/login/google/callback", server.googleLoginCallback)
	router.POST("/api/users", server.register) // Register

	// product api

	// product := router.Group("/api/products")

	// product.GET("", server.getAllProducts)
	// product.GET("/special-products", server.getSpecialProducts)
	server.router = router
}

func (server *Server) Start(address string) {
	server.router.Logger.Fatal(server.router.Start(address))
}
