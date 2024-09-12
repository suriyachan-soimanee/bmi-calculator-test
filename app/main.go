package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"

	"github.com/suriyachan-soimanee/bmi-calculator/bmi"
	"github.com/suriyachan-soimanee/bmi-calculator/internal/rest"
)

const (
	defaultTimeout = 30
	defaultAddress = ":9091"
)

func main() {
	// prepare echo
	e := echo.New()

	// Build service Layer
	bmi := bmi.NewBMIService()
	rest.NewBMIHandler(e, bmi)

	// Start Server
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = defaultAddress
	}
	log.Fatal(e.Start(address)) //nolint
}
