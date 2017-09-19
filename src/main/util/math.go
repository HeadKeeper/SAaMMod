package util

import "math"

func getExpectedValue(values []float64) float64 {
	var sum float64
	for _, element := range values {
		sum = sum + element
	}
	return sum / float64(len(values))
}

func getDispersion(values []float64, expectedValue float64) float64 {
	var sum float64
	for _, element := range values {
		sum = sum + math.Pow(element - expectedValue, 2)
	}
	return sum / float64(len(values))
}

func getStandardDeviation(dispersion float64) float64 {
	return math.Sqrt(dispersion)
}

func GetStatisticValues(values []float64) (float64, float64, float64) {
	expValue := getExpectedValue(values)
	dispersion := getDispersion(values, expValue)
	standardDeviation := getStandardDeviation(dispersion)
	return expValue, dispersion, standardDeviation
}


/* 1 - Get Linear Random values*/
func GetLinear(values []float64, a, b float64) []float64 {
	var newValues []float64
	for index, _ := range values {
		value := a + (b - a) * values[index]
		newValues = append(newValues, value)
	}
	return newValues
}

func GetLinearValues(a, b float64) (float64, float64) {
	m := (a + b) / 2
	dispersion := math.Pow((b - a), 2)/12
	return m, dispersion
}

/* 2 - Get Gauss Random values */
func GetGauss(expValue, standardDeviation float64, count int, values []float64) []float64 {
	var result []float64
	index := 0
	for index < len(values) - 1 {
		sum := 0.0
		for i := 0; i < count; i++ {
			sum = sum + values[index]
			index++
			if (index > len(values) - 1) {
				break;
			}
		}
		newValue := expValue + standardDeviation * math.Sqrt(12 / float64(count)) * (sum - float64(count) / 2)
		result = append(result, newValue)
	}
	return result
}