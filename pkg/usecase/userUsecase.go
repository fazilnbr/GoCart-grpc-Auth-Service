package usecase

import (
	"crypto/md5"
	"fmt"

	"github.com/fazilnbr/banking-grpc-auth-service/pkg/domain"
	repository "github.com/fazilnbr/banking-grpc-auth-service/pkg/repository/interface"
	interfaces "github.com/fazilnbr/banking-grpc-auth-service/pkg/usecase/interface"
)

type userUseCase struct {
	userRepo repository.UserRepository
}

// Register implements interfaces.UserUseCase
func (c *userUseCase) Register(user domain.User) (int, error) {
	userId, err := c.userRepo.FindUser(user)
	if err != nil {
		return userId, err
	}
	user.Password = HashPassword(user.Password)
	userId, err = c.userRepo.CreateUser(user)
	return userId, err
}

func HashPassword(password string) string {
	data := []byte(password)
	password = fmt.Sprintf("%x", md5.Sum(data))
	return password

}

func VerifyPassword(requestPassword, dbPassword string) bool {

	requestPassword = fmt.Sprintf("%x", md5.Sum([]byte(requestPassword)))
	return requestPassword == dbPassword
}

func NewUserUseCase(repo repository.UserRepository) interfaces.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}
