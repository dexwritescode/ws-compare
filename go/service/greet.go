package service

import (
	"fmt"

	"github.com/dexterdarwich/ws-compare/model"
)

// GreetingService GreetingService struct
type GreetingService struct {
}

// NewGreetingService Returns a new GreetingService
func NewGreetingService() *GreetingService {
	return &GreetingService{}
}

// GetGreeting Returns the Greet object
func (gs *GreetingService) GetGreeting(name string) *model.Greet {
	var m model.Greet
	m.Greeting = fmt.Sprintf("Hello, %s!", name)
	return &m
}
