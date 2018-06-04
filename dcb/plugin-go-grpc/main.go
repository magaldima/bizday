package main

import (
	"math"
	"time"

	"github.com/magaldima/bizday/dcb"

	plugin "github.com/hashicorp/go-plugin"
)

// dcb30_360 is a real implementation with day count basis 30/360
type dcb30_360 struct{}

func (dcb30_360) DaysInYear(t time.Time) (int32, error) {
	return 360, nil
}

func (dcb30_360) DaysInMonth(t time.Time) (int32, error) {
	return 30, nil
}

func (dcb30_360) Delta(start, end time.Time) (int32, error) {
	d1, m1, y1 := start.Day(), start.Month(), start.Year()
	d2, m2, y2 := end.Day(), end.Month(), end.Year()

	// 1. if d1 is 31, change d1 to 30
	if d1 == 31 {
		d1 = 30
	}
	// 2. if d2 is 31, change d2 to 30
	if d2 == 31 {
		d2 = 30
	}

	// 3. get number of days
	return int32(math.Max((360*float64(y2-y1) + float64(30*(int32(m2)-int32(m1))) + float64(d2-d1)), 0)), nil
}

func (c dcb30_360) Alpha(start, end time.Time) (float64, error) {
	delta, err := c.Delta(start, end)
	if err != nil {
		return 0, err
	}
	den, err := c.DaysInYear(start)
	if err != nil {
		return 0, err
	}
	return float64(delta) / float64(den), nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: dcb.Handshake,
		Plugins: map[string]plugin.Plugin{
			"30_360": &dcb.DcbPlugin{Impl: &dcb30_360{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
