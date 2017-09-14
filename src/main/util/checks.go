package util

import (
	"math"
	"main/generator"
)

func getPairsCount(values []float64) int {
	var counter int
	for i := 0; i < len(values); i += 2 {
		value := math.Pow(values[i], 2) + math.Pow(values[i+1], 2)
		if (value < 1) {
			counter++
		}
	}
	return counter
}

func GetIndirectionIndications(values []float64) float64 {
	k := getPairsCount(values)
	return float64(2 * k) / math.Pow10(6)
}

func GetPeriod(values []float64) int {
	last := values[len(values) - 1]
	for i := len(values) - 2; i > 0; i-- {
		if (values[i] == last) {
			return len(values) - i
		}
	}
	return 0
}

func GetAperiod(values []float64, period, a, m int) int {
	newValues := *generator.LemerMethod(period, a, m)
	for i := 0; i < len(values) - 1 - period; i++ {
		if (newValues[i] == newValues[i + period]) {
			return i;
		}
	}
	return 0
}