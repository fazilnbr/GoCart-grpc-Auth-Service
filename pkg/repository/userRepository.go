package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/fazilnbr/banking-grpc-auth-service/pkg/domain"
	interfaces "github.com/fazilnbr/banking-grpc-auth-service/pkg/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

// FindByName implements interfaces.UserRepository
func (c *userDatabase) FindByName(ctx context.Context, userId int) (domain.User, error) {
	var user domain.User
	err := c.DB.Where("id_user = ?", userId).First(&user).Error
	fmt.Println("err from repos find name", err)
	return user, err
}

// FindUserWithEmail implements interfaces.UserRepository
func (c *userDatabase) FindUserWithEmail(user domain.User) (domain.User, error) {
	err := c.DB.Where("email = ?", user.Email).First(&user).Error
	return user, err
}

// FindUserWithUserName implements interfaces.UserRepository
func (c *userDatabase) FindUserWithUserName(user domain.User) (domain.User, error) {
	err := c.DB.Where("user_name = ?", user.UserName).First(&user).Error
	return user, err
}

// FindUser implements interfaces.UserRepository
func (c *userDatabase) FindUser(user domain.User) (int, error) {
	err := c.DB.Where("user_name = ?", user.UserName).First(&user).Error
	if err == nil {
		return 0, errors.New("Username alredy exist")
	}
	err = c.DB.Where("email = ?", user.Email).First(&user).Error
	if err == nil {
		return 0, errors.New("Email alredy exist")
	}
	if err == gorm.ErrRecordNotFound {
		return 0, nil
	}

	return user.IdUser, err
}

// Register implements interfaces.UserRepository
func (c *userDatabase) CreateUser(user domain.User) (int, error) {
	result := c.DB.Create(&user)
	return user.IdUser, result.Error
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}
