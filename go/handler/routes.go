package handler

import (
	"github.com/labstack/echo/v4"
)

// Register Registers the routes
func (h *Handler) Register(e *echo.Echo) {
	e.GET("/", h.GetHelloWorld)
}
