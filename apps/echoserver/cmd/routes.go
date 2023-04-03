package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sharekhan/echoserver/internal/handlers"
)

// Skip Auth Middleware
func registerPublicRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world !!")
	})
}

// Go through Auth middlewares
// func registerPrivateRoutes(e *echo.Echo, s *handlers.Stats) {
// 	e.Use(s.Process)
// 	e.GET("/stats", s.Handle) // Endpoint to get stats
// 	groupMiddleware := e.Group("/users")
// 	groupMiddleware.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
// 		if username == "joe" && password == "secret" {
// 			return true, nil
// 		}
// 		return false, nil
// 	}))
// 	groupMiddleware.GET("/users/:id", handlers.GetUsers)
// 	// Server header
// 	e.Use(handlers.ServerHeader)

// 	// Handler
// 	e.GET("/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Hello, World!")
// 	})
// }

func registerRoutes(e *echo.Echo, s *handlers.Stats) {
	// registerPrivateRoutes(e, s)
	registerPublicRoutes(e)
}
