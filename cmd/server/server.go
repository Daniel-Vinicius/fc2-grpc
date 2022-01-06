package main

import (
	"log"
	"net"

	"github.com/Daniel-Vinicius/fc2-grpc/pb"
	"github.com/Daniel-Vinicius/fc2-grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())

	// Without this line evans grpc client does not work
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}
