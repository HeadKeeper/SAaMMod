package generator

import (
)

func LemerMethod(x0, a, m, n int) *[]float64 {
	var elementsCount = n
	var randomValues []float64
	previewX := getNewValue(x0, a, m)
	for i := 0; i < elementsCount; i++ {
		randomValues = append(randomValues, float64(previewX) / float64(m))
		previewX = getNewValue(previewX, a, m)
	}
	return &randomValues
}

func getNewValue(previewValue, a, m int) int {
	return (a * previewValue) % m
}
