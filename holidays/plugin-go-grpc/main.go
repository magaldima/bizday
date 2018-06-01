package main

import (
	"github.com/hashicorp/go-plugin"
	api "github.com/magaldima/bizday/api"
	"github.com/magaldima/bizday/holidays/shared"
)

// Holiday is a real implementation of Holiday
type Holiday struct{}

func (Holiday) IsHoliday(api.Date) bool {
	return false
}

func (Holiday) HolidaysBetween(api.Date, api.Date) int32 {
	return 0
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			"us-holiday": &shared.HolidayPlugin{Impl: &Holiday{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
