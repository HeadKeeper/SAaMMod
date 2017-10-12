package main

import (
	"fmt"
	"main/generator"
	"main/util"
	"math"
	"main/histogram"
	"main/labs/lab3"
)

/*
	Values for check:
	1 141 134456 1000000 - P < 50000
	7 17000 160001 1000000 - P > 50000
	7 17000 160001 1000000 1 7 1 1 123 55 6
	7 17 121 1000000
*/

func main() {
	thirdLabMain()
}

func secondLabMain() {
	x0, a, m, n, left, right, lambda, nu, gaussExpValue, gaussStandDev, count := util.GetInitValues()
	//fmt.Printf("%v %v %v %v %v %v %v %v %v %v %v", x0, a, m, n, left, right, lambda, nu, gaussExpValue, gaussStandDev, count)
	values := generator.LemerMethod(x0, a, m, n)
	expValue, expDispertion, standardDeviation := util.GetStatisticValues(*values)
	/*gaussValues := util.GetTriangle(1, 7, *values)//util.GetGauss(expValue, standardDeviation, 6, *values)
	histogram.DrawHistogram(gaussValues)*/
	fmt.Println("---------------LINEAR---------------")
	linear := util.GetLinear(*values, left, right)
	histogram.DrawHistogram(linear)
	expValue, expDispertion, standardDeviation = util.GetStatisticValues(linear)
	fmt.Printf("M=%f --- D=%f --- SD=%f\n", expValue, expDispertion, standardDeviation)
	fmt.Println("---------------GAUSS---------------")
	gauss := util.GetGauss(gaussExpValue, gaussStandDev, count, *values)
	histogram.DrawHistogram(gauss)
	expValue, expDispertion, standardDeviation = util.GetStatisticValues(gauss)
	fmt.Printf("M=%f --- D=%f --- SD=%f\n", expValue, expDispertion, standardDeviation)
	fmt.Println("---------------EXPONENTIAL---------------")
	exponential := util.GetExponential(lambda, *values)
	histogram.DrawHistogram(exponential)
	expValue, expDispertion, standardDeviation = util.GetStatisticValues(exponential)
	fmt.Printf("M=%f --- D=%f --- SD=%f\n", expValue, expDispertion, standardDeviation)
	fmt.Println("---------------GAMMA---------------")
	gamma := util.GetGamma(lambda, int(nu), *values)
	histogram.DrawHistogram(gamma)
	expValue, expDispertion, standardDeviation = util.GetStatisticValues(gamma)
	fmt.Printf("M=%f --- D=%f --- SD=%f\n", expValue, expDispertion, standardDeviation)
	fmt.Println("---------------TRIANGLE---------------")
	triangle := util.GetTriangle(left, right, *values)
	histogram.DrawHistogram(triangle)
	expValue, expDispertion, standardDeviation = util.GetStatisticValues(triangle)
	fmt.Printf("M=%f --- D=%f --- SD=%f\n", expValue, expDispertion, standardDeviation)
	fmt.Println("---------------SIMPSON---------------")
	simpson := util.GetSimpson(left, right, *values)
	histogram.DrawHistogram(simpson)
	expValue, expDispertion, standardDeviation = util.GetStatisticValues(simpson)
	fmt.Printf("M=%f --- D=%f --- SD=%f\n", expValue, expDispertion, standardDeviation)
}

func firstLabMain() {
	x0, a, m, n, _, _, _, _, _, _, _ := util.GetInitValues()
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

func thirdLabMain() {
	//lab3.Setup()
	lab3.Perform()
}