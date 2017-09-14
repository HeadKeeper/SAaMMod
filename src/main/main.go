package main

import (
	"fmt"
	"main/generator"
	"main/util"
	"github.com/aybabtme/uniplot/histogram"
	"os"
	"math"
)

/*
	Values for check 1 141 134456
*/

func main() {
	x0, a, m := util.GetInitValues()
	values := generator.LemerMethod(x0, a, m)
	bins := 20
	hist := histogram.Hist(bins, *values)
	maxWidth := 10
	err := histogram.Fprint(os.Stdout, hist, histogram.Linear(maxWidth))
	if (err != nil) {
		fmt.Println(err)
	}
	expValue, dispertion, standardDeviation := util.GetStatisticValues(*values)
	checkValue := util.GetIndirectionIndications(*values)
	period := util.GetPeriod(*values) - 1
	aperiod := util.GetAperiod(*values, period, a, m) + period
	fmt.Printf("M=%f --- D=%f --- SD=%f", expValue, dispertion, standardDeviation)
	fmt.Println()
	fmt.Printf("%f -> %f", checkValue, math.Pi / 4)
	fmt.Println()
	fmt.Printf("P=%d, L=%d", period, aperiod)
}