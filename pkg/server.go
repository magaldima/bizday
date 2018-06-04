package pkg

import (
	"sync"
	"time"

	"github.com/magaldima/bizday/dcb"
	"github.com/magaldima/bizday/holiday"

	"github.com/hashicorp/go-plugin"

	api "github.com/magaldima/bizday/api"

	"github.com/magaldima/bizday/datecalc"
	"github.com/rickb777/date"
	context "golang.org/x/net/context"
)

// server implements the bizday API
type server struct {
	dcbRegistry     dcbRegistry
	holidayRegistry holidayRegistry
}

// New creates a new server
func New(dcbProtocol plugin.ClientProtocol, holidayProtocol plugin.ClientProtocol) api.DateCalcServer {
	return &server{
		dcbRegistry: dcbRegistry{
			source: dcbProtocol,
			mu:     sync.Mutex{},
			dcbs:   make(map[string]dcb.DayCountBasis),
		},
		holidayRegistry: holidayRegistry{
			source:   holidayProtocol,
			mu:       sync.Mutex{},
			holidays: make(map[string]holiday.Holiday),
		},
	}
}

// DaysBetween returns the number of days between a pair of dates using the day count basis (DCB)
func (s *server) DaysBetween(ctx context.Context, req *api.BinaryDateRequest) (*api.NumberOfDaysResponse, error) {
	days, err := s.daysBetween(req.Cal.DayCountBasis, req.Start.Time(), req.End.Time())
	if err != nil {
		return nil, err
	}
	return &api.NumberOfDaysResponse{Days: days}, nil
}

// CalendarDaysBetween returns the number of days between a pair of dates using the calendar
func (s *server) CalendarDaysBetween(ctx context.Context, req *api.BinaryDateRequest) (*api.NumberOfDaysResponse, error) {
	days := datecalc.CalendarDaysBetween(req.Start.Time(), req.End.Time())
	return &api.NumberOfDaysResponse{Days: days}, nil
}

// BizDaysBetween returns the number of business day dates between a pair of dates
func (s *server) BizDaysBetween(ctx context.Context, req *api.BinaryDateRequest) (*api.NumberOfDaysResponse, error) {
	total := datecalc.CalendarDaysBetween(req.Start.Time(), req.End.Time())
	weekends := datecalc.WeekendDaysBetween(req.Start.Time(), req.End.Time())
	// get holiday plugin
	holiday, err := s.getHoliday(req.Cal.Holiday)
	if err != nil {
		return nil, err
	}
	holidays, err := holiday.Delta(req.Start.Time(), req.End.Time())
	if err != nil {
		return nil, err
	}
	return &api.NumberOfDaysResponse{Days: total - weekends - holidays}, nil
}

// WeekdaysBetween returns the number of weekday dates between a pair of dates
func (s *server) WeekdaysBetween(ctx context.Context, req *api.BinaryDateRequest) (*api.NumberOfDaysResponse, error) {
	weekdays := datecalc.WeekdayDaysBetween(req.Start.Time(), req.End.Time())
	return &api.NumberOfDaysResponse{Days: weekdays}, nil
}

// WeekendsBetween returns the number of weekend dates between a pair of dates
func (s *server) WeekendsBetween(ctx context.Context, req *api.BinaryDateRequest) (*api.NumberOfDaysResponse, error) {
	weekends := datecalc.WeekendDaysBetween(req.Start.Time(), req.End.Time())
	return &api.NumberOfDaysResponse{Days: weekends}, nil
}

// HolidaysBetween returns the number of holiday dates between a pair of dates
func (s *server) HolidaysBetween(ctx context.Context, req *api.BinaryDateRequest) (*api.NumberOfDaysResponse, error) {
	// get holiday plugin
	holiday, err := s.getHoliday(req.Cal.Holiday)
	if err != nil {
		return nil, err
	}
	holidays, err := holiday.Delta(req.Start.Time(), req.End.Time())
	if err != nil {
		return nil, err
	}
	return &api.NumberOfDaysResponse{Days: holidays}, nil
}

// BizDaysInMonth returns the number of business day dates in the month of the provided date
func (s *server) BizDaysInMonth(ctx context.Context, req *api.UnaryDateRequest) (*api.NumberOfDaysResponse, error) {
	month := api.GetTimeMonth(req.Date)
	lastDay := date.DaysIn(int(req.Date.Year), month)
	start := datecalc.Date(int(req.Date.Year), month, 0)
	end := datecalc.Date(int(req.Date.Year), month, lastDay)

	total := datecalc.CalendarDaysBetween(start, end)
	weekends := datecalc.WeekendDaysBetween(start, end)

	// get holiday plugin
	holiday, err := s.getHoliday(req.Cal.Holiday)
	if err != nil {
		return nil, err
	}
	holidays, err := holiday.Delta(start, end)
	if err != nil {
		return nil, err
	}
	return &api.NumberOfDaysResponse{Days: total - weekends - holidays}, nil
}

