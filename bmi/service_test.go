package bmi_test

import (
	"testing"

	"github.com/suriyachan-soimanee/bmi-calculator/bmi"
)

func TestCalculateBMI(t *testing.T) {
	bmiService := bmi.NewBMIService()

	// Test case: Normal height and weight
	height := 1.55
	weight := 47.0
	expectedBMI := weight / (height * height)

	bmi := bmiService.CalculateBMI(height, weight)
	if bmi != expectedBMI {
		t.Errorf("Expected BMI: %.2f, but got %.2f", expectedBMI, bmi)
	}

	// Test case: Height is zero
	bmi = bmiService.CalculateBMI(0, weight)
	if bmi != 0 {
		t.Errorf("Expected BMI: 0, but got %.2f", bmi)
	}
}
