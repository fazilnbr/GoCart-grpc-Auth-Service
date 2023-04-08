package usecase

import (
	 "os"
	 interfaces "github.com/fazilnbr/banking-grpc-auth-service/pkg/usecase/interface"
)

type jwtUsecase struct {
	SecretKey string
}

func NewJWTUsecase() interfaces.JWTUsecase {
	return &jwtUsecase{
		SecretKey: os.Getenv("SECRET_KEY"),
	}
}
