package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "helloworld/protobuf"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:3000"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		e := conn.Close()
		if e != nil {
			log.Fatal(e)
		}
	}()

	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	hello, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greeting: %v", err)
	}
	log.Printf("Greeting: %s", hello.GetMessage())

	helloAgain, err := c.SayHelloAgain(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greeting again: %v", err)
	}
	log.Printf("Greeting: %s", helloAgain.GetMessage())
}
