package server

import (
	"context"
	pb "service/examples/pkg/pb/v1"
)

func (e *handler) ReadEcho(ctx context.Context, req *pb.FetchEchoRequest) (*pb.Echo, error) {
	return &pb.Echo{Id: req.GetId()}, nil
}

// Auth一般都是随着账号默认创建一个，其他的都是绑定
func (e *handler) FetchEcho(ctx context.Context, req *pb.FetchEchoRequest) (*pb.EchoResponse, error) {
	data := []*pb.Echo{
		{Id: 1, Name: "name1"},
		{Id: 2, Name: "name2"},
	}
	return &pb.EchoResponse{Count: 2, Data: data}, nil
}
