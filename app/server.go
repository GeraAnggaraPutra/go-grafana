package app

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	zlog "github.com/rs/zerolog/log"
)

func StartServer() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(HttpLoggerMiddleware())
	e.Use(echoprometheus.NewMiddleware(os.Getenv("GRAFANA_NAME")))
	e.Use(TrackCustomMetrics)

	e.GET("/metrics", echoprometheus.NewHandler())
	e.GET("/health", healthHandler)
	e.GET("/users", usersHandler)
	e.GET("/error", errorHandler)

	if err := e.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))); err != nil {
		log.Fatal(err)
	}
}

func healthHandler(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Up and running!",
	})
}

func usersHandler(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello from /users",
	})
}

func errorHandler(c *echo.Context) error {
	err := errors.New("manual trigger error for testing")

	zlog.Error().Err(err).
		Str("method", c.Request().Method).
		Str("path", c.Path()).
		Int("status", http.StatusInternalServerError).
		Msg("An error occurred in /error endpoint")

	return c.JSON(http.StatusInternalServerError, map[string]string{
		"error": "Internal Server Error",
	})
}
