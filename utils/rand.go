package utils

import (
	"math/rand"
	"time"
)

var r *rand.Rand

// GetRangeNum 生成随机整数 digit：位数
func GetRangeNum(digit int) int {
	var max, min int = 1, 1
	if digit > 0 {
		for i := 0; i < digit; i++ {
			max = max * 10
		}
		for i := 0; i < digit-1; i++ {
			min = min * 10
		}
	}
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
