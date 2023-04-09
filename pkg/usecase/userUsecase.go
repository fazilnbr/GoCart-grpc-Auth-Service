package usecase

import (
	"github.com/fazilnbr/banking-grpc-auth-service/pkg/domain"
	repository "github.com/fazilnbr/banking-grpc-auth-service/pkg/repository/interface"
	interfaces "github.com/fazilnbr/banking-grpc-auth-service/pkg/usecase/interface"
)

type userUseCase struct {
	userRepo repository.UserRepository
}

// Register implements interfaces.UserUseCase
func (c *userUseCase) Register(user domain.User) (int, error) {
	userId, err := c.userRepo.Register(user)
	return userId, err
}

func NewUserUseCase(repo repository.UserRepository) interfaces.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}
