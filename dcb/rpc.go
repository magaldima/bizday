package dcb

import (
	"net/rpc"
	"time"
)

type RPCClient struct{ client *rpc.Client }

func (m *RPCClient) DaysInYear(t time.Time) (int32, error) {
	var resp int32
	err := m.client.Call("Plugin.DaysInYear", map[string]interface{}{
		"time": t,
	}, &resp)
	if err != nil {
		return 0, err
	}
	return resp, nil
}

func (m *RPCClient) DaysInMonth(t time.Time) (int32, error) {
	var resp int32
	err := m.client.Call("Plugin.DaysInMonth", map[string]interface{}{
		"time": t,
	}, &resp)
	if err != nil {
		return 0, err
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

func (m *RPCClient) Alpha(start, end time.Time) (float64, error) {
	var resp float64
	err := m.client.Call("Plugin.Alpha", map[string]interface{}{
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
	Impl DayCountBasis
}

func (m *RPCServer) DaysInYear(args map[string]interface{}, resp *interface{}) (int32, error) {
	return m.Impl.DaysInYear(args["time"].(time.Time))
}

func (m *RPCServer) DaysInMonth(args map[string]interface{}, resp *interface{}) (int32, error) {
	return m.Impl.DaysInMonth(args["time"].(time.Time))
}

func (m *RPCServer) Delta(args map[string]interface{}, resp *interface{}) (int32, error) {
	return m.Impl.Delta(args["start"].(time.Time), args["end"].(time.Time))
}

func (m *RPCServer) Alpha(args map[string]interface{}, resp *interface{}) (float64, error) {
	return m.Impl.Alpha(args["start"].(time.Time), args["end"].(time.Time))
}
