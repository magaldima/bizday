package pkg

import (
	api "github.com/magaldima/bizday/api"
	"github.com/magaldima/bizday/datecalc"
)

func (s *server) DaysBetween(dcb string, d1 api.Date, d2 api.Date) (int32, error) {
	cal, err := s.getCalendar(dcb)
	if err != nil {
		return 0, err
	}
	return cal.Delta(d1, d2), nil
}

func (s *server) AddDaysUsingDCB(dcb string, date api.Date, days int32) (api.Date, error) {
	d := datecalc.Date(int(date.Year), api.GetTimeMonth(&date), int(date.Day))

	cal, err := s.getCalendar(dcb)
	if err != nil {
		return api.Date{}, err
	}
	daysInMonth := cal.DaysInMonth(date)
	monthsToAdd := int(days / daysInMonth)
	daysToAdd := int(days % daysInMonth)

	res := d.AddDate(0, monthsToAdd, daysToAdd)

	return *api.ConvertToProtoDate(res), nil
}

// AlphaFixed computes the day count fraction with constant denominator
func (s *server) AlphaFixed(dcb string, d1 api.Date, d2 api.Date) (float64, error) {
	cal, err := s.getCalendar(dcb)
	if err != nil {
		return 0, err
	}
	delta := cal.Delta(d1, d2)
	den := cal.DaysInYear(d1)
	return float64(delta) / float64(den), nil
}

// AlphaVariable computes the day count fraction with variable denominator
func (s *server) AlphaVariable(dcb string, d1 api.Date, d2 api.Date) (float64, error) {
	cal, err := s.getCalendar(dcb)
	if err != nil {
		return 0, err
	}
	diy1 := int(cal.DaysInYear(d1))
	diy2 := int(cal.DaysInYear(d2))

	d1a := datecalc.Date(int(d1.Year), api.GetTimeMonth(&d1), int(d1.Day))
	d2a := datecalc.Date(int(d2.Year), api.GetTimeMonth(&d2), int(d2.Day))

	doy1 := d1a.YearDay()
	doy2 := d2a.YearDay()

	frac := float64(diy1*doy2-diy2*doy1) / float64(diy1*diy2)
	return float64(d2.Year-d1.Year) + frac, nil
}
