package functions

import "math"

func Variance(numbers []float64) float64 {
	var squaredsum float64
	mean := 0.00
	n := float64(len(numbers))
	sum := 0.00
	for _, number := range numbers {
		sum += number
	}
	mean = float64(sum) / n

	for _, number := range numbers {
		deviation := float64(number) - mean
		squaredsum += (math.Pow(deviation, 2))
	}
	return squaredsum / n
}
