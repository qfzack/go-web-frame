package handler

import (
	"context"
	"fmt"
	"log"
	pb "qfzack/grpc-demo/api/proto/server"
)

// custom Server struct inplements the MyRpcServer interface
type RPCServer struct {
	BaseHandler
	pb.UnimplementedMyRpcServer
}

func NewRPCServer() *RPCServer {
	return &RPCServer{}
}

// override the func SayHello defined in the gRPC
func (s *RPCServer) SayHello(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Printf("request recieved: %v", req.Name)
	return &pb.Response{
		Message: fmt.Sprintf("%s from server", req.Name),
	}, nil
}
