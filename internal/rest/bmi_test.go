package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockBMIService struct {
	mock.Mock
}

func (m *MockBMIService) CalculateBMI(height, weight float64) float64 {
	args := m.Called(height, weight)
	return args.Get(0).(float64)
}

func TestCalculateBMI(t *testing.T) {
	e := echo.New()

	// Create mock
	mockService := new(MockBMIService)
	handler := &BMIHandler{Service: mockService}

	// Test case: Normal height and weight
	req := httptest.NewRequest(http.MethodGet, "/bmi?height=1.85&weight=75", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService.On("CalculateBMI", 1.85, 75.0).Return(21.913805697589478)

	// Invoke the handler
	if assert.NoError(t, handler.CalculateBMI(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"BMI": 21.913805697589478}`, rec.Body.String())
	}

	// Test case: Invalid height
	req = httptest.NewRequest(http.MethodGet, "/bmi?height=0&weight=75", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	if assert.NoError(t, handler.CalculateBMI(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, `{"error": "Invalid height"}`, rec.Body.String())
	}
}
