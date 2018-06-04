package main

import (
	"time"

	"github.com/hashicorp/go-plugin"
	"github.com/magaldima/bizday/holiday"
)

// Holiday is a real implementation of Holiday
type Holiday struct{}

func (Holiday) IsHoliday(d time.Time) (bool, error) {
	return false, nil
}

func (Holiday) Delta(d1, d2 time.Time) (int32, error) {
	return 0, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: holiday.Handshake,
		Plugins: map[string]plugin.Plugin{
			"US": &holiday.HolidayPlugin{Impl: &Holiday{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
