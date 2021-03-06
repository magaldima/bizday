package dcb

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
	MagicCookieKey:   "DCB_PLUGIN",
	MagicCookieValue: "dcb",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"30_360": &DcbPlugin{},
	//todo: add more plugin here for all different types of day count basis implementations
}

// DayCountBasis is the interface that we're exposing as a plugin.
type DayCountBasis interface {
	DaysInYear(time.Time) (int32, error)
	DaysInMonth(time.Time) (int32, error)
	Delta(time.Time, time.Time) (int32, error)
	Alpha(time.Time, time.Time) (float64, error)
}

// DcbPlugin is the implementation of plugin.Plugin so we can serve/consume this.
// We also implement GRPCPlugin so that this plugin can be served over
// gRPC.
type DcbPlugin struct {
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl DayCountBasis
}

func (p *DcbPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &RPCServer{Impl: p.Impl}, nil
}

func (*DcbPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &RPCClient{client: c}, nil
}

func (p *DcbPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	RegisterDayCountBasisServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *DcbPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: NewDayCountBasisClient(c)}, nil
}
