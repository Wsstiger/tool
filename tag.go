package tool

import (
	"time"
)

func CreateTag() int64 {
	return time.Now().UnixNano()
}
