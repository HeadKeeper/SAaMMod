package util

import "math"

func getExpectedValue(values []float64) float64 {
	var sum float64
	for _, element := range values {
		sum = sum + element
	}
	return sum / float64(len(values))
}

func getDispertion(values []float64, expectedValue float64) float64 {
	var sum float64
	for _, element := range values {
		sum = sum + math.Pow((element - expectedValue), 2)
	}
	return sum / float64(len(values))
}

func getStandartDeviation(dispertion float64) float64 {
	return math.Sqrt(dispertion)
}

func GetStatisticValues(values []float64) (float64, float64, float64) {
	expValue := getExpectedValue(values)
	dispertion := getDispertion(values, expValue)
	standartDeviation := getStandartDeviation(dispertion)
	return expValue, dispertion, standartDeviation
}