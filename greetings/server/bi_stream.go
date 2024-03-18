package main

import (
	pb "greetings/proto"
	"io"
	"log"
	"strconv"
)

func (s *GreetingsServer) SayHelloBidirectionalStreaming(stream pb.Greetings_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got Request with name: %v, age: %v", req.GetName(), req.GetAge())
		res := &pb.HelloResponse{
			Message: "Hello " + req.GetName() + " of age " + strconv.Itoa(int(req.GetAge())),
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
