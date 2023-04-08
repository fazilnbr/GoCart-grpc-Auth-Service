package usecase

import (
	repository "github.com/fazilnbr/banking-grpc-auth-service/pkg/repository/interface"
	interfaces "github.com/fazilnbr/banking-grpc-auth-service/pkg/usecase/interface"
)

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) interfaces.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}
