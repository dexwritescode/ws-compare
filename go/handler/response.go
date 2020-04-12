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

type fibonacciResponse struct {
	Input  int   `json:"input"`
	Output int64 `json:"output"`
}

func newHelloWorldResponse(c echo.Context, h *model.HelloWorld) *helloResponse {
	hr := new(helloResponse)
	hr.Message = h.Message
	return hr
}

func newGreetingResponse(c echo.Context, g *model.Greet) *greetingResponse {
	gr := new(greetingResponse)
	gr.Greeting = g.Greeting
	return gr
}

func newFibonacciResponse(c echo.Context, f *model.Fibonacci) *fibonacciResponse {
	fr := new(fibonacciResponse)
	fr.Input = f.Input
	fr.Output = f.Output
	return fr
}
