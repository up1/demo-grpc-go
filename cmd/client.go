package main

import (
	"context"
	"fmt"
	"time"

	pb "demo/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func CallSayHi(client pb.HelloWorldClient) (*pb.HelloResponse, error) {
	// Timeout 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	hello := pb.Hello{
		Id:      1,
		Message: "Hello world ...",
	}
	resp, err := client.SayHi(ctx, &hello)
	statusCode := status.Code(err)
	if statusCode != codes.OK {
		return nil, err
	}
	fmt.Printf("Response: %d, statusCode: %s\n", resp.Id, statusCode.String())
	return resp, err
}

func StartPingPongClient() {
	// Disable transport security for this example only,
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial("127.0.0.1:9000", opts...)
	if err != nil {
		panic(err)
	}

	client := pb.NewHelloWorldClient(conn)

	resp, err := CallSayHi(client)
	if err != nil {
		panic(err)
	}
	fmt.Println("Finish hello world with response = ", resp)
}

func main() {
	StartPingPongClient()
}
