package main

import (
	"github.com/dexterdarwich/ws-compare/handler"
	"github.com/dexterdarwich/ws-compare/service"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	helloWorld := service.NewHelloWorldService()
	greeting := service.NewGreetingService()
	handler := handler.NewHandler(helloWorld, greeting)
	handler.Register(e)
	e.Logger.Fatal(e.Start(":8080"))
}
