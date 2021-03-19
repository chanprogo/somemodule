package randomnumber

import (
	"crypto/rand"
	"math/big"
	mathrand "math/rand"
	"time"
)

// Random generate string
func GetRandomNumString(n int) string {
	const alphanum = "012356789"

	var bytes = make([]byte, n)
	mathrand.Read(bytes)

	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}

	return string(bytes)
}

// 随机数组生成
func generateRandomNumber(start int, end int, count int) []int {

	if end < start || (end-start) < count {
		return nil
	}

	// 存放结果的slice
	nums := make([]int, 0)

	r := mathrand.New(mathrand.NewSource(time.Now().UnixNano()))

	for len(nums) < count {
		num := r.Intn((end - start))

		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}

func RandInt64(min, max int64) int64 {
	maxBigInt := big.NewInt(max)
	i, _ := rand.Int(rand.Reader, maxBigInt)
	iInt64 := i.Int64()
	if iInt64 < min {
		iInt64 = RandInt64(min, max) //应该用参数接一下
	}
	return iInt64
}
