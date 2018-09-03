package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/aren55555/aren55555/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Stream(cs pb.Chat_StreamServer) error {
	fmt.Println("streaming")

	m := &pb.Message{Body: "server up!"}
	cs.Send(m)

	r, err := cs.Recv()
	if err != nil {
		panic(err)
	}
	fmt.Println(r)

	return nil
}

func main() {
	fmt.Println("aren55555 server")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChatServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
