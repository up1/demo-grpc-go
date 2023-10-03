package hello

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type HelloServer struct {
}

// mustEmbedUnimplementedHelloWorldServer implements hello.HelloWorldServer.
func (s *HelloServer) mustEmbedUnimplementedHelloWorldServer() {
	panic("unimplemented")
}

func (s *HelloServer) SayHi(ctx context.Context, hello *Hello) (*HelloResponse, error) {
	fmt.Println("Called sayHi")

	resp := HelloResponse{
		Id:      hello.Id,
		Message: hello.Message,
	}

	return &resp, nil
}

func StartServer() {
	server := HelloServer{}
	lis, err := net.Listen("tcp", "localhost:9000")
	grpcServer := grpc.NewServer()
	RegisterHelloWorldServer(grpcServer, &server)

	// Start grpcServer
	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
