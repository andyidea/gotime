// time.go

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
	date := TimeToDate(t.Add(time.Hour * 24))
	return DateStrToTime(date)
}

// 获取当日晚上24点（次日0点）的时间戳
func Get24timeUnix(t time.Time) int64 {
	date := TimeToDate(t.Add(time.Hour * 24))
	t24 := DateStrToTime(date)
	return GetTimeUnix(t24)
}

// 获取今天晚上24点（次日0点）的时间戳
func GetToday24timeUnix() int64 {
	return Get24timeUnix(time.Now())
}

// 时间转换成日期字符串 (time.Time to "2006-01-02")
func TimeToDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// 获取当前的日期字符串
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
