package pkg

import (
	"time"

	"github.com/magaldima/bizday/datecalc"
	"github.com/rickb777/date"
	context "golang.org/x/net/context"
)

// Server implements the bizday API
type Server struct {
}

// BizDaysBetween returns the number of business day dates between a pair of dates
func (s Server) BizDaysBetween(ctx context.Context, req *BinaryDateRequest) (*NumberOfDaysResponse, error) {
	total := datecalc.CalendarDaysBetween(req.Start.time(), req.End.time())
	weekends := datecalc.WeekendDaysBetween(req.Start.time(), req.End.time())
	holidays := datecalc.HolidaysBetween(req.Start.time(), req.End.time())
	return &NumberOfDaysResponse{Days: total - weekends - holidays}, nil
}

// WeekdaysBetween returns the number of weekday dates between a pair of dates
func (s Server) WeekdaysBetween(ctx context.Context, req *BinaryDateRequest) (*NumberOfDaysResponse, error) {
	weekdays := datecalc.WeekdayDaysBetween(req.Start.time(), req.End.time())
	return &NumberOfDaysResponse{Days: weekdays}, nil
}

// WeekendsBetween returns the number of weekend dates between a pair of dates
func (s Server) WeekendsBetween(ctx context.Context, req *BinaryDateRequest) (*NumberOfDaysResponse, error) {
	weekends := datecalc.WeekendDaysBetween(req.Start.time(), req.End.time())
	return &NumberOfDaysResponse{Days: weekends}, nil
}

// HolidaysBetween returns the number of holiday dates between a pair of dates
func (s Server) HolidaysBetween(ctx context.Context, req *BinaryDateRequest) (*NumberOfDaysResponse, error) {
	holidays := datecalc.HolidaysBetween(req.Start.time(), req.End.time())
	return &NumberOfDaysResponse{Days: holidays}, nil
}

// BizDaysInMonth returns the number of business day dates in the month of the provided date
func (s Server) BizDaysInMonth(ctx context.Context, req *UnaryDateRequest) (*NumberOfDaysResponse, error) {
	month := getTimeMonth(req.Date)
	lastDay := date.DaysIn(int(req.Date.Year), month)
	start := datecalc.Date(int(req.Date.Year), month, 0)
	end := datecalc.Date(int(req.Date.Year), month, lastDay)

	total := datecalc.CalendarDaysBetween(start, end)
	weekends := datecalc.WeekendDaysBetween(start, end)
	holidays := datecalc.HolidaysBetween(start, end)
	return &NumberOfDaysResponse{Days: total - weekends - holidays}, nil
}

// BizDaysInYear returns the number of business day dates in the year of the provided date
func (s Server) BizDaysInYear(ctx context.Context, req *UnaryDateRequest) (*NumberOfDaysResponse, error) {
	start := datecalc.Date(int(req.Date.Year), time.January, 1)
	end := start.AddDate(1, 0, 0)

	total := datecalc.CalendarDaysBetween(start, end)
	weekends := datecalc.WeekendDaysBetween(start, end)
	holidays := datecalc.HolidaysBetween(start, end)
	return &NumberOfDaysResponse{Days: total - weekends - holidays}, nil
}

// IsBizDay returns true if the date provided is a business day
func (s Server) IsBizDay(ctx context.Context, req *UnaryDateRequest) (*UnaryBoolResponse, error) {
	d := datecalc.Date(int(req.Date.Year), getTimeMonth(req.Date), int(req.Date.Day))
	return &UnaryBoolResponse{Ok: datecalc.IsBusinessDay(d)}, nil
}

func (s Server) IsFirstBizDayOfMonth(context.Context, *UnaryDateRequest) (*UnaryBoolResponse, error) {
	return nil, nil
}

func (s Server) IsLastBizDayOfMonth(context.Context, *UnaryDateRequest) (*UnaryBoolResponse, error) {
	return nil, nil
}

func (s Server) FirstBizDayOfMonth(context.Context, *UnaryDateRequest) (*UnaryDateResponse, error) {
	return nil, nil
}

func (s Server) LastBizDayOfMonth(context.Context, *UnaryDateRequest) (*UnaryDateResponse, error) {
	return nil, nil
}

func (s Server) FirstBizDayOfQtr(context.Context, *UnaryDateRequest) (*UnaryDateResponse, error) {
	return nil, nil
}

func (s Server) LastBizDayOfQtr(context.Context, *UnaryDateRequest) (*UnaryDateResponse, error) {
	return nil, nil
}

// NextBizDay returns the next business day
func (s Server) NextBizDay(ctx context.Context, req *UnaryDateRequest) (*UnaryDateResponse, error) {
	cur := datecalc.Date(int(req.Date.Year), getTimeMonth(req.Date), int(req.Date.Day))
	next := datecalc.AddBusinessDays(cur, 1)
	return &UnaryDateResponse{Date: convertToProtoDate(next)}, nil
}

// PrevBizDay returns the previous business day
func (s Server) PrevBizDay(ctx context.Context, req *UnaryDateRequest) (*UnaryDateResponse, error) {
	cur := datecalc.Date(int(req.Date.Year), getTimeMonth(req.Date), int(req.Date.Day))
	prev := datecalc.AddBusinessDays(cur, -1)
	return &UnaryDateResponse{Date: convertToProtoDate(prev)}, nil
}

// AddBizDays returns the business day that corresponds to the transformation of the given date by the given offset
func (s Server) AddBizDays(ctx context.Context, req *UnaryTransformRequest) (*UnaryDateResponse, error) {
	cur := datecalc.Date(int(req.Date.Year), getTimeMonth(req.Date), int(req.Date.Day))
	updated := datecalc.AddBusinessDays(cur, int(req.Offset))
	return &UnaryDateResponse{Date: convertToProtoDate(updated)}, nil
}
