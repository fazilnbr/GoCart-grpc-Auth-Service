package handler

import (
	"context"
	"net/http"

	"github.com/fazilnbr/banking-grpc-auth-service/pkg/domain"
	"github.com/fazilnbr/banking-grpc-auth-service/pkg/pb"
	usecase "github.com/fazilnbr/banking-grpc-auth-service/pkg/usecase/interface"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
	jwtUsecase  usecase.JWTUsecase
}

func (cr *UserHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := domain.User{
		UserName: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
	userId, err := cr.userUseCase.Register(user)
	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusUnprocessableEntity,
			Id:     int64(userId),
			Error:  "err",
		}, err
	}

	return &pb.RegisterResponse{
		Status: http.StatusOK,
		Id:     int64(userId),
	}, nil

}

// Login implements pb.AuthServiceServer
func (cr *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user := domain.User{
		Email:    req.Email,
		UserName: req.Username,
		Password: req.Password,
	}
	err := cr.userUseCase.Login(user)
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusUnprocessableEntity,
		}, err
	}

	return &pb.LoginResponse{
		Status: http.StatusOK,
	}, nil
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
