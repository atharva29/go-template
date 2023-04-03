package main

import (
	"github.com/atharva29/echoserver/internal/handlers"
	"github.com/atharva29/go-template/pkg/kafka"
	"github.com/labstack/echo/v4"
)

func main() {
	s := handlers.NewStats()
	kafka.KafkaInitialize()
	e := echo.New()
	registerRoutes(e, s)
	e.Logger.Fatal(e.Start(":1323"))
}
