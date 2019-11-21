package tool

import (
	"fmt"
	"time"
)

func CurrentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func CurrentTimeSecond() int64 {
	return time.Now().Unix()
}

func CurrentTimeNanoSecond() int64 {
	return time.Now().UnixNano()
}

func Unixtime2Str(unixtime int64) string {
	tm := time.Unix(unixtime, 0)
	return tm.Format("2006-01-02 15:04:05")
}

func Str2Time(format, src string) (time.Time, error) {
	return time.ParseInLocation(format, src, time.Local)
}

func CurrentTimeAdd2Str(t time.Duration) string {
	return time.Now().Add(t).Format("2006-01-02 15:04:05")
}

func Second2Human(freeSec int64) string {
	days := freeSec / 86400
	freeSec = freeSec - days*86400
	hours := freeSec / 3600
	freeSec = freeSec - hours*3600
	mins := freeSec / 60
	secs := freeSec - mins*60

	s := ""
	if days > 0 {
		s = s + fmt.Sprintf("%v天", days)
	}
	if hours > 0 {
		s = s + fmt.Sprintf("%v小时", hours)
	}
	if mins > 0 {
		s = s + fmt.Sprintf("%v分钟", mins)
	}
	if secs > 0 {
		s = s + fmt.Sprintf("%v秒", secs)
	}
	return s
}
