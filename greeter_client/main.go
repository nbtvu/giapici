package main

import (
	"context"
	pb "github.com/nbtvu/giapici/greeter"
	"google.golang.org/grpc"
	"log"
)

func main() {

	//nlbAddr := "grpc-nlb-1f5862c5481c6dc3.elb.ap-southeast-1.amazonaws.com:80"
	nlbAddr := "localhost:8800"

	connCount := 1
	streamsPerConn := 50
	streamCount := connCount*streamsPerConn
	reqCount := 1000

	streams := []pb.Greeter_CounterClient{}
	for i := 0; i < connCount; i++ {
		// setup a connection to greeter server
		conn, err := grpc.Dial(nlbAddr, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}

		client := pb.NewGreeterClient(conn)
		for j := 0; j < streamsPerConn; j++ {
			stream, _ := client.Counter(context.Background())
			streams = append(streams, stream)
		}
	}


	for i := 0; i < reqCount; i++ {
		stream := streams[i%streamCount]
		req := pb.CounterRequest{
			Num: int64(i),
		}
		err := stream.Send(&req)
		if err != nil{
			log.Fatalf("failed to count, %v", err)
		}
	}

	for _, stream := range streams {
		_, _ = stream.CloseAndRecv()
	}
}
