package main

import (
	"context"
	"net"

	"github.com/gunsluo/go-example/grpc/gateway/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	sAddress = "0.0.0.0:19000"
)

type Service struct {
}

func (s *Service) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "hello, " + req.Name,
	}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &Service{})

	listener, err := net.Listen("tcp", sAddress)
	if err != nil {
		panic(err)
	}

	logrus.WithField("addr", sAddress).Println("Starting server")
	server.Serve(listener)
}
