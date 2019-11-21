package tool

import (
	"fmt"
	"strconv"
	"time"
)

// 字节转兆字节
func Byte2MByte(num, round int) float64 {
	mb := float64(num) / 1024 / 1024
	formatStr := "%." + strconv.Itoa(round) + "f"
	mb, _ = strconv.ParseFloat(fmt.Sprintf(formatStr, mb), 64)
	return mb
}

// 字节转兆比特
func Byet2Mbit(num, round int) float64 {
	mb := float64(num) / 1024 / 1024 * 8
	formatStr := "%." + strconv.Itoa(round) + "f"
	mb, _ = strconv.ParseFloat(fmt.Sprintf(formatStr, mb), 64)
	return mb
}

// 统计函数执行时长
func CountTimeUsed() func() {
	start := time.Now()
	return func() {
		fmt.Printf("use time: %v", time.Since(start))
	}
}
