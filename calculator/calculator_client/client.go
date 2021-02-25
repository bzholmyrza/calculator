package main

import (
	"context"
	"fmt"
	"gitlab.com/bzholmyrza/calculator/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
)


func main()  {
	fmt.Println("Hello I'm a client")

	conn,err := grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v",err)
	}
	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)

	doPrimeNumberDecomposition(c)
	//doComputeAverage(c)
}

func doPrimeNumberDecomposition(c calculatorpb.CalculatorServiceClient)  {
	ctx := context.Background()
	request := &calculatorpb.CalculatorRequest{NumberRequest: &calculatorpb.Calculator{
		Number: 60,
	}}

	responce,err := c.PrimeNumberDecomposition(ctx,request)
	if err != nil {
		log.Fatalf("error while calling Greet RPC $v",err)
	}
	log.Printf("response from Greet:%v",responce.Result)
}

func doComputeAverage(c calculatorpb.CalculatorServiceClient)  {
	ctx := context.Background()
	request := &calculatorpb.CalculatorComputeAverageRequest{NumberRequest: &calculatorpb.CalculatorNumbers{
		Number: "1, 2, 3, 4",
	}}

	responce,err := c.ComputeAverage(ctx,request)
	if err != nil {
		log.Fatalf("error while calling Greet RPC $v",err)
	}
	log.Printf("response from Greet:%v",responce.Result)
}

