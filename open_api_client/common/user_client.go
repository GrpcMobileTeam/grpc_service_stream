package main

import (
	"os"
	"log"
	"time"

	"google.golang.org/grpc"
	"golang.org/x/net/context"
	pb "qyd/model/common"
	"io"
)

const (
	address      = "localhost:50051"
	defaultPhone = "13120168539"
)

func callSmsCode(c pb.UserClient) {
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

func callStreamSmsCode(c pb.UserClient, rect *pb.SmsCodeInput) {
	log.Printf("Looking for features within %v", rect)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := c.StreamSmsCode(ctx, rect)
	if err != nil {
		log.Fatalf("%v.ListFeatures(_) = _, %v", c, err)
	}
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", c, err)
		}
		log.Println(feature)
	}
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserClient(conn)

	// simple call func
	go callSmsCode(c)

	// stream call func
	go callStreamSmsCode(c, &pb.SmsCodeInput{Phone: "131"})

	select {}
}
