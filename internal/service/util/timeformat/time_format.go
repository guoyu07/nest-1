package timeformat

import (
	"time"
	"strconv"
	"fmt"
)

const (
	TIME_FORMAT = "20060102150405"
	DATE_FORMAT = "20060102"
	MONTH_FORMAT = "200601"
)

func FormatTime(t time.Time) int64 {
	result, _ := strconv.ParseInt(t.Format(TIME_FORMAT), 10, 64)
	return result
}

func FormatDate(t time.Time) int64 {
	result, _ := strconv.ParseInt(t.Format(DATE_FORMAT), 10, 64)
	return result
}

func FormatMonth(t time.Time) int64 {
	result, _ := strconv.ParseInt(t.Format(MONTH_FORMAT), 10, 64)
	return result
}

func FormatTimeNow() int64 {
	return FormatTime(time.Now())
}

func FormatDateNow() int64 {
	return FormatDate(time.Now())
}

func FormatMonthNow() int64 {
	return FormatMonth(time.Now())
}

func ParseTime(n int64) *time.Time {
	t, err := time.ParseInLocation(TIME_FORMAT, fmt.Sprintf("%v", n), time.Local)
	if err != nil {
		return nil
	}
	return &t
}

func ParseDate(n int64) *time.Time {
	t, err := time.ParseInLocation(DATE_FORMAT, fmt.Sprintf("%v", n), time.Local)
	if err != nil {
		return nil
	}
	return &t
}

func ParseMonth(n int64) *time.Time {
	t, err := time.ParseInLocation(MONTH_FORMAT, fmt.Sprintf("%v", n), time.Local)
	if err != nil {
		return nil
	}
	return &t
}