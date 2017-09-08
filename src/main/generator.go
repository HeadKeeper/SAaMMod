package main

import (
	"fmt"
)

func LemerMethod() int {
	fmt.Println("Enter init value:")
	var initValue int;
	fmt.Scanf("%d", &initValue);

	return initValue;
}
