package api

import (
	db "BE_Ecommerce/db/sqlc"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	store  db.Store
	router *echo.Echo
}

func NewServer(store db.Store) (*Server, error) {
	server := &Server{
		store: store,
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

	server.router = router
}

func (server *Server) Start(address string) {
	server.router.Logger.Fatal(server.router.Start(address))
}
