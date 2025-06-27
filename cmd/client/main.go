package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "qfzack/go-web-starter/api/proto/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "server address")
	name = flag.String("name", "qfzack", "request name")
)

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMyRpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.Request{Name: *name})
	if err != nil {
		log.Fatalf("could not access: %v", err)
	}
	log.Printf("Accessing: %s", r.GetMessage())
}
