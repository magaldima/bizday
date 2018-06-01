package shared

import (
	"context"

	api "github.com/magaldima/bizday/api"
	"github.com/magaldima/bizday/holidays"
)

// GRPCClient is an implementation of Holiday that talks over RPC.
type GRPCClient struct{ client holidays.HolidayClient }

func (m *GRPCClient) IsHoliday(date api.Date) bool {
	resp, err := m.client.IsHoliday(context.Background(), &date)
	if err != nil {
		// add err to the interface
		return false
	}
	return resp.GetOk()
}

func (m *GRPCClient) HolidaysBetween(start api.Date, end api.Date) int32 {
	resp, err := m.client.HolidaysBetween(context.Background(), &holidays.BinaryDateRequest{
		Start: &start,
		End:   &end,
	})
	if err != nil {
		return 0
	}
	return resp.GetDays()
}

// GRPCServer is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl Holiday
}

func (m *GRPCServer) IsHoliday(ctx context.Context, date *api.Date) (*holidays.BoolResponse, error) {
	ok := m.Impl.IsHoliday(*date)
	return &holidays.BoolResponse{Ok: ok}, nil
}

func (m *GRPCServer) HolidaysBetween(ctx context.Context, req *holidays.BinaryDateRequest) (*api.NumberOfDaysResponse, error) {
	days := m.Impl.HolidaysBetween(*req.Start, *req.End)
	return &api.NumberOfDaysResponse{Days: days}, nil
}
