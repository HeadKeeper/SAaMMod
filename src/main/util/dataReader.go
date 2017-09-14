package util

import "fmt"

func GetInitValues() (int, int, int) {
	var x0, a, m int
	fmt.Println("X0 A M:")
	fmt.Scanf("%d %d %d", &x0, &a, &m)
	return x0, a, m
}
