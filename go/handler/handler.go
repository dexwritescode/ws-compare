package handler

import (
	"github.com/dexterdarwich/ws-compare/service"
)

// Handler The handler of the web service
type Handler struct {
	helloWorld *service.HelloWorldService
	greeting   *service.GreetingService
	fibonacci  *service.FibonacciService
}

// NewHandler Create a new handler
func NewHandler(h *service.HelloWorldService, g *service.GreetingService, f *service.FibonacciService) *Handler {
	return &Handler{
		helloWorld: h,
		greeting:   g,
		fibonacci:  f,
	}
}
