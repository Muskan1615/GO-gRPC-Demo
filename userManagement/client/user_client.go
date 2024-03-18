package main

import (
	"context"
	"log"
	"time"
	pb "usermanagement/proto"

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

	c := pb.NewUserManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	newUsers := make(map[string]int32)
	newUsers["Max"] = 45
	newUsers["Spenser"] = 30
	for name, age := range newUsers {
		res, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Printf("Could not create user: %v", err)
		}
		log.Printf(`
		User Details:
		Name: %s
		Age: %d
		Id: %d`, res.GetName(), res.GetAge(), res.GetId())
	}

}
