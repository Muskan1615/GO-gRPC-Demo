package main

import (
	pb "greetings/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

const port = ":50051"

type GreetingsServer struct {
	pb.UnimplementedGreetingsServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreetingsServer(s, &GreetingsServer{})
	log.Printf("Server listening on port %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Printf("Failed to serve: %v", err)
	}
}
