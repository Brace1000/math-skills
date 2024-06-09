package functions

import "sort"

func Median(data []float64) float64 {
	sort.Float64s(data)
	midpoint := len(data) / 2
	if len(data)%2 == 0 {
		return (data[midpoint-1] + data[midpoint]) / 2.0
	} else {
		return data[midpoint]
	}
}
