package pkg

import (
	"time"

	"github.com/magaldima/bizday/holiday"
)

// DateTime returns the Time corresponding to
//	yyyy-mm-dd 00:00:00 + 0 nanoseconds
// in Universal Coordinated Time (UTC)
func DateTime(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func (s *server) daysBetween(dcb string, d1, d2 time.Time) (int32, error) {
	cal, err := s.getDCB(dcb)
	if err != nil {
		return 0, err
	}
	return cal.Delta(d1, d2)
}

func (s *server) AddDaysUsingDCB(dcb string, d time.Time, days int32) (time.Time, error) {
	cal, err := s.getDCB(dcb)
	if err != nil {
		return time.Time{}, err
	}
	daysInMonth, err := cal.DaysInMonth(d)
	if err != nil {
		return time.Time{}, err
	}
	monthsToAdd := int(days / daysInMonth)
	daysToAdd := int(days % daysInMonth)

	res := d.AddDate(0, monthsToAdd, daysToAdd)
	return res, nil
}

// AddBusinessDays adds n business days to time t and returns the new time
// n can be positive or negative depending on incrementing or decrementing the days
func (s *server) AddBusinessDays(curr time.Time, n int, holiday holiday.Holiday) (time.Time, error) {
	if n == 0 {
		return curr, nil
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
		ok, err := s.isBusinessDay(next, holiday)
		if err != nil {
			return next, err
		}
		if ok {
			count++
		}
		curr = next
	}
	return next, nil
}

// IsBusinessDay returns true if the day provided is a business day
func (s *server) isBusinessDay(t time.Time, holiday holiday.Holiday) (bool, error) {
	isHoliday, err := holiday.IsHoliday(t)
	if err != nil {
		return false, err
	}
	return !(isWeekendDay(t.Weekday()) || isHoliday), nil
}

// IsWeekendDay returns true if the day provided is either Saturday or Sunday
func isWeekendDay(day time.Weekday) bool {
	if day == time.Saturday || day == time.Sunday {
		return true
	}
	return false
}
