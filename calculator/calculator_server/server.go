package main

import (
	"context"
	"fmt"
	"gitlab.com/bzholmyrza/calculator/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

type Server struct{
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (s *Server) PrimeNumberDecomposition(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	fmt.Printf("PrimeNumberDecomposition function was invoked with %v \n",req)
	number := int64(req.GetNumberRequest().Number)
	var div int64 = 2
	result := " "
	for number > 1 {
		for number % div == 0 {
			result += " " + strconv.FormatInt(div, 10)
			number /= div
		}
		div++
	}
	res := &calculatorpb.CalculatorResponse{Result: result}
	return res, nil
}


/*func (s *Server) ComputeAverage(ctx context.Context, req *calculatorpb.CalculatorComputeAverageRequest) (*calculatorpb.CalculatorComputeAverageResponse, error) {
	fmt.Printf("ComputeAverage function was invoked with %v \n",req)
	numbers := strings.Split(req.NumberRequest.Number, ",")
	var result int64 = 0
	for i:= 0; i < len(numbers); i++ {
		result = result + strconv.Atoi(numbers[i])
	}
	res := &calculatorpb.CalculatorComputeAverageResponse{Result: result}
	return res, nil
}*/

func main()  {
	l,err := net.Listen("tcp","0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v",err)
	}

	s:=grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &Server{})
	if err := s.Serve(l);err!=nil {
		log.Fatalf("failed to serve:%v",err)
	}
}

