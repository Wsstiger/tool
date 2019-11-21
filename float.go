package tool

import (
	"fmt"
	"math"
	"strconv"
)

// 浮点型精度
func KeepPrecision(number float64, precision int) float64 {
	format := "%." + fmt.Sprintf("%v", precision) + "f"
	tmp := fmt.Sprintf(format, number)
	result, _ := strconv.ParseFloat(tmp, 64)
	return result
}

// 优先用这个
func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}
