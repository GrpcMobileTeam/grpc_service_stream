package common

import (
	pb "grpc_service_stream/model/common"
	"golang.org/x/net/context"
	"log"
	"fmt"
	"time"
)

type User struct {
}

func (u User) StreamSmsCode(rect *pb.SmsCodeInput, stream pb.User_StreamSmsCodeServer) error {
	log.Print("StreamSmsCode is running")
	for i := 1; i <= 5; i++ {
		sendBody := fmt.Sprintf("hello client, I'm going to start run %d times;", i)
		if err := stream.Send(&pb.SmsCodeOutput{VerificationCode: sendBody}); err != nil {
			return err
		}
		log.Print(sendBody)
		time.Sleep(time.Second * 2)
	}

	return nil
}

func (u User) SmsCode(ctx context.Context, in *pb.SmsCodeInput) (*pb.SmsCodeOutput, error) {
	log.Println("SmsCode run")
	return &pb.SmsCodeOutput{VerificationCode: "Hello again " + in.Phone}, nil
}
