package main

import (
	"context"
	pb "github.com/nbtvu/giapici/greeter"
	"google.golang.org/grpc"
	"log"
	"net"
)

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *greeterServer) Counter(ctx context.Context, request *pb.CounterRequest) (*pb.CounterResponse, error) {
	log.Printf("Received: %v", request.GetNum())
	return &pb.CounterResponse{
		ResNum: -request.GetNum(),
		Ip: "not-implemented",
	}, nil
}


func main()  {
	lis, err := net.Listen("tcp",":8800")
	if err != nil {
		log.Fatalf("failed to listen, %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &greeterServer{})

	log.Println("starting server now....")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server, %v", err)
	}
}



//protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/chat.proto
