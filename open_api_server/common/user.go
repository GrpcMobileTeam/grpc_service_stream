package common

import (
	pb "qyd/model/common"
	"golang.org/x/net/context"
)

type User struct {
}

func (u User) SmsCode(ctx context.Context, in *pb.SmsCodeInput) (*pb.SmsCodeOutput, error) {
	return &pb.SmsCodeOutput{VerificationCode: "Hello again " + in.Phone}, nil
}
