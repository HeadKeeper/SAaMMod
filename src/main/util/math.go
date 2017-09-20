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

func max(r1, r2 float64) float64 {
	if r1 >= r2 {
		return r1
	} else {
		return r2
	}
}

func getNewLinearValue(a, b, value float64) float64 {
	return a + (b - a) * value
}


/* 1 - Get Linear Random values*/
func GetLinear(values []float64, a, b float64) []float64 {
	var newValues []float64
	for index, _ := range values {
		value := getNewLinearValue(a, b, values[index])
		newValues = append(newValues, value)
	}
	return newValues
}

func GetLinearValues(a, b float64) (float64, float64) {
	m := (a + b) / 2
	dispersion := math.Pow((b - a), 2)/12
	return m, dispersion
}

/* 2 - Get Gauss Random values (EXP and STANDDEV is immutable) */
func GetGauss(expValue, standardDeviation float64, count int, values []float64) []float64 {
	var result []float64
	index := 0
	for index < len(values) - 1 {
		var sum float64
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

/* 3 - Get exponential values */
func GetExponential(lambda float64, values []float64) []float64 {
	var result []float64
	for index, _ := range values {
		element := -(1 / lambda) * math.Log(values[index])
		result = append(result, element)
	}
	return result
}

func GetExponentialValues(lambda float64) (float64, float64) {
	return (1 / lambda), (1 / math.Pow(lambda, 2))
}

/* 4 - Get Gamma values */
func GetGamma(lambda float64, nu int, values []float64) []float64 {
	var result []float64
	index := 0
	for index < len(values) - 1 {
		sum := 0.0
		for i := 0; i < nu; i++ {
			sum = sum + math.Log(values[index])
			index++
			if (index > len(values) - 1) {
				break;
			}
		}
		newValue := -(1 / lambda) * (sum)
		result = append(result, newValue)
	}
	return result
}

func GetGammaValues(lambda float64, nu float64) (float64, float64) {
	return (nu / lambda), (nu / math.Pow(nu, 2))
}

/* 5 - Get Triangle values */

func GetTriangle(a, b float64, values []float64) []float64 {
	var result []float64
	for i := 0; i < len(values) - 2; i++ {
		result = append(result, getNewLinearValue(a, b, max(values[i], values[i+1])))
	}
	return result
}

/* 6 - Get Simpson values */

func GetSimpson(a, b float64, values []float64) []float64 {
	firstValues := GetLinear(values, a/2, b/2)
	secondValues := GetLinear(values, a/2, b/2)
	var result []float64
	for i := 0; i < len(values) - 2; i+=2 {
		result = append(result, firstValues[i] + secondValues[i+1])
	}
	return result
}