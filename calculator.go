package calculator

import (
	"errors"
	"math"
)

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiplication(a, b float64) float64 {
	return a * b
}

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero not allowed")
	}
	return (a / b), nil
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("sqrt by negative numbers not allowed")
	}
	return math.Sqrt(a), nil
}
