package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/fazilnbr/banking-grpc-auth-service/pkg/pb"
	usecase "github.com/fazilnbr/banking-grpc-auth-service/pkg/usecase/interface"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
	jwtUsecase  usecase.JWTUsecase
}

func (cr *UserHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	return &pb.RegisterResponse{
		Status: http.StatusUnprocessableEntity,
		Id:     1,
		Error:  fmt.Sprint(errors.New("email already exist")),
	}, nil

}

// Login implements pb.AuthServiceServer
func (*UserHandler) Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error) {
	panic("unimplemented")
}

// Refresh implements pb.AuthServiceServer
func (*UserHandler) Refresh(context.Context, *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	panic("unimplemented")
}

// mustEmbedUnimplementedAuthServiceServer implements pb.AuthServiceServer
func (*UserHandler) mustEmbedUnimplementedAuthServiceServer() {
	panic("unimplemented")
}

func NewUserHandler(usecase usecase.UserUseCase, jwtusecase usecase.JWTUsecase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
		jwtUsecase:  jwtusecase,
	}
}
