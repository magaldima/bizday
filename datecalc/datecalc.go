package datecalc

import (
	"time"

	"github.com/rickb777/date"
)

var Weekdays = []time.Weekday{
	time.Sunday,
	time.Monday,
	time.Tuesday,
	time.Wednesday,
	time.Thursday,
	time.Friday,
	time.Saturday,
}

type quarter struct {
	month int
	day   int
}

var (
	Q1 = quarter{3, 31}
	Q2 = quarter{6, 30}
	Q3 = quarter{9, 30}
	Q4 = quarter{12, 31}
)

// Date returns the Time corresponding to
//	yyyy-mm-dd 00:00:00 + 0 nanoseconds
// in Universal Coordinated Time (UTC)
func Date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// CalendarDaysBetween returns the number of days between two dates using the calendar
// does not assume proper date order
func CalendarDaysBetween(t1, t2 time.Time) int32 {
	var d1, d2 date.Date
	if t1.After(t2) {
		d1 = date.New(t2.Year(), t2.Month(), t2.Day())
		d2 = date.New(t1.Year(), t1.Month(), t1.Day())
	} else {
		d1 = date.New(t1.Year(), t1.Month(), t1.Day())
		d2 = date.New(t2.Year(), t2.Month(), t2.Day())
	}
	return int32(d2.Sub(d1))
}

// WeekendDaysBetween returns the number of weekend days between two dates
// we do not assume proper date order
func WeekendDaysBetween(t1, t2 time.Time) int32 {
	var count int32
	for _, day := range Weekdays {
		if IsWeekendDay(day) {
			count += daysOfWeekBetween(t1, t2, day)
		}
	}
	return count
}

// WeekdayDaysBetween returns the number of weekday days between two dates
// we do not assume proper date order
func WeekdayDaysBetween(t1, t2 time.Time) int32 {
	var count int32
	for _, day := range Weekdays {
		if !IsWeekendDay(day) {
			count += daysOfWeekBetween(t1, t2, day)
		}
	}
	return count
}

// HolidaysBetween returns the number of holidays between two dates
func HolidaysBetween(t1, t2 time.Time) int32 {
	//todo
	return 0
}

// IsHoliday returns true if the day provided is a holiday
func IsHoliday(t time.Time) bool {
	return false
}

// daysOfWeekBetween returns the number of days of this day between two dates
// the range is inclusive of the two dates
// does not assume the dates are properly-ordered
func daysOfWeekBetween(t1, t2 time.Time, day time.Weekday) int32 {
	var d1, d2 date.Date
	if t1.After(t2) {
		d1 = date.New(t2.Year(), t2.Month(), t2.Day())
		d2 = date.New(t1.Year(), t1.Month(), t1.Day())
	} else {
		d1 = date.New(t1.Year(), t1.Month(), t1.Day())
		d2 = date.New(t2.Year(), t2.Month(), t2.Day())
	}

	// find the first day with this Weekday after the starting date
	start := int32(d1.DaysSinceEpoch())
	end := int32(d2.DaysSinceEpoch())
	d1Dow := d1.Weekday()
	start += (int32(day) + 7 - int32(d1Dow)) % 7
	if start == end {
		return 1
	}
	if start > end {
		return 0
	}

	// now count the number of weeks and add one for the start day
	return ((end - start) / 7) + 1
}

// IsWeekendDay returns true if the day provided is either Saturday or Sunday
func IsWeekendDay(day time.Weekday) bool {
	if day == time.Saturday || day == time.Sunday {
		return true
	}
	return false
}

// IsBusinessDay returns true if the day provided is a business day
func IsBusinessDay(t time.Time) bool {
	return !(IsWeekendDay(t.Weekday()) || IsHoliday(t))
}

// AddBusinessDays adds n business days to time t and returns the new time
// n can be positive or negative depending on incrementing or decrementing the days
func AddBusinessDays(curr time.Time, n int) time.Time {
	if n == 0 {
		return curr
	}

	var dir int
	if n < 0 {
		dir = -1
		n = -n
	} else {
		dir = 1
	}

	var next time.Time
	count := 0
	for count < n {
		next = curr.AddDate(0, 0, dir) // add one calendar interval at a time
		if ok := IsBusinessDay(next); ok {
			count++
		}
		curr = next
	}
	return next
}
