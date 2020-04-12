package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetGreeting returns JSON GetGreeting response
func (h *Handler) GetGreeting(c echo.Context) error {
	name := c.Param("name")
	greet := h.greeting.GetGreeting(name)
	return c.JSON(http.StatusOK, newGreetingResponse(c, greet))
}
