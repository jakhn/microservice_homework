package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "app/protos/user_service"
)

func main() {
	conn, err := grpc.Dial("localhost:9002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	// resp, err := c.GetUserById(context.Background(), &pb.UserPrimeryKey{Id: 2})
	// if err != nil {
	// 	log.Println("error whiling get by id:", err.Error())
	// 	return
	// }

	// resp, err := c.Add(context.Background(), &pb.Sum{Num1: 1, Num2: 2})
	// if err != nil {
	// 	log.Println("error whiling adding nums:", err.Error())
	// 	return
	// }

	//================ SUBTRACTION ==================
	// resp, err := c.Subtraction(context.Background(), &pb.SubtReq{Num1: 1.1, Num2: 2.2})
	// if err != nil {
	// 	log.Println("error whiling subtracting nums:", err.Error())
	// 	return
	// }

	//================ MULTIPLICATION ==================
	// resp, err := c.Multiplication(context.Background(), &pb.MultReq{Num1: -1.1, Num2: 2.2})
	// if err != nil {
	// 	log.Println("error whiling multiplication nums:", err.Error())
	// 	return
	// }

	//================ DIVISION ==================
	// resp, err := c.Division(context.Background(), &pb.DivReq{Num1: -1.1, Num2: 2.2})
	// if err != nil {
	// 	log.Println("error whiling division nums:", err.Error())
	// 	return
	// }

	//================ SQUARE ROOT ==================
	// resp, err := c.SquareRoot(context.Background(), &pb.SqrtReq{Num1: 5.25})
	// if err != nil {
	// 	log.Println("error while finding square root of the num:", err.Error())
	// 	return
	// }

	//================ POWER ==================
	// resp, err := c.Power(context.Background(), &pb.PowReq{Num1: 2, Num2: 10})
	// if err != nil {
	// 	log.Println("error while finding power of the num:", err.Error())
	// 	return
	// }

	//================ ARRAY MIN  ==================
	// resp, err := c.ArrayMin(context.Background(), &pb.MnReq{Nums: []int32{2, 3, 4, 3, 42, 1}})
	// if err != nil {
	// 	log.Println("error while finding  min number of the given array", err.Error())
	// 	return
	// }

	// 	//================ ARRAY MAX  ==================
	resp, err := c.ArrayMax(context.Background(), &pb.MxReq{Nums: []int32{2, 3, 4, 3, 42, 1}})
	if err != nil {
		log.Println("error while finding  max number of the given array", err.Error())
		return
	}

	fmt.Println(resp)

}
