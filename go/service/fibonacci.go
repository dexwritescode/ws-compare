package service

import (
	"github.com/dexterdarwich/ws-compare/model"
)

// FibonacciService FibonacciService struct
type FibonacciService struct {
}

// NewFibonacciService Returns a new FibonacciService
func NewFibonacciService() *FibonacciService {
	return &FibonacciService{}
}

// GetFibonacci Returns the Fibonacci object
func (fs *FibonacciService) GetFibonacci(number int) *model.Fibonacci {
	var f model.Fibonacci
	f.Input = number
	f.Output = fibonacci(number)
	return &f
}

func fibonacci(number int) int64 {
	if number <= 1 {
		return int64(number)
	}
	return fibonacci(number-1) + fibonacci(number-2)
}
