package tool

import (
	"math/rand"
	"strconv"
	"time"
)

// 生成n位code码
func CreateCode(n int) string {
	// 撒一颗随机种子
	rand.Seed(time.Now().UnixNano())
	code := ""
	for i := 0; i < n; i++ {
		code = code + strconv.Itoa(rand.Intn(10))
	}

	return code
}
