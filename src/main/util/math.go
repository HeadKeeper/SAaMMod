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
	for index := range values {
		value := a + (b - a)*values[index]
		newValues = append(newValues, value)
	}
	return newValues
}

/* 2 - Get Gauss Random values */
func GetGauss(expValue, standardDeviation float64, count int, values []float64) float64 {
	var sum float64
	for _, element := range values {
		sum = sum + element - float64(len(values))/2
	}
	result := sum * math.Sqrt(12 / float64(len(values))) * standardDeviation
	return result + expValue
}