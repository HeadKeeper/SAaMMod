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


/* 1 - Get Linear Random values*/
func GetLinear(values []float64, a, b float64) []float64 {
	var newValues []float64
	for index, _ := range values {
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