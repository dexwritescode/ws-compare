package handler

import (
	"github.com/dexterdarwich/ws-compare/model"
	"github.com/labstack/echo/v4"
)

type helloResponse struct {
	Message string `json:"message"`
}

func newHelloWorldResponse(c echo.Context, h *model.HelloWorld) *helloResponse {
	hr := new(helloResponse)
	hr.Message = h.Message
	return hr
}
