package holiday

import (
	"net/rpc"
	"time"
)

type RPCClient struct{ client *rpc.Client }

func (m *RPCClient) IsHoliday(t time.Time) (bool, error) {
	var resp bool
	err := m.client.Call("Plugin.IsHoliday", map[string]interface{}{
		"time": t,
	}, &resp)
	if err != nil {
		return false, err
	}
	return resp, nil
}

func (m *RPCClient) Delta(start, end time.Time) (int32, error) {
	var resp int32
	err := m.client.Call("Plugin.Delta", map[string]interface{}{
		"start": start,
		"end":   end,
	}, &resp)
	if err != nil {
		return 0, err
	}
	return resp, nil
}

// RPCServer is the RPC server that RPCClient talks to, conforming to
// the requirements of net/rpc
type RPCServer struct {
	// This is the real implementation
	Impl Holiday
}

func (m *RPCServer) IsHoliday(args map[string]interface{}, resp *interface{}) (bool, error) {
	return m.Impl.IsHoliday(args["time"].(time.Time))
}

func (m *RPCServer) Delta(args map[string]interface{}, resp *interface{}) (int32, error) {
	return m.Impl.Delta(args["start"].(time.Time), args["end"].(time.Time))
}
