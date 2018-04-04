package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"grpc_service_stream/open_api_server/common"
	"google.golang.org/grpc/reflection"

	pb "grpc_service_stream/model/common"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &common.User{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Println("server running ...")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
