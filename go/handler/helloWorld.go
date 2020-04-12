package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetHelloWorld returns JSON GetHelloWorld response
func (h *Handler) GetHelloWorld(c echo.Context) error {
	hello := h.helloWorld.GetHelloWorld()
	return c.JSON(http.StatusOK, newHelloWorldResponse(c, hello))
}
