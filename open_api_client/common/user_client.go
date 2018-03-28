package main

import (
	"os"
	"log"
	"time"

	"google.golang.org/grpc"
	"golang.org/x/net/context"
	pb "qyd/proto/common"
)

const (
	address      = "localhost:50051"
	defaultPhone = "13120168539"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserClient(conn)

	// Contact the server and print out its response.
	phone := defaultPhone
	if len(os.Args) > 1 {
		phone = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SmsCode(ctx, &pb.SmsCodeInput{Phone: phone})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.VerificationCode)
}
