package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/aren55555/aren55555/protos"
	"google.golang.org/grpc"
)

var (
	flagServer  = flag.String("server", "localhost:50051", "the address of the server")
	flagMessage = flag.String("message", "test", "the message to repeadedlty send")
)

func main() {
	flag.Parse()
	fmt.Println("aren55555 client")

	// Set up a connection to the server.
	conn, err := grpc.Dial(*flagServer, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewChatClient(conn)

	// Contact the server and print out its response.
	streamClient, err := client.Stream(context.Background())
	if err != nil {
		log.Fatalf("could not get stream client: %v", err)
	}

	go func() {
		for {
			in, err := streamClient.Recv()
			fmt.Println("\n\n\nReceived value")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println(in)
		}
	}()

	waitc := make(chan struct{})

	msg := &pb.Message{Body: *flagMessage}
	go func() {
		for {
			log.Println("Sleeping...")
			time.Sleep(2 * time.Second)
			log.Println("Sending msg...")
			streamClient.Send(msg)
		}
	}()
	<-waitc
}
