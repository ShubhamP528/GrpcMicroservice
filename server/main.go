package main

import (
	"context"
	pb "grpc/proto" // Import the generated code
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedMyServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func (s *server) GetHelloPage(ctx context.Context, _ *pb.HelloRequest) (*pb.HelloResponse, error) {
	htmlPage := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Hello Page</title>
    </head>
    <body>
        <h1>Hello, World!</h1>
    </body>
    </html>
    `
	return &pb.HelloResponse{Message: htmlPage}, nil
}

// func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
//     return &pb.HelloResponse{Message: "Hello, " + in.Name}, nil
// }
