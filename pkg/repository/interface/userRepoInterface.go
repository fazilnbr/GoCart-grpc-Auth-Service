package interfaces

import (
	"github.com/fazilnbr/banking-grpc-auth-service/pkg/domain"
)

type UserRepository interface {
	Register(user domain.User) (int, error)
}
