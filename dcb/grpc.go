package dcb

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes"
	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
)

// GRPCClient is an implementation of Holiday that talks over RPC.
type GRPCClient struct{ client DayCountBasisClient }

func (m *GRPCClient) DaysInYear(t time.Time) (int32, error) {
	timestamp, err := ptypes.TimestampProto(t)
	if err != nil {
		return 0, err
	}
	resp, err := m.client.DaysInYear(context.Background(), timestamp)
	if err != nil {
		// todo: add err to the interface
		return 0, err
	}
	return resp.GetValue(), nil
}

func (m *GRPCClient) DaysInMonth(t time.Time) (int32, error) {
	timestamp, err := ptypes.TimestampProto(t)
	if err != nil {
		return 0, err
	}
	resp, err := m.client.DaysInMonth(context.Background(), timestamp)
	if err != nil {
		// todo: add err to the interface
		return 0, err
	}
	return resp.GetValue(), nil
}

func (m *GRPCClient) Delta(start, end time.Time) (int32, error) {
	startTimestamp, err := ptypes.TimestampProto(start)
	if err != nil {
		return 0, err
	}
	endTimestamp, err := ptypes.TimestampProto(end)
	if err != nil {
		return 0, err
	}
	resp, err := m.client.Delta(context.Background(), &BinaryDateRequest{
		Start: startTimestamp,
		End:   endTimestamp,
	})
	if err != nil {
		return 0, err
	}
	return resp.GetValue(), nil
}

func (m *GRPCClient) Alpha(start, end time.Time) (float64, error) {
	startTimestamp, err := ptypes.TimestampProto(start)
	if err != nil {
		return 0, err
	}
	endTimestamp, err := ptypes.TimestampProto(end)
	if err != nil {
		return 0, err
	}
	resp, err := m.client.Alpha(context.Background(), &BinaryDateRequest{
		Start: startTimestamp,
		End:   endTimestamp,
	})
	if err != nil {
		return 0, err
	}
	return resp.GetValue(), nil
}

// GRPCServer is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl DayCountBasis
}

func (m *GRPCServer) DaysInYear(ctx context.Context, timestamp *google_protobuf.Timestamp) (*NumberOfDaysResponse, error) {
	t, err := ptypes.Timestamp(timestamp)
	if err != nil {
		return nil, err
	}
	days, err := m.Impl.DaysInYear(t)
	if err != nil {
		return nil, err
	}
	return &NumberOfDaysResponse{Value: days}, nil
}

func (m *GRPCServer) DaysInMonth(ctx context.Context, timestamp *google_protobuf.Timestamp) (*NumberOfDaysResponse, error) {
	t, err := ptypes.Timestamp(timestamp)
	if err != nil {
		return nil, err
	}
	days, err := m.Impl.DaysInMonth(t)
	if err != nil {
		return nil, err
	}
	return &NumberOfDaysResponse{Value: days}, nil
}

func (m *GRPCServer) Delta(ctx context.Context, req *BinaryDateRequest) (*NumberOfDaysResponse, error) {
	start, err := ptypes.Timestamp(req.Start)
	if err != nil {
		return nil, err
	}
	end, err := ptypes.Timestamp(req.End)
	if err != nil {
		return nil, err
	}
	days, err := m.Impl.Delta(start, end)
	if err != nil {
		return nil, err
	}
	return &NumberOfDaysResponse{Value: days}, nil
}

func (m *GRPCServer) Alpha(ctx context.Context, req *BinaryDateRequest) (*DayCountFractionResponse, error) {
	start, err := ptypes.Timestamp(req.Start)
	if err != nil {
		return nil, err
	}
	end, err := ptypes.Timestamp(req.End)
	if err != nil {
		return nil, err
	}
	frac, err := m.Impl.Alpha(start, end)
	if err != nil {
		return nil, err
	}
	return &DayCountFractionResponse{Value: frac}, nil
}
