package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	pb "usermanagement/proto"

	"google.golang.org/grpc"
)

const port = ":50051"

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, newUser *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", newUser.GetName())
	userId := int32(rand.Intn(1000))
	return &pb.User{Name: newUser.GetName(), Age: newUser.GetAge(), Id: userId}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, &UserManagementServer{})
	log.Printf("Server listening on port %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
