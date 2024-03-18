package main

import (
	pb "greetings/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreetingsClient(conn)
	names := &pb.NamesList{
		Names: []*pb.Obj{
			{
				Name: "Max",
				Age:  30,
			},
			{
				Name: "Helen",
				Age:  40,
			},
		},
	}
	callHelloBidirectionalStream(c, names)
}
