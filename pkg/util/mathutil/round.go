package mathutil

import (
	"fmt"
	"math"
	"strconv"
)

func Round(f float64, n int) float64 {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n)+"f", f)
	inst, _ := strconv.ParseFloat(floatStr, 64)
	return inst
}

func RoundByTrunc(f float64, n int) float64 {
	pow10n := math.Pow10(n)
	return math.Trunc(f*pow10n) / pow10n
}
