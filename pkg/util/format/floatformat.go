package format

import (
	"fmt"
)

// 成交量格式化
func VolumeFormat(f float64) string {
	switch {
	case f < 0.000001:
		return fmt.Sprintf("%.8f", f)
	case f < 0.0001:
		return fmt.Sprintf("%.6f", f)
	case f < 0.01:
		return fmt.Sprintf("%.4f", f)
	case f < 1:
		return fmt.Sprintf("%.2f", f)
	case f < 100:
		return fmt.Sprintf("%.1f", f)
	default:
		return fmt.Sprintf("%.0f", f)

	}
}

// 人民币价格格式化
func RmbPriceFormat(f float64) string {
	switch {
	case f < 0.000001:
		return fmt.Sprintf("%.8f", f)
	case f < 0.0001:
		return fmt.Sprintf("%.6f", f)
	case f < 0.01:
		return fmt.Sprintf("%.4f", f)
	default:
		return fmt.Sprintf("%.2f", f)
	}
}

// 价格格式化
func PriceFormat(f float64) string {
	switch {
	case f < 0.00001:
		return fmt.Sprintf("%.10f", f)
	case f < 0.001:
		return fmt.Sprintf("%.8f", f)
	case f < 0.1:
		return fmt.Sprintf("%.6f", f)
	case f < 1000:
		return fmt.Sprintf("%.4f", f)
	default:
		return fmt.Sprintf("%.2f", f)
	}
}

// 盘口深度数量格式化，最多保留4位数字
func DepthNumFormat(f float64) string {
	switch {
	case f >= 1000000:
		tmp := fmt.Sprintf("%.3f", f/1000000)
		return tmp[0:5] + "M"
	case f >= 1000:
		tmp := fmt.Sprintf("%.3f", f/1000)
		return tmp[0:5] + "K"
	default:
		tmp := fmt.Sprintf("%.3f", f)
		return tmp[0:5]
	}
}
