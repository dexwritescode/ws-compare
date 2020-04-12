package handler

import (
	"github.com/dexterdarwich/ws-compare/model"
	"github.com/labstack/echo/v4"
)

type helloResponse struct {
	Message string `json:"message"`
}

type greetingResponse struct {
	Greeting string `json:"greeting"`
}

func newHelloWorldResponse(c echo.Context, h *model.HelloWorld) *helloResponse {
	hr := new(helloResponse)
	hr.Message = h.Message
	return hr
}

func newGreetingResponse(c echo.Context, h *model.Greet) *greetingResponse {
	gr := new(greetingResponse)
	gr.Greeting = h.Greeting
	return gr
}
