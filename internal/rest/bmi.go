package rest

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BMIService interface {
	CalculateBMI(height, weight float64) float64
}

type BMIHandler struct {
	Service BMIService
}

func NewBMIHandler(e *echo.Echo, u BMIService) {
	handler := &BMIHandler{
		Service: u,
	}

	e.GET("/bmi", handler.CalculateBMI)
}

func (h *BMIHandler) CalculateBMI(c echo.Context) error {
	heightStr := c.QueryParam("height")
	weightStr := c.QueryParam("weight")

	height, err := strconv.ParseFloat(heightStr, 64)
	if err != nil || height <= 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid height"})
	}

	weight, err := strconv.ParseFloat(weightStr, 64)
	if err != nil || weight <= 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid weight"})
	}

	bmi := h.Service.CalculateBMI(height, weight)

	return c.JSON(http.StatusOK, echo.Map{"BMI": bmi})
}
