package interfaces

import "github.com/fazilnbr/banking-grpc-auth-service/pkg/domain"

type UserUseCase interface {
	Login(user domain.User) (error)
	Register(user domain.User) (int, error)
}
