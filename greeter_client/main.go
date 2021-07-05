package main

import (
	"context"
	pb "github.com/nbtvu/giapici/greeter"
	"google.golang.org/grpc"
	"log"
)

func main() {

	nlbAddr := "localhost:8800"
	clients := []pb.GreeterClient{}
	for i := 0; i < 1; i++ {
		// setup a connection to greeter server
		conn, err := grpc.Dial(nlbAddr, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer func() {
			_ = conn.Close()
		}()

		client := pb.NewGreeterClient(conn)
		clients = append(clients, client)
	}

	for i := 0; i < 100; i++ {
		client := clients[i%1]
		res, err := client.Counter(context.Background(), &pb.CounterRequest{Num: int64(i)})
		if err != nil{
			log.Fatalf("failed to count, %v", err)
		}
		log.Println("result: ", res.ResNum, ", ", res.Ip)
	}
}
