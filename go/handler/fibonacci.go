package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetFibonacci returns JSON GetFibonacci response
func (h *Handler) GetFibonacci(c echo.Context) error {
	number, _ := strconv.Atoi(c.Param("number"))
	fibonacci := h.fibonacci.GetFibonacci(number)
	return c.JSON(http.StatusOK, newFibonacciResponse(c, fibonacci))
}
