package datecalc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/rickb777/date"
)

func TestCalendarDaysBetween(t *testing.T) {
	var d1, d2 time.Time
	// same day
	d1 = Date(2018, time.May, 31)
	d2 = Date(2018, time.May, 31)
	assert.Equal(t, int32(0), CalendarDaysBetween(d1, d2))

	// week
	d1 = Date(2018, time.May, 10)
	d2 = d1.AddDate(0, 0, 7)
	assert.Equal(t, int32(7), CalendarDaysBetween(d1, d2))

	// month - starting in month with 31 days
	d1 = Date(2018, time.May, 10)
	d2 = d1.AddDate(0, 1, 0)
	assert.Equal(t, int32(31), CalendarDaysBetween(d1, d2))

	// month - starting in month with 30 days
	d1 = Date(2018, time.June, 10)
	d2 = d1.AddDate(0, 1, 0)
	assert.Equal(t, int32(30), CalendarDaysBetween(d1, d2))

	// month - starting in month with 28 days
	d1 = Date(2018, time.February, 10)
	d2 = d1.AddDate(0, 1, 0)
	assert.Equal(t, int32(28), CalendarDaysBetween(d1, d2))

	// month - starting in month with 29 days
	d1 = Date(2020, time.February, 10)
	d2 = d1.AddDate(0, 1, 0)
	assert.Equal(t, int32(29), CalendarDaysBetween(d1, d2))

	// non leap year
	d1 = Date(2018, time.January, 1)
	d2 = d1.AddDate(1, 0, 0)
	assert.Equal(t, int32(365), CalendarDaysBetween(d1, d2))

	// leap year
	d1 = Date(2020, time.January, 1)
	d2 = d1.AddDate(1, 0, 0)
	assert.Equal(t, int32(366), CalendarDaysBetween(d1, d2))

	// case where d1 > d2 - ordering shouldn't matter
	assert.Equal(t, int32(366), CalendarDaysBetween(d2, d1))
}

func TestWeekdayDaysBetween(t *testing.T) {
	var d1, d2 time.Time
	// same day - non weekend
	d1 = Date(2018, time.May, 31)
	d2 = Date(2018, time.May, 31)
	assert.Equal(t, int32(1), WeekdayDaysBetween(d1, d2))

	// same day - weekend
	d1 = Date(2018, time.June, 2)
	d2 = Date(2018, time.June, 2)
	assert.Equal(t, int32(0), WeekdayDaysBetween(d1, d2))

	// month span
	d1 = Date(2018, time.May, 31)
	d2 = Date(2018, time.June, 30)
	assert.Equal(t, int32(22), WeekdayDaysBetween(d1, d2))

	// case where d1 > d2 - ordering shouldn't matter
	d1 = Date(2018, time.May, 31)
	d2 = Date(2018, time.May, 1)
	assert.Equal(t, int32(23), WeekdayDaysBetween(d1, d2))
}

func TestWeekendDaysBetween(t *testing.T) {
	var d1, d2 time.Time
	// same day - non weekend
	d1 = Date(2018, time.May, 31)
	d2 = Date(2018, time.May, 31)
	assert.Equal(t, int32(0), WeekendDaysBetween(d1, d2))

	// same day - weekend
	d1 = Date(2018, time.June, 2)
	d2 = Date(2018, time.June, 2)
	assert.Equal(t, int32(1), WeekendDaysBetween(d1, d2))

	// month span
	d1 = Date(2018, time.May, 31)
	d2 = Date(2018, time.June, 30)
	assert.Equal(t, int32(9), WeekendDaysBetween(d1, d2))

	// case where d1 > d2 - ordering shouldn't matter
	d1 = Date(2018, time.May, 31)
	d2 = Date(2018, time.May, 1)
	assert.Equal(t, int32(8), WeekendDaysBetween(d1, d2))
}

func BenchmarkDaysIn(b *testing.B) {
	for n := 0; n < b.N; n++ {
		date.DaysIn(1994, time.May)
	}
}
