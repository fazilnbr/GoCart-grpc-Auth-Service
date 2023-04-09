package repository

import (
	"github.com/fazilnbr/banking-grpc-auth-service/pkg/domain"
	interfaces "github.com/fazilnbr/banking-grpc-auth-service/pkg/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

// Register implements interfaces.UserRepository
func (c *userDatabase) Register(user domain.User) (int, error) {
	result:=c.DB.Create(&user)
	return user.IdUser,result.Error
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}
