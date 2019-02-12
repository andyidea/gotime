// gotime_test.go

package gotime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetTimeUnix(t *testing.T) {
	now := time.Now()
	unix := GetTimeUnix(now)
	assert.Equal(t, unix, now.Unix())
}

func TestGetNowTimeUnix(t *testing.T) {
	unix := GetNowTimeUnix()
	assert.NotZero(t, unix)
}

func TestGet24time(t *testing.T) {
	now := time.Now()
	time24 := Get24Time(now)
	assert.NotZero(t, time24.Unix())
}

func TestGet24timeUnix(t *testing.T) {
	now := time.Now()
	time24 := Get24Time(now)
	unix := Get24TimeUnix(now)
	assert.NotZero(t, unix)
	assert.Equal(t, unix, time24.Unix())
}

func TestGetToday24time(t *testing.T) {
	time24 := GetToday24Time()
	assert.NotZero(t, time24.Unix())
}

func TestGetToday24timeUnix(t *testing.T) {
	unix := GetToday24TimeUnix()
	assert.NotZero(t, unix)
}

func TestTimeToDate(t *testing.T) {
	now := time.Now()
	date := TimeToDate(now)
	assert.NotZero(t, date)
}

func TestGetNowDateStr(t *testing.T) {
	date := GetNowDateStr()
	assert.NotZero(t, date)
}

func TestTimeToDateTime(t *testing.T) {
	now := time.Now()
	dateTime := TimeToDateTime(now)
	assert.NotZero(t, dateTime)
}

func TestGetNowStr(t *testing.T) {
	nowStr := GetNowStr()
	assert.NotZero(t, nowStr)
}

func TestDateStrToTime(t *testing.T) {
	dateStr := "2016-01-01"
	ti := DateStrToTime(dateStr)
	assert.NotZero(t, ti.Unix())
}

func TestDateTimeStrToTime(t *testing.T) {
	dateTimeStr := "2016-01-01 10:01:01"
	ti := DateTimeStrToTime(dateTimeStr)
	assert.NotZero(t, ti.Unix())
}

func TestTimeTodayStrToTime(t *testing.T) {
	timeStr := "10:01:01"
	ti := TimeTodayStrToTime(timeStr)
	assert.NotZero(t, ti.Unix())
}

func TestIsWeekend(t *testing.T) {
	now := time.Now()
	IsWeekend(now)
}
