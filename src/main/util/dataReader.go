package util

import "fmt"

func GetInitValues() (int, int, int, int) {
	var x0, a, m, n int
	fmt.Println("X0 A M N:")
	fmt.Scanf("%d %d %d %d", &x0, &a, &m, &n)
	return x0, a, m, n
}
