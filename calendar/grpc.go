package calendar

import (
	"context"

	api "github.com/magaldima/bizday/api"
)

// GRPCClient is an implementation of Holiday that talks over RPC.
type GRPCClient struct{ client CalendarClient }

func (m *GRPCClient) DaysInYear(date api.Date) int32 {
	resp, err := m.client.DaysInYear(context.Background(), &date)
	if err != nil {
		// todo: add err to the interface
		return 0
	}
	return resp.GetDays()
}

func (m *GRPCClient) DaysInMonth(date api.Date) int32 {
	resp, err := m.client.DaysInMonth(context.Background(), &date)
	if err != nil {
		// todo: add err to the interface
		return 0
	}
	return resp.GetDays()
}

func (m *GRPCClient) Delta(start api.Date, end api.Date) int32 {
	resp, err := m.client.Delta(context.Background(), &BinaryDateRequest{
		Start: &start,
		End:   &end,
	})
	if err != nil {
		return 0
	}
	return resp.GetDays()
}

func (m *GRPCClient) Alpha(start api.Date, end api.Date) float64 {
	resp, err := m.client.Alpha(context.Background(), &BinaryDateRequest{
		Start: &start,
		End:   &end,
	})
	if err != nil {
		return 0
	}
	return resp.GetValue()
}

// GRPCServer is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl Calendar
}

func (m *GRPCServer) DaysInYear(ctx context.Context, date *api.Date) (*api.NumberOfDaysResponse, error) {
	days := m.Impl.DaysInYear(*date)
	return &api.NumberOfDaysResponse{Days: days}, nil
}

func (m *GRPCServer) DaysInMonth(ctx context.Context, date *api.Date) (*api.NumberOfDaysResponse, error) {
	days := m.Impl.DaysInMonth(*date)
	return &api.NumberOfDaysResponse{Days: days}, nil
}

func (m *GRPCServer) Delta(ctx context.Context, req *BinaryDateRequest) (*api.NumberOfDaysResponse, error) {
	days := m.Impl.Delta(*req.Start, *req.End)
	return &api.NumberOfDaysResponse{Days: days}, nil
}

func (m *GRPCServer) Alpha(ctx context.Context, req *BinaryDateRequest) (*DayCountFractionResponse, error) {
	frac := m.Impl.Alpha(*req.Start, *req.End)
	return &DayCountFractionResponse{Value: frac}, nil
}
