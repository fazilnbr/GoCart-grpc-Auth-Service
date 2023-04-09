package repository

import (
	"errors"

	"github.com/fazilnbr/banking-grpc-auth-service/pkg/domain"
	interfaces "github.com/fazilnbr/banking-grpc-auth-service/pkg/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

// FindUser implements interfaces.UserRepository
func (c *userDatabase) FindUser(user domain.User) (int, error) {
	err := c.DB.Where("user_name = ?", user.UserName).First(&user).Error
	if err == nil {
		return 0, errors.New("username alredy exist")
	}
	// if err == gorm.ErrRecordNotFound {
	// 	return 0, nil
	// }
	// if err != nil {
	// 	return 0, err
	// }
	err = c.DB.Where("email = ?", user.Email).First(&user).Error
	if err == nil {
		return 0, errors.New("email alredy exist")
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
