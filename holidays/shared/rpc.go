package shared

import (
	"net/rpc"

	api "github.com/magaldima/bizday/api"
)

type RPCClient struct{ client *rpc.Client }

func (m *RPCClient) IsHoliday(date api.Date) bool {
	var resp bool
	m.client.Call("Plugin.IsHoliday", map[string]interface{}{
		"date": date,
	}, &resp)
	return resp
}

func (m *RPCClient) HolidaysBetween(start api.Date, end api.Date) int32 {
	var resp int32
	m.client.Call("Plugin.HolidaysBetween", map[string]interface{}{
		"start": start,
		"end":   end,
	}, &resp)
	return resp
}

// RPCServer is the RPC server that RPCClient talks to, conforming to
// the requirements of net/rpc
type RPCServer struct {
	// This is the real implementation
	Impl Holiday
}

func (m *RPCServer) IsHoliday(args map[string]interface{}, resp *interface{}) bool {
	return m.Impl.IsHoliday(args["date"].(api.Date))
}

func (m *RPCServer) HolidaysBetween(args map[string]interface{}, resp *interface{}) int32 {
	return m.Impl.HolidaysBetween(args["start"].(api.Date), args["end"].(api.Date))
}
