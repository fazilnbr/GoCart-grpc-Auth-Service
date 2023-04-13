package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/fazilnbr/banking-grpc-auth-service/pkg/domain"
	"github.com/fazilnbr/banking-grpc-auth-service/pkg/pb"
	usecase "github.com/fazilnbr/banking-grpc-auth-service/pkg/usecase/interface"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
	jwtUsecase  usecase.JWTUsecase
}

func (cr *UserHandler) TokenRefresh(ctx context.Context, req *pb.TokenRefreshRequest) (*pb.TokenRefreshResponse, error) {

	fmt.Println("the tocker : ", req.Token)
	ok, claims := cr.jwtUsecase.VerifyToken(req.Token)
	if !ok {
		return &pb.TokenRefreshResponse{
			Status: http.StatusUnauthorized,
			Error:  fmt.Sprint(errors.New("token verification failed")),
		}, nil
	}

	fmt.Println("//////////////////////////////////", claims.UserName)
	accesstoken, err := cr.jwtUsecase.GenerateAccessToken(int(claims.UserId), claims.UserName, "user")

	if err != nil {
		return &pb.TokenRefreshResponse{
			Status: http.StatusUnauthorized,
			Error:  fmt.Sprint(errors.New("unable to generate access token")),
		}, errors.New(err.Error())
	}
	return &pb.TokenRefreshResponse{
		Status: http.StatusOK,
		Token:  accesstoken,
	}, nil

}

// Validate implements pb.AuthServiceServer
func (cr *UserHandler) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	ok, claims := cr.jwtUsecase.VerifyToken(req.Token)
	fmt.Println("claims", claims.UserId)
	if !ok {
		return &pb.ValidateResponse{
			Status: http.StatusUnauthorized,
			Error:  fmt.Sprint(errors.New("token verification failed")),
		}, nil
	}

	user, err := cr.userUseCase.FindByName(ctx, int(claims.UserId))

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusUnauthorized,
			Error:  fmt.Sprint(errors.New("user not found with token credentials")),
		}, errors.New(err.Error())
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: int64(user.IdUser),
		Source: fmt.Sprint(claims.Source),
	}, nil
}

func (cr *UserHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := domain.User{
		UserName: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
	_, err := cr.userUseCase.Register(user)
	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusUnprocessableEntity,
			// Id:     int64(userId),
			Error: "err",
		}, err
	}

	return &pb.RegisterResponse{
		Status: http.StatusOK,
		// Id:     int64(userId),
	}, nil

}

// Login implements pb.AuthServiceServer
func (cr *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user := domain.User{
		Email:    req.Email,
		UserName: req.Username,
		Password: req.Password,
	}
	userData, err := cr.userUseCase.Login(user)

	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  err.Error(),
		}, err
	}

	accesToken, err := cr.jwtUsecase.GenerateAccessToken(userData.IdUser, req.Username, "user")
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "error while generating token",
		}, err

	}
	refreshTocken, err := cr.jwtUsecase.GenerateRefreshToken(userData.IdUser, req.Username, "user")
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "error while generating token",
		}, err

	}

	return &pb.LoginResponse{
		Status:       http.StatusOK,
		Accesstoken:  accesToken,
		Refreshtoken: refreshTocken,
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
