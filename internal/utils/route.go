package utils

import (
	"context"
	"github.com/osamikoyo/floof-db/internal/data"
	"github.com/osamikoyo/floof-db/pkg/pb"
	"google.golang.org/grpc"
)

type DataRouter struct {
	pb.UnimplementedFloofyServer
	Ring data.RingHandler
}

func (d DataRouter) Create(ctx context.Context, in *pb.AddReq, opts ...grpc.CallOption) (*pb.Response, error) {
	d.Ring.Storage.Create(in.Key, in.Value)
	return &pb.Response{Message: "success"}, nil
}

func (d DataRouter) Get(ctx context.Context, in *pb.GetReq, opts ...grpc.CallOption) (*pb.GetResp, error) {
	value, err := d.Ring.Storage.Get(in.Key)
	if err != nil {
		return nil, err
	}

	return &pb.GetResp{
		Value: value,
	}, nil
}

func (d DataRouter) Update(ctx context.Context, in *pb.UpdateReq, opts ...grpc.CallOption) (*pb.Response, error) {
	d.Ring.Storage.Update(in.Key, in.NewValue)

	return &pb.Response{Message: "success"}, nil
}
