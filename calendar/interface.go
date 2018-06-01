package calendar

import (
	"context"
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	api "github.com/magaldima/bizday/api"
	"google.golang.org/grpc"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "CALENDAR_PLUGIN",
	MagicCookieValue: "calendar",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"30_360": &CalendarPlugin{},
	//todo: add more plugin here for all different types of calendars
}

// Calendar is the interface that we're exposing as a plugin.
type Calendar interface {
	DaysInYear(api.Date) int32
	DaysInMonth(api.Date) int32
	Delta(api.Date, api.Date) int32
	Alpha(api.Date, api.Date) float64
}

// CalendarPlugin is the implementation of plugin.Plugin so we can serve/consume this.
// We also implement GRPCPlugin so that this plugin can be served over
// gRPC.
type CalendarPlugin struct {
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl Calendar
}

func (p *CalendarPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &RPCServer{Impl: p.Impl}, nil
}

func (*CalendarPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &RPCClient{client: c}, nil
}

func (p *CalendarPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	RegisterCalendarServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *CalendarPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: NewCalendarClient(c)}, nil
}
