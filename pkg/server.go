package pkg

import (
	"sync"
	"time"

	"github.com/magaldima/bizday/calendar"

	"github.com/hashicorp/go-plugin"

	api "github.com/magaldima/bizday/api"
	"github.com/magaldima/bizday/holidays/shared"

	"github.com/magaldima/bizday/datecalc"
	"github.com/rickb777/date"
	context "golang.org/x/net/context"
)

// server implements the bizday API
type server struct {
	calendarRegistry calendarRegistry
	holidayRegistry  holidayRegistry
}

// New creates a new server
func New(calendarProtocol plugin.ClientProtocol, holidayProtocol plugin.ClientProtocol) api.DateCalcServer {
	return &server{
		calendarRegistry: calendarRegistry{
			source:    calendarProtocol,
			mu:        sync.Mutex{},
			calendars: make(map[string]calendar.Calendar),
		},
		holidayRegistry: holidayRegistry{
			source:   holidayProtocol,
			mu:       sync.Mutex{},
			holidays: make(map[string]shared.Holiday),
		},
	}
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
	holidays := holiday.Delta(*req.Start, *req.End)
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
	holidays := holiday.Delta(*req.Start, *req.End)
	return &api.NumberOfDaysResponse{Days: holidays}, nil
}

// BizDaysInMonth returns the number of business day dates in the month of the provided date
func (s *server) BizDaysInMonth(ctx context.Context, req *api.UnaryDateRequest) (*api.NumberOfDaysResponse, error) {
	month := api.GetTimeMonth(req.Date)
	lastDay := date.DaysIn(int(req.Date.Year), month)
	start := datecalc.Date(int(req.Date.Year), month, 0)
	end := datecalc.Date(int(req.Date.Year), month, lastDay)

	startDate := api.Date{
		Year:  req.Date.Year,
		Month: req.Date.Month,
		Day:   1,
	}
	endDate := api.Date{
		Year:  req.Date.Year,
		Month: req.Date.Month,
		Day:   int32(lastDay),
	}

	total := datecalc.CalendarDaysBetween(start, end)
	weekends := datecalc.WeekendDaysBetween(start, end)

	// get holiday plugin
	holiday, err := s.getHoliday(req.Cal.Holiday)
	if err != nil {
		return nil, err
	}
	holidays := holiday.Delta(startDate, endDate)
	return &api.NumberOfDaysResponse{Days: total - weekends - holidays}, nil
}

// BizDaysInYear returns the number of business day dates in the year of the provided date
func (s *server) BizDaysInYear(ctx context.Context, req *api.UnaryDateRequest) (*api.NumberOfDaysResponse, error) {
	start := datecalc.Date(int(req.Date.Year), time.January, 1)
	end := start.AddDate(1, 0, 0)

	total := datecalc.CalendarDaysBetween(start, end)
	weekends := datecalc.WeekendDaysBetween(start, end)

	startDate := api.Date{
		Year:  req.Date.Year,
		Month: api.Month_January,
		Day:   1,
	}
	endDate := api.Date{
		Year:  int32(end.Year()),
		Month: api.Month_January,
		Day:   1,
	}

	// get holiday plugin
	holiday, err := s.getHoliday(req.Cal.Holiday)
	if err != nil {
		return nil, err
	}
	holidays := holiday.Delta(startDate, endDate)

	return &api.NumberOfDaysResponse{Days: total - weekends - holidays}, nil
}

// IsBizDay returns true if the date provided is a business day
func (s *server) IsBizDay(ctx context.Context, req *api.UnaryDateRequest) (*api.UnaryBoolResponse, error) {
	d := datecalc.Date(int(req.Date.Year), api.GetTimeMonth(req.Date), int(req.Date.Day))
	return &api.UnaryBoolResponse{Ok: datecalc.IsBusinessDay(d)}, nil
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
	next := datecalc.AddBusinessDays(cur, 1)
	return &api.UnaryDateResponse{Date: api.ConvertToProtoDate(next)}, nil
}

// PrevBizDay returns the previous business day
func (s *server) PrevBizDay(ctx context.Context, req *api.UnaryDateRequest) (*api.UnaryDateResponse, error) {
	cur := datecalc.Date(int(req.Date.Year), api.GetTimeMonth(req.Date), int(req.Date.Day))
	prev := datecalc.AddBusinessDays(cur, -1)
	return &api.UnaryDateResponse{Date: api.ConvertToProtoDate(prev)}, nil
}

// AddBizDays returns the business day that corresponds to the transformation of the given date by the given offset
func (s *server) AddBizDays(ctx context.Context, req *api.UnaryTransformRequest) (*api.UnaryDateResponse, error) {
	cur := datecalc.Date(int(req.Date.Year), api.GetTimeMonth(req.Date), int(req.Date.Day))
	updated := datecalc.AddBusinessDays(cur, int(req.Offset))
	return &api.UnaryDateResponse{Date: api.ConvertToProtoDate(updated)}, nil
}
