package util

import (
	"time"
)

const (
	DateTimeFormatString = "2000-01-01 23:59:59"
	DateFormatString     = "2000-01-01"
)

func BeginOfDay(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, tm.Location())
}

func BeginOfYesterday(tm time.Time) time.Time {
	return BeginOfDay(tm.Add(-24 * time.Hour))
}

func BeginOfTomorrow(tm time.Time) time.Time {
	return BeginOfDay(tm.Add(24 * time.Hour))
}

func DurationUntilTomorrow(tm time.Time) time.Duration {
	tom := BeginOfDay(tm.Add(24 * time.Hour))
	return tom.Sub(tm)
}

func NextMonday(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month(), tm.Day()+(7-int(tm.Weekday()))%7+1, 0, 0, 0, 0, time.Local)
}

func FirstDayOfLastMonth() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).AddDate(0, -1, 0)
}

func LastDayOfLastMonth() time.Time {
	now := time.Now()
	firstDayOfThisMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	//往前挪1秒
	return firstDayOfThisMonth.Add(-time.Second)
}

func DayDiff(before, after time.Time) int64 {
	hours := BeginOfDay(after).Sub(BeginOfDay(before)).Hours()
	return int64(hours / 24)
}

func MonthDiff(before, after time.Time) int64 {
	beforeYear := before.Year()
	beforeMonth := int(before.Month())
	afterYear := after.Year()
	afterMonth := int(after.Month())

	return int64((afterYear-beforeYear)*12 + (afterMonth - beforeMonth))

}

var weekdayMap = map[int]string{
	1: "星期一",
	2: "星期二",
	3: "星期三",
	4: "星期四",
	5: "星期五",
	6: "星期六",
	0: "星期日",
}

func GetChineseWeekday(ts time.Time) string {
	return weekdayMap[int(ts.Weekday())]
}

func RemainingSecondsToday() time.Duration {
	now := time.Now()
	return BeginOfTomorrow(now).Sub(now)
}

func IsZeroTime(tm time.Time) bool {
	maxZeroTime := time.Date(1970, time.January, 1, 9, 0, 0, 0, time.Local)
	if tm.Before(maxZeroTime) {
		return true
	} else {
		return false
	}
}
