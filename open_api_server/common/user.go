package common

import (
	pb "qyd/model/common"
	"golang.org/x/net/context"
	"log"
)

type User struct {
}

func (u User) SmsCode(ctx context.Context, in *pb.SmsCodeInput) (*pb.SmsCodeOutput, error) {
	log.Println("SmsCode run")
	return &pb.SmsCodeOutput{VerificationCode: "Hello again " + in.Phone}, nil
}
