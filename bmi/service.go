package bmi

import (
	"github.com/suriyachan-soimanee/bmi-calculator/domain"
)

type BMIService interface {
	CalculateBMI(height, weight float64) float64
}

type bmiService struct{}

func NewBMIService() BMIService {
	return &bmiService{}
}

func (u *bmiService) CalculateBMI(height, weight float64) float64 {

	bmi := domain.BMI{
		Height: height,
		Weight: weight,
	}

	// calculate BMI
	if bmi.Height == 0 {
		return 0
	}
	return bmi.Weight / (bmi.Height * bmi.Height)
}
