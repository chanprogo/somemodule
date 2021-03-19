package mathutil

import (
	"math/big"
	"strconv"
)

func HexToInt64(num string) int64 {
	v := num[2:]
	s, err := strconv.ParseInt(v, 16, 64)
	if err == nil {
		return s
	}
	return 0
}

func HexToTenStr(value string) string {
	bignumber := big.NewInt(0)
	bignumber.SetString(value, 0)
	return bignumber.String()
}
