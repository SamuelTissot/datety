package datety

import (
	"time"
)

// IsSameDay returns true if both dates are on the same day, same month and same year
func IsSameDay(t1, t2 time.Time) bool {
	t1 = t1.UTC()
	t2 = t2.UTC()
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// IsSameMonth return true if both date are on the same month and year
func IsSameMonth(t1, t2 time.Time) bool {
	t1 = t1.UTC()
	t2 = t2.UTC()
	y1, m1, _ := t1.Date()
	y2, m2, _ := t2.Date()
	return y1 == y2 && m1 == m2
}

// IsSameYear returns true if both date are on the same year
func IsSameYear(t1, t2 time.Time) bool {
	t1 = t1.UTC()
	t2 = t2.UTC()
	y1, _, _ := t1.Date()
	y2, _, _ := t2.Date()
	return y1 == y2
}

// IsSamWithinThreshold return true if t1 is between t2 - threshold AND t2 + threshold
func IsSamWithinThreshold(t1, t2 time.Time, threshold time.Duration) bool {
	if t1.Equal(t2) {
		return true
	}
	if t1.After(t2.Add(-1*threshold)) && t1.Before(t2.Add(threshold)) {
		return true
	}

	return false
}

// IsToday return true if date is today
func IsToday(date time.Time) bool {
	return IsSameDay(date, time.Now())
}

// NumberOfMonths return the number of month separating from to
func NumberOfMonths(from time.Time, to time.Time) int {
	if from.After(to) {
		return 0
	}

	if (from.Month() == to.Month()) && (from.Year() == to.Year()) {
		return 0
	}

	return 1 + NumberOfMonths(from.AddDate(0, 1, 0), to)
}

// NumberOfDays return the number of days separating from, to
func NumberOfDays(from, to time.Time) int {
	if from.After(to) {
		return 0
	}

	if IsSameDay(from, to) {
		return 0
	}

	return 1 + NumberOfDays(from.AddDate(0, 0, 1), to)
}

// NumberOfDays return the number of hours between from and two
func NumberOfHours(from, to time.Time) int {
	d := to.Sub(from)
	return int(d / time.Hour)
}

// TodayAtMidnight return today's date floored to midnight
func TodayAtMidnight() time.Time {
	return DayFloor(time.Now())
}

// BeginningOfMonth returns the the time of the first day of the month of t
func BeginningOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// HourFloor return the time with min:sec:nsec to 0:0:0
func HourFloor(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
}

// DayFloor set the day to midnight
func DayFloor(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// Ceil returns the time with the hour set to 23:59:59:9999999
func Ceil(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 9999999, t.Location())
}
