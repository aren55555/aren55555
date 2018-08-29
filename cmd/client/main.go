package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/aren55555/aren55555/protos"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	fmt.Println("aren55555 client")

	m := &pb.Message{Body: "Hello World"}
	fmt.Println(m)

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChatClient(conn)

	fmt.Println(c)

	// Contact the server and print out its response.
	ctx := context.Background()
	sc, err := c.Stream(ctx)
	if err != nil {
		log.Fatalf("could not get stream client: %v", err)
	}
	fmt.Println(sc)

	sc.Send(m)

}
