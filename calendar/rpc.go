package calendar

import (
	"net/rpc"

	api "github.com/magaldima/bizday/api"
)

type RPCClient struct{ client *rpc.Client }

func (m *RPCClient) DaysInYear(date api.Date) int32 {
	var resp int32
	m.client.Call("Plugin.DaysInYear", map[string]interface{}{
		"date": date,
	}, &resp)
	return resp
}

func (m *RPCClient) DaysInMonth(date api.Date) int32 {
	var resp int32
	m.client.Call("Plugin.DaysInMonth", map[string]interface{}{
		"date": date,
	}, &resp)
	return resp
}

func (m *RPCClient) Delta(start api.Date, end api.Date) int32 {
	var resp int32
	m.client.Call("Plugin.Delta", map[string]interface{}{
		"start": start,
		"end":   end,
	}, &resp)
	return resp
}

func (m *RPCClient) Alpha(start api.Date, end api.Date) float64 {
	var resp float64
	m.client.Call("Plugin.Alpha", map[string]interface{}{
		"start": start,
		"end":   end,
	}, &resp)
	return resp
}

// RPCServer is the RPC server that RPCClient talks to, conforming to
// the requirements of net/rpc
type RPCServer struct {
	// This is the real implementation
	Impl Calendar
}

func (m *RPCServer) DaysInYear(args map[string]interface{}, resp *interface{}) int32 {
	return m.Impl.DaysInYear(args["date"].(api.Date))
}

func (m *RPCServer) DaysInMonth(args map[string]interface{}, resp *interface{}) int32 {
	return m.Impl.DaysInMonth(args["date"].(api.Date))
}

func (m *RPCServer) Delta(args map[string]interface{}, resp *interface{}) int32 {
	return m.Impl.Delta(args["start"].(api.Date), args["end"].(api.Date))
}

func (m *RPCServer) Alpha(args map[string]interface{}, resp *interface{}) float64 {
	return m.Impl.Alpha(args["start"].(api.Date), args["end"].(api.Date))
}
