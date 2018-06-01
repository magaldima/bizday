package main

import (
	"math"

	plugin "github.com/hashicorp/go-plugin"
	api "github.com/magaldima/bizday/api"
	"github.com/magaldima/bizday/calendar"
)

// Calendar30_360 is a real implementation of Calendar with day count basis 30/360
type Calendar30_360 struct{}

func (Calendar30_360) DaysInYear(d api.Date) int32 {
	return 360
}

func (Calendar30_360) DaysInMonth(d api.Date) int32 {
	return 30
}

func (Calendar30_360) Delta(start api.Date, end api.Date) int32 {
	d1, m1, y1 := start.Day, start.Month, start.Year
	d2, m2, y2 := end.Day, end.Month, end.Year

	// 1. if d1 is 31, change d1 to 30
	if d1 == 31 {
		d1 = 30
	}
	// 2. if d2 is 31, change d2 to 30
	if d2 == 31 {
		d2 = 30
	}

	// 3. get number of days
	return int32(math.Max((360*float64(y2-y1) + float64(30*(int32(m2)-int32(m1))) + float64(d2-d1)), 0))
}

func (c Calendar30_360) Alpha(start api.Date, end api.Date) float64 {
	delta := c.Delta(start, end)
	den := c.DaysInYear(start)
	return float64(delta) / float64(den)
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: calendar.Handshake,
		Plugins: map[string]plugin.Plugin{
			"30_360": &calendar.CalendarPlugin{Impl: &Calendar30_360{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
