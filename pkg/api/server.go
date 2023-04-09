package http

import (
	"fmt"
	"log"
	"net"

	"github.com/fazilnbr/banking-grpc-auth-service/pkg/api/handler"
	"github.com/fazilnbr/banking-grpc-auth-service/pkg/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewGRPCServer(userHandler *handler.UserHandler, grpcPort string) {
	lis, err := net.Listen("tcp", ":"+grpcPort)
	fmt.Println("grpcPort/////", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, userHandler)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func NewServerHTTP(userHandler *handler.UserHandler) *ServerHTTP {
	engine := gin.New()
	go NewGRPCServer(userHandler, "50081")
	// Use logger from Gin
	engine.Use(gin.Logger())

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":8000")
}
