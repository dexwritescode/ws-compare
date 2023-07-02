package service

import (
	"github.com/dexterdarwich/ws-compare/model"
)

// HelloWorldService HelloWorldService struct
type HelloWorldService struct {
}

// NewHelloWorldService Returns a new HelloService
func NewHelloWorldService() *HelloWorldService {
	return &HelloWorldService{}
}

// GetHelloWorld Returns the Hello object
func (hs *HelloWorldService) GetHelloWorld() *model.HelloWorld {
	var m model.HelloWorld
	m.Message = "Hello, World!"
	return &m
}
