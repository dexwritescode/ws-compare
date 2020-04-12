package handler

import (
	"github.com/dexterdarwich/ws-compare/service"
)

// Handler The handler of the web service
type Handler struct {
	helloWorld *service.HelloWorldService
}

// NewHandler Create a new handler
func NewHandler(h *service.HelloWorldService) *Handler {
	return &Handler{
		helloWorld: h,
	}
}
