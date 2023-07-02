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
	f.Output = fibonacci(int64(number))
	return &f
}

func fibonacci(number int64) int64 {
	if number <= 1 {
		return number
	}
	return fibonacci(number-1) + fibonacci(number-2)
}
