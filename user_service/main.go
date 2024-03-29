package main

import (
	"context"
	"errors"
	"fmt"
	"net"

	"google.golang.org/grpc"

	pb "app/protos/user_service"
)

type user_server struct {
	*pb.UnimplementedUserServiceServer
}

func (u *user_server) GetUserById(ctx context.Context, req *pb.UserPrimeryKey) (*pb.User, error) {
	users := []pb.User{
		{
			Id:       1,
			FullName: "Shaxboz Norbekov",
			Age:      24,
		},
		{
			Id:       2,
			FullName: "Jahongir Normurodov",
			Age:      25,
		},
		{
			Id:       3,
			FullName: "Samandar Foziljonov",
			Age:      20,
		},
		{
			Id:       4,
			FullName: "Moxirbek Abduvaliev",
			Age:      21,
		},
	}

	for _, user := range users {
		if req.Id == user.Id {
			return &user, nil
		}
	}

	return nil, errors.New("Not found user")
}

func (s *user_server) Add(ctx context.Context, req *pb.Sum) (*pb.SumA_B, error) {
	sum_of := req.Num1 + req.Num2
	return &pb.SumA_B{SumOf: sum_of}, nil
}

func (s *user_server) Division(ctx context.Context, req *pb.DivReq) (*pb.DivResp, error) {
	var division float32
	if req.Num1 == 0 {
		return &pb.DivResp{Num3: 0}, nil	
	}
	if req.Num2 == 0 {
		return nil, errors.New("Bo'luvchi son 0 ga teng! Boshqa son kiriting!")
	}
	division = req.Num1 / req.Num2
	return &pb.DivResp{Num3: division}, nil
}

func (s *user_server) Multiplication(ctx context.Context, req *pb.MultReq) (*pb.MultResp, error) {
	var multiplication float32
	multiplication = req.Num1 * req.Num2
	return &pb.MultResp{Num3: multiplication}, nil
}

func (s *user_server) Subtraction(ctx context.Context, req *pb.SubtReq) (*pb.SubtResp, error) {
	var subtraction float32
	subtraction = req.Num1 - req.Num2
	return &pb.SubtResp{Num3: subtraction}, nil
}

func (s *user_server) SquareRoot(ctx context.Context, req *pb.SqrtReq) (*pb.SqrtResp, error) {
	var (
		start     float32 = 0
		end       float32 = req.Num1
		mid       float32
		ans       float32
		increment float32 = 0.1
	)

	if req.Num1 <= 0 {
		return nil, errors.New("0 dan kotta son kiriting")
	}
	for start <= end {
		mid = (start + end) / 2

		if mid*mid == req.Num1 {
			ans = mid
			break
		}

		if mid*mid < req.Num1 {
			ans = start
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	for i := 0; i < 5; i++ {
		for ans*ans <= req.Num1 {
			ans += increment
		}
		ans = ans - increment
		increment = increment / 10
	}
	return &pb.SqrtResp{Num3: ans}, nil

}

func (s *user_server) Power(ctx context.Context, req *pb.PowReq) (*pb.PowResp, error) {
	var result float32 = 1
	if req.Num2 == 0 {
		return &pb.PowResp{Num3: 1}, nil
	}
	if req.Num2 < 0 {
		return nil, errors.New("Power should be positive number!")

	}
	for req.Num2 != 0 {
		result *= req.Num1
		req.Num2--
	}

	return &pb.PowResp{Num3: result}, nil
}

func (s *user_server) ArrayMin(ctx context.Context, req *pb.MnReq) (*pb.MnResp, error) {
	result := pb.MnResp{}
	result.Num = req.Nums[0]
	for i := 0; i < len(req.Nums); i++ {
		if result.Num > req.Nums[i] {
			result.Num = req.Nums[i]
		}
	}

	return &result, nil
}

func (s *user_server) ArrayMax(ctx context.Context, req *pb.MxReq) (*pb.MxResp, error) {
	result := pb.MxResp{}
	result.Num = req.Nums[0]
	for i := 0; i < len(req.Nums); i++ {
		if result.Num < req.Nums[i] {
			result.Num = req.Nums[i]
		}
	}

	return &result, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9002")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &user_server{})

	fmt.Println("Listening :9002...")
	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
