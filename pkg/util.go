package pkg

import (
	"fmt"
	"time"

	"github.com/rickb777/date"
)

func (d *Date) date() date.Date {
	return date.New(int(d.Year), getTimeMonth(d), int(d.Day))
}

func (d *Date) time() time.Time {
	return time.Date(int(d.Year), getTimeMonth(d), int(d.Day), 0, 0, 0, 0, time.UTC)
}

// helper function to find the golang time Month from the protocol buffer date
func getTimeMonth(d *Date) time.Month {
	switch d.Month {
	case Month_January:
		return time.January
	case Month_February:
		return time.February
	case Month_March:
		return time.March
	case Month_April:
		return time.April
	case Month_May:
		return time.May
	case Month_June:
		return time.June
	case Month_July:
		return time.July
	case Month_August:
		return time.August
	case Month_September:
		return time.September
	case Month_October:
		return time.October
	case Month_November:
		return time.November
	case Month_December:
		return time.December
	default:
		panic(fmt.Errorf("unknown month of date %v", d))
	}
}

func getDateMonth(t time.Time) Month {
	switch t.Month() {
	case time.January:
		return Month_January
	case time.February:
		return Month_February
	case time.March:
		return Month_March
	case time.April:
		return Month_April
	case time.May:
		return Month_May
	case time.June:
		return Month_June
	case time.July:
		return Month_July
	case time.August:
		return Month_August
	case time.September:
		return Month_September
	case time.October:
		return Month_October
	case time.November:
		return Month_November
	case time.December:
		return Month_December
	default:
		panic(fmt.Errorf("unknown month of time %v", t))
	}
}

func convertToProtoDate(t time.Time) *Date {
	return &Date{
		Year:  int32(t.Year()),
		Month: getDateMonth(t),
		Day:   int32(t.Day()),
	}
}
