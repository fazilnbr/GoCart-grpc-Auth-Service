package interfaces

import (
	"context"

	"github.com/fazilnbr/banking-grpc-auth-service/pkg/domain"
)

type UserUseCase interface {
	Login(user domain.User) (domain.User, error)
	Register(user domain.User) (int, error)
	FindByName(ctx context.Context, userId int) (domain.User, error)
}
