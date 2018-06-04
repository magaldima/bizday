package holiday

import (
	"context"
	"net/rpc"
	"time"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "HOLIDAY_PLUGIN",
	MagicCookieValue: "holiday",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"US": &HolidayPlugin{},
	//todo: add more plugin here for all different types of calendars
}

// Holiday is the interface that we're exposing as a plugin.
type Holiday interface {
	IsHoliday(time.Time) (bool, error)
	Delta(time.Time, time.Time) (int32, error)
}

// HolidayPlugin is the implementation of plugin.Plugin so we can serve/consume this.
// We also implement GRPCPlugin so that this plugin can be served over
// gRPC.
type HolidayPlugin struct {
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl Holiday
}

func (p *HolidayPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &RPCServer{Impl: p.Impl}, nil
}

func (*HolidayPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &RPCClient{client: c}, nil
}

func (p *HolidayPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	RegisterHolidayServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *HolidayPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: NewHolidayClient(c)}, nil
}
