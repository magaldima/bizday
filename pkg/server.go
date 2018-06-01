package pkg

import (
	"time"

	api "github.com/magaldima/bizday/api"

	"github.com/magaldima/bizday/datecalc"
	"github.com/rickb777/date"
	context "golang.org/x/net/context"
)

// Server implements the bizday API
type Server struct {
}

// BizDaysBetween returns the number of business day dates between a pair of dates
func (s Server) BizDaysBetween(ctx context.Context, req *api.BinaryDateRequest) (*api.NumberOfDaysResponse, error) {
	total := datecalc.CalendarDaysBetween(req.Start.Time(), req.End.Time())
	weekends := datecalc.WeekendDaysBetween(req.Start.Time(), req.End.Time())
	holidays := datecalc.HolidaysBetween(req.Start.Time(), req.End.Time())
	return &api.NumberOfDaysResponse{Days: total - weekends - holidays}, nil
}

// WeekdaysBetween returns the number of weekday dates between a pair of dates
func (s Server) WeekdaysBetween(ctx context.Context, req *api.BinaryDateRequest) (*api.NumberOfDaysResponse, error) {
	weekdays := datecalc.WeekdayDaysBetween(req.Start.Time(), req.End.Time())
	return &api.NumberOfDaysResponse{Days: weekdays}, nil
}

// WeekendsBetween returns the number of weekend dates between a pair of dates
func (s Server) WeekendsBetween(ctx context.Context, req *api.BinaryDateRequest) (*api.NumberOfDaysResponse, error) {
	weekends := datecalc.WeekendDaysBetween(req.Start.Time(), req.End.Time())
	return &api.NumberOfDaysResponse{Days: weekends}, nil
}

// HolidaysBetween returns the number of holiday dates between a pair of dates
func (s Server) HolidaysBetween(ctx context.Context, req *api.BinaryDateRequest) (*api.NumberOfDaysResponse, error) {
	holidays := datecalc.HolidaysBetween(req.Start.Time(), req.End.Time())
	return &api.NumberOfDaysResponse{Days: holidays}, nil
}

// BizDaysInMonth returns the number of business day dates in the month of the provided date
func (s Server) BizDaysInMonth(ctx context.Context, req *api.UnaryDateRequest) (*api.NumberOfDaysResponse, error) {
	month := api.GetTimeMonth(req.Date)
	lastDay := date.DaysIn(int(req.Date.Year), month)
	start := datecalc.Date(int(req.Date.Year), month, 0)
	end := datecalc.Date(int(req.Date.Year), month, lastDay)

	total := datecalc.CalendarDaysBetween(start, end)
	weekends := datecalc.WeekendDaysBetween(start, end)
	holidays := datecalc.HolidaysBetween(start, end)
	return &api.NumberOfDaysResponse{Days: total - weekends - holidays}, nil
}

// BizDaysInYear returns the number of business day dates in the year of the provided date
func (s Server) BizDaysInYear(ctx context.Context, req *api.UnaryDateRequest) (*api.NumberOfDaysResponse, error) {
	start := datecalc.Date(int(req.Date.Year), time.January, 1)
	end := start.AddDate(1, 0, 0)

	total := datecalc.CalendarDaysBetween(start, end)
	weekends := datecalc.WeekendDaysBetween(start, end)
	holidays := datecalc.HolidaysBetween(start, end)
	return &api.NumberOfDaysResponse{Days: total - weekends - holidays}, nil
}

// IsBizDay returns true if the date provided is a business day
func (s Server) IsBizDay(ctx context.Context, req *api.UnaryDateRequest) (*api.UnaryBoolResponse, error) {
	d := datecalc.Date(int(req.Date.Year), api.GetTimeMonth(req.Date), int(req.Date.Day))
	return &api.UnaryBoolResponse{Ok: datecalc.IsBusinessDay(d)}, nil
}

func (s Server) IsFirstBizDayOfMonth(context.Context, *api.UnaryDateRequest) (*api.UnaryBoolResponse, error) {
	return nil, nil
}

func (s Server) IsLastBizDayOfMonth(context.Context, *api.UnaryDateRequest) (*api.UnaryBoolResponse, error) {
	return nil, nil
}

func (s Server) FirstBizDayOfMonth(context.Context, *api.UnaryDateRequest) (*api.UnaryDateResponse, error) {
	return nil, nil
}

func (s Server) LastBizDayOfMonth(context.Context, *api.UnaryDateRequest) (*api.UnaryDateResponse, error) {
	return nil, nil
}

func (s Server) FirstBizDayOfQtr(context.Context, *api.UnaryDateRequest) (*api.UnaryDateResponse, error) {
	return nil, nil
}

func (s Server) LastBizDayOfQtr(context.Context, *api.UnaryDateRequest) (*api.UnaryDateResponse, error) {
	return nil, nil
}

// NextBizDay returns the next business day
func (s Server) NextBizDay(ctx context.Context, req *api.UnaryDateRequest) (*api.UnaryDateResponse, error) {
	cur := datecalc.Date(int(req.Date.Year), api.GetTimeMonth(req.Date), int(req.Date.Day))
	next := datecalc.AddBusinessDays(cur, 1)
	return &api.UnaryDateResponse{Date: api.ConvertToProtoDate(next)}, nil
}

// PrevBizDay returns the previous business day
func (s Server) PrevBizDay(ctx context.Context, req *api.UnaryDateRequest) (*api.UnaryDateResponse, error) {
	cur := datecalc.Date(int(req.Date.Year), api.GetTimeMonth(req.Date), int(req.Date.Day))
	prev := datecalc.AddBusinessDays(cur, -1)
	return &api.UnaryDateResponse{Date: api.ConvertToProtoDate(prev)}, nil
}

// AddBizDays returns the business day that corresponds to the transformation of the given date by the given offset
func (s Server) AddBizDays(ctx context.Context, req *api.UnaryTransformRequest) (*api.UnaryDateResponse, error) {
	cur := datecalc.Date(int(req.Date.Year), api.GetTimeMonth(req.Date), int(req.Date.Day))
	updated := datecalc.AddBusinessDays(cur, int(req.Offset))
	return &api.UnaryDateResponse{Date: api.ConvertToProtoDate(updated)}, nil
}
