package main

import (
	pb "github.com/nbtvu/giapici/greeter"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *greeterServer) Counter(stream pb.Greeter_CounterServer) error {
	log.Println("Counter function")
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// Close the connection and return the response to the client
			return stream.SendAndClose(&pb.CounterResponse{
				ResNum: int64(count),
				Ip: "not-implemented",
			})
		}

		if err != nil {
			log.Fatalf("Error when reading client request stream: %v", err)
		}

		num := req.GetNum()
		log.Println("Received counter with num: ", num)
		count++
	}
	return nil
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
