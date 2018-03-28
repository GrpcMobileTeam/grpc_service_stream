package main

import (
	pb "qyd/model/common"
	"google.golang.org/grpc"
	"net"
	"qyd/open_api_server/common"
	"google.golang.org/grpc/reflection"
	"log"
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
