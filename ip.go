package tool

import (
	"fmt"
	"strconv"
	"strings"
)

// 获取下一个IP，只支持IPv4，不支持IPv6
func GetNextIP(currIP, endIP string) string {
	currItems, endItems := strings.Split(currIP, "."), strings.Split(endIP, ".")
	currIntItems, endIntItems := []int{0, 0, 0, 0}, []int{0, 0, 0, 0}
	for i := 0; i < 4; i++ {
		currIntItems[i], _ = strconv.Atoi(currItems[i])
		endIntItems[i], _ = strconv.Atoi(endItems[i])
	}
	nextIntItems := currIntItems
	nextIntItems[3] = nextIntItems[3] + 1
	for i := 3; i >= 0; i-- {
		if nextIntItems[i] > 255 {
			nextIntItems[i] = nextIntItems[i] - 255
			if i-1 >= 0 {
				nextIntItems[i-1] = nextIntItems[i-1] + 1
			} else {
				return ""
			}
		}
	}
	for i := 0; i < 4; i++ {
		if nextIntItems[i] < endIntItems[i] {
			return fmt.Sprintf("%v.%v.%v.%v", nextIntItems[0], nextIntItems[1], nextIntItems[2], nextIntItems[3])
		} else if nextIntItems[i] == endIntItems[i] {
			continue
		} else {
			return ""
		}
	}
	return fmt.Sprintf("%v.%v.%v.%v", nextIntItems[0], nextIntItems[1], nextIntItems[2], nextIntItems[3])
}
