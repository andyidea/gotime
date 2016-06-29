package gotime

import (
	"time"
)

// 获取时间戳   return:1464059756
func GetTimeUnix(t time.Time) int64 {
	return t.Local().Unix()
}

// 获取当前时间的时间戳 return:1464059756
func GetNowTimeUnix() int64 {
	return GetTimeUnix(time.Now())
}

// 获取当日晚上24点（次日0点）的时间
func Get24time(t time.Time) time.Time {
	dateStr := TimeToDate(t.Add(time.Hour * 24))
	return DateStrToTime(dateStr)
}

// 获取当日晚上24点（次日0点）的时间戳
func Get24timeUnix(t time.Time) int64 {
	t24 := Get24time(t)
	return GetTimeUnix(t24)
}

// 获取今天晚上24点（次日0点）的时间
func GetToday24time() time.Time {
	return Get24time(time.Now())
}

// 获取今天晚上24点（次日0点）的时间戳
func GetToday24timeUnix() int64 {
	return Get24timeUnix(time.Now())
}

// 时间转换成日期字符串 (time.Time to "2006-01-02")
func TimeToDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// 获取当前时间的日期字符串（"2006-01-02"）
func GetNowDateStr() string {
	return TimeToDate(time.Now())
}

// 时间转换成日期+时间字符串 (time.Time to "2006-01-02 15:04:05")
func TimeToDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// 获取当前的时期+时间字符串
func GetNowStr() string {
	return TimeToDateTime(time.Now())
}

// 日期字符串转换成时间 ("2006-01-02" to time.Time)
func DateStrToTime(d string) time.Time {
	t, _ := time.ParseInLocation("2006-01-02", d, time.Local)
	return t
}

// 日期+时间字符串转换成时间 ("2006-01-02 15:04:05" to time.Time)
func DateTimeStrToTime(dt string) time.Time {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", dt, time.Local)
	return t
}

// 时间字符串转换成时间 ("15:04:05" to time.Time)
func TimeTodayStrToTime(dt string) time.Time {
	now := time.Now()
	strNowDate := TimeToDate(now)
	return DateTimeStrToTime(strNowDate + " " + dt)
}

// 是否是周末
func IsWeekend(t time.Time) bool {
	wd := t.Weekday()
	if wd == time.Sunday || wd == time.Saturday {
		return true
	}
	return false
}
