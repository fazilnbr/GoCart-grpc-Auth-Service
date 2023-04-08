package http

import (
	"fmt"
	"log"
	"net"

	"github.com/fazilnbr/banking-grpc-auth-service/pkg/api/handler"
	"google.golang.org/grpc"
)

func NewGRPCServer(userHandler *handler.UserHandler, grpcPort string) {
	lis, err := net.Listen("tcp", ":"+grpcPort)
	fmt.Println("grpcPort/////", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// pb.RegisterAuthServiceServer(grpcServer, userHandler)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
