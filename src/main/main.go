package main

import (
	"fmt"
	"main/generator"
	"main/util"
	"math"
	"main/histogram"
)

/*
	Values for check:
	1 141 134456 - P < 50000
	7 17000 160001 - P > 50000
	7 17 121
*/

func main() {

}

func secondLabMain() {

}

func firstLabMain() {
	x0, a, m, n := util.GetInitValues()
	values := generator.LemerMethod(x0, a, m, n)
	err := histogram.DrawHistogram(*values)
	if err != nil {
		fmt.Println(err)
	}
	expValue, dispersion, standardDeviation := util.GetStatisticValues(*values)
	checkValue := util.GetIndirectionIndications(*values)
	period := util.GetPeriod(*values) - 1
	aperiodic := util.GetAperiodic(*values, period, a, m) + period
	fmt.Printf("M=%f --- D=%f --- SD=%f", expValue, dispersion, standardDeviation)
	fmt.Println()
	fmt.Printf("%f -> %f", checkValue, math.Pi / 4)
	fmt.Println()
	fmt.Printf("P=%d, L=%d", period, aperiodic)
	fmt.Println()
	//fmt.Println(util.GetGauss(expValue, standardDeviation, *values))
}