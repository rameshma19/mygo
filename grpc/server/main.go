package main

import (
	"context"
	"net"

	"github.com/mygo/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", "localhost:4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterCalcServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}

}

func (s *server) Add(ctx context.Context, req *proto.Request) (res *proto.Response, err error) {
	a, b := req.GetA(), req.GetB()
	result := a + b

	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, req *proto.Request) (res *proto.Response, err error) {
	a, b := req.GetA(), req.GetB()
	result := a * b

	return &proto.Response{Result: result}, nil
}

func (s *server) Divide(ctx context.Context, req *proto.Request) (res *proto.Response, err error) {
	a, b := req.GetA(), req.GetB()
	result := a / b

	return &proto.Response{Result: result}, nil
}
