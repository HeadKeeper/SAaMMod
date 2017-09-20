package util

import "fmt"

func GetInitValues() (int, int, int, int, float64, float64, float64, int, float64, float64, int) {
	var x0, a, m, n, count, nu int
	var left, right, expValue, standDev, lambda  float64
	fmt.Println("X0 A M N [A,B] LAMBDA Nu EXPVAL STANDDEV GAUSS_COUNT:")
	fmt.Scanf("%d %d %d %d %f %f %f %d %f %f %d", &x0, &a, &m, &n, &left, &right, &lambda, &nu, &expValue, &standDev, &count)
	return x0, a, m, n, left, right, lambda, nu, expValue, standDev, count
}