// BizDaysInYear returns the number of business day dates in the year of the provided date
func (s *server) BizDaysInYear(ctx context.Context, req *api.UnaryDateRequest) (*api.NumberOfDaysResponse, error) {
	start := datecalc.Date(int(req.Date.Year), time.January, 1)
	end := start.AddDate(1, 0, 0)

	total := datecalc.CalendarDaysBetween(start, end)
	weekends := datecalc.WeekendDaysBetween(start, end)

	// get holiday plugin
	holiday, err := s.getHoliday(req.Cal.Holiday)
	if err != nil {
		return nil, err
	}
	holidays, err := holiday.Delta(start, end)
	if err != nil {
		return nil, err
	}
	return &api.NumberOfDaysResponse{Days: total - weekends - holidays}, nil
}

// IsBizDay returns true if the date provided is a business day
func (s *server) IsBizDay(ctx context.Context, req *api.UnaryDateRequest) (*api.UnaryBoolResponse, error) {
	d := datecalc.Date(int(req.Date.Year), api.GetTimeMonth(req.Date), int(req.Date.Day))
	holiday, err := s.getHoliday(req.Cal.Holiday)
	if err != nil {
		return nil, err
	}
	res, err := s.isBusinessDay(d, holiday)
	if err != nil {
		return nil, err
	}
	return &api.UnaryBoolResponse{Ok: res}, nil
}

func (s *server) IsFirstBizDayOfMonth(context.Context, *api.UnaryDateRequest) (*api.UnaryBoolResponse, error) {
	return nil, nil
}

func (s *server) IsLastBizDayOfMonth(context.Context, *api.UnaryDateRequest) (*api.UnaryBoolResponse, error) {
	return nil, nil
}

func (s *server) FirstBizDayOfMonth(context.Context, *api.UnaryDateRequest) (*api.UnaryDateResponse, error) {
	return nil, nil
}

func (s *server) LastBizDayOfMonth(context.Context, *api.UnaryDateRequest) (*api.UnaryDateResponse, error) {
	return nil, nil
}

func (s *server) FirstBizDayOfQtr(context.Context, *api.UnaryDateRequest) (*api.UnaryDateResponse, error) {
	return nil, nil
}

func (s *server) LastBizDayOfQtr(context.Context, *api.UnaryDateRequest) (*api.UnaryDateResponse, error) {
	return nil, nil
}

// NextBizDay returns the next business day
func (s *server) NextBizDay(ctx context.Context, req *api.UnaryDateRequest) (*api.UnaryDateResponse, error) {
	cur := datecalc.Date(int(req.Date.Year), api.GetTimeMonth(req.Date), int(req.Date.Day))
	holiday, err := s.getHoliday(req.Cal.Holiday)
	if err != nil {
		return nil, err
	}
	next, err := s.AddBusinessDays(cur, 1, holiday)
	if err != nil {
		return nil, err
	}
	return &api.UnaryDateResponse{Date: api.ConvertToProtoDate(next)}, nil
}

// PrevBizDay returns the previous business day
func (s *server) PrevBizDay(ctx context.Context, req *api.UnaryDateRequest) (*api.UnaryDateResponse, error) {
	cur := datecalc.Date(int(req.Date.Year), api.GetTimeMonth(req.Date), int(req.Date.Day))
	holiday, err := s.getHoliday(req.Cal.Holiday)
	if err != nil {
		return nil, err
	}
	prev, err := s.AddBusinessDays(cur, -1, holiday)
	if err != nil {
		return nil, err
	}
	return &api.UnaryDateResponse{Date: api.ConvertToProtoDate(prev)}, nil
}

// AddBizDays returns the business day that corresponds to the transformation of the given date by the given offset
func (s *server) AddBizDays(ctx context.Context, req *api.UnaryTransformRequest) (*api.UnaryDateResponse, error) {
	cur := datecalc.Date(int(req.Date.Year), api.GetTimeMonth(req.Date), int(req.Date.Day))
	holiday, err := s.getHoliday(req.Cal.Holiday)
	if err != nil {
		return nil, err
	}
	updated, err := s.AddBusinessDays(cur, int(req.Offset), holiday)
	if err != nil {
		return nil, err
	}
	return &api.UnaryDateResponse{Date: api.ConvertToProtoDate(updated)}, nil
}
