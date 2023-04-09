package interfaces

import "github.com/fazilnbr/banking-grpc-auth-service/pkg/domain"

type UserUseCase interface {
	Register(user domain.User) (int, error)
}
