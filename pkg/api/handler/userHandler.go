package handler

import (
	usecase "github.com/fazilnbr/banking-grpc-auth-service/pkg/usecase/interface"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
	jwtUsecase  usecase.JWTUsecase
}


func NewUserHandler(usecase usecase.UserUseCase, jwtusecase usecase.JWTUsecase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
		jwtUsecase:  jwtusecase,
	}
}
