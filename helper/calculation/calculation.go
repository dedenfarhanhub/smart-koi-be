package calculation

import "math"

func FindMinAndMax(a []int64) (min int64, max int64) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func FindMinFloat(a float64, b float64) (min float64) {
	return math.Min(a, b)
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}