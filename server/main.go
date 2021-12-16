package main

import (
	"context"
	"helloworld/interceptor"
	"log"
	"net"
	"time"

	pb "helloworld/protobuf"

	"google.golang.org/grpc"
)

const (
	port = ":3000"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	time.Sleep(2 * time.Second)
	return &pb.HelloReply{
		Message: "Hello " + in.GetName(),
	}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	time.Sleep(2 * time.Second)
	return &pb.HelloReply{
		Message: "Hello again " + in.GetName(),
	}, nil
}

func main() {
	var err error
	var listener net.Listener

	listener, err = net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			interceptor.UnaryCacheServerInterceptor,
		),
	)
	pb.RegisterGreeterServer(s, &server{})

	log.Println("starting grpc server")
	if err = s.Serve(listener); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
