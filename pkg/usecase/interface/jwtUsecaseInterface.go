package interfaces

import (
	"github.com/fazilnbr/banking-grpc-auth-service/pkg/domain"
	"github.com/golang-jwt/jwt"
)

type JWTUsecase interface {
	GenerateRefreshToken(userid int, username string, role string) (string, error)
	GenerateAccessToken(userid int, username string, role string) (string, error)
	VerifyToken(signedToken string) (bool, *domain.SignedDetails)
	GetTokenFromString(signedToken string, claims *domain.SignedDetails) (*jwt.Token, error)
}
