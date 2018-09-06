package main

import (
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/aren55555/aren55555/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct {
	servers []pb.Chat_StreamServer
}

var connections map[string]*pb.Chat_StreamServer

func (s *server) Stream(cs pb.Chat_StreamServer) error {
	s.servers = append(s.servers, cs)

	for {
		in, err := cs.Recv()
		fmt.Println("\n\n\nReceived value")
		if err == io.EOF {
			fmt.Println("EOF")
			return nil
		}
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		fmt.Println(in)

		for _, server := range s.servers {
			if server == cs {
				continue
			}
			server.Send(in)
		}
	}
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
