package domain

import (
	"testing"
)

func TestBMICreation(t *testing.T) {
	bmi := BMI{
		Height: 1.85,
		Weight: 90,
	}

	if bmi.Height != 1.85 || bmi.Weight != 90 {
		t.Errorf("BMI not correctly created: %+v", bmi)
	}
}
