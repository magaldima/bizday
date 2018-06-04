package holiday

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes"
	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
)

// GRPCClient is an implementation of Holiday that talks over RPC.
type GRPCClient struct{ client HolidayClient }

func (m *GRPCClient) IsHoliday(t time.Time) (bool, error) {
	timestamp, err := ptypes.TimestampProto(t)
	if err != nil {
		return false, err
	}
	resp, err := m.client.IsHoliday(context.Background(), timestamp)
	if err != nil {
		return false, err
	}
	return resp.GetOk(), nil
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

// GRPCServer is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl Holiday
}

func (m *GRPCServer) IsHoliday(ctx context.Context, timestamp *google_protobuf.Timestamp) (*BoolResponse, error) {
	t, err := ptypes.Timestamp(timestamp)
	if err != nil {
		return nil, err
	}
	ok, err := m.Impl.IsHoliday(t)
	if err != nil {
		return nil, err
	}
	return &BoolResponse{Ok: ok}, nil
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
