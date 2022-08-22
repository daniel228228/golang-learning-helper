package helper

import (
	"math"
)

func phi_f(x, precision float64) (result float64) {
	for i := float64(0); i < math.Abs(x); i += precision {
		result += precision * (math.Exp(-0.5*i*i) + math.Exp(-0.5*(i+precision)*(i+precision))) / 2.0
	}

	result /= math.Sqrt(2 * math.Pi)

	if x < 0 {
		result *= -1
	}

	return
}

func phi_d(x float64) float64 {
	return math.Exp(-0.5*x*x) / math.Sqrt(2*math.Pi)
}

// GetValueByX takes the value of the argument X and precision and returns the value of the Laplace function with given X and precision
func GetValueByX(x, precision float64) float64 {
	return phi_f(x, precision)
}

// GetXByValue takes the interval [a, b], the value of the Laplace function and the precision and returns the value of the argument X and the number of iterations of the X search
func GetXByValue(a, b, value, precision float64) (float64, uint) {
	result := float64(0.5)
	result_prev := float64(0)
	iterations := uint(0)

	for math.Abs(result-result_prev) > precision {
		result_prev = result
		result = result - (phi_f(result, precision)-value)/phi_d(result)
		iterations++
	}

	return result, iterations
}
