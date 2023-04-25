package utils

import "time"

// 工具包，我们自己的一些小工具可以写在这里
func UnixToTime(timestamp int) string {
	// 我们传入一个时间戳，就会返回我们一个年月日时分秒
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}
