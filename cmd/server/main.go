package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "qfzack/grpc-demo/api/proto/server"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "server port")
)

type server struct {
	pb.UnimplementedMyRpcServer
}

func (s *server) SayHello(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Printf("request recieved: %v", req.Name)
	return &pb.Response{
		Message: fmt.Sprintf("%s from server", req.Name),
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMyRpcServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
