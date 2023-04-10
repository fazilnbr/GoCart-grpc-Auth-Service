package interfaces

import (
	"context"

	"github.com/fazilnbr/banking-grpc-auth-service/pkg/domain"
)

type UserRepository interface {
	FindUser(user domain.User) (int, error)
	CreateUser(user domain.User) (int, error)
	FindUserWithUserName(user domain.User) (domain.User, error)
	FindUserWithEmail(user domain.User) (domain.User, error)
	FindByName(ctx context.Context, userId int) (domain.User, error)
}
