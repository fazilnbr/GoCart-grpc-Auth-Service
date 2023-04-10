package usecase

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"

	"github.com/fazilnbr/banking-grpc-auth-service/pkg/domain"
	repository "github.com/fazilnbr/banking-grpc-auth-service/pkg/repository/interface"
	interfaces "github.com/fazilnbr/banking-grpc-auth-service/pkg/usecase/interface"
)

type userUseCase struct {
	userRepo repository.UserRepository
}

// FindByName implements interfaces.UserUseCase
func (c *userUseCase) FindByName(ctx context.Context, userId int) (domain.User, error) {
	user, err := c.userRepo.FindByName(ctx, userId)
	return user, err
}

// Login implements interfaces.UserUseCase
func (c *userUseCase) Login(user domain.User) (domain.User, error) {

	if user.UserName != "" {
		userData, err := c.userRepo.FindUserWithUserName(user)
		if err != nil {
			return userData, errors.New("invalid user name or password")
		}
		if !VerifyPassword(user.Password, userData.Password) {
			return userData, errors.New("invalid user name or password")
		}
		return userData, nil
	} else if user.Email != "" {
		userData, err := c.userRepo.FindUserWithEmail(user)
		fmt.Println("userdata", userData)
		if err != nil {
			return userData, errors.New("invalid email or password")
		}
		if !VerifyPassword(user.Password, userData.Password) {
			return userData, errors.New("invalid email or password")
		}
		return userData, nil
	}
	return user, nil

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
