package usecase

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fazilnbr/banking-grpc-auth-service/pkg/domain"
	interfaces "github.com/fazilnbr/banking-grpc-auth-service/pkg/usecase/interface"
	"github.com/golang-jwt/jwt"
)

type jwtUsecase struct {
	SecretKey string
}

// GenerateRefreshToken implements interfaces.jwtUsecase
func (j *jwtUsecase) GenerateRefreshToken(userid int, username string, role string) (string, error) {
	claims := &domain.SignedDetails{
		UserId:   uint(userid),
		UserName: username,
		Source:   "refreshtoken",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 12 * 7).Unix(),
		},
	}
	// fmt.Printf("\n\nrefresh time : %v\n\n", time.Hour*12*7)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(j.SecretKey))

	if err != nil {
		log.Println(err)
	}

	return signedToken, err
}

// GenerateRefreshToken implements interfaces.jwtUsecase
func (j *jwtUsecase) GenerateAccessToken(userid int, username string, role string) (string, error) {

	claims := &domain.SignedDetails{
		UserId:   uint(userid),
		UserName: username,
		Source:   "accesstoken",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(500)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(j.SecretKey))

	if err != nil {
		log.Println(err)
	}

	return signedToken, err
}

// // GetTokenFromString implements interfaces.jwtUsecase
func (j *jwtUsecase) GetTokenFromString(signedToken string, claims *domain.SignedDetails) (*jwt.Token, error) {
	return jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.SecretKey), nil
	})

}

// VerifyToken implements interfaces.jwtUsecase
func (j *jwtUsecase) VerifyToken(signedToken string) (bool, *domain.SignedDetails) {
	claims := &domain.SignedDetails{}
	token, err := j.GetTokenFromString(signedToken, claims)

	if err != nil {
		return false, claims
	}

	if token.Valid {
		if e := claims.Valid(); e == nil {
			return true, claims
		}
	}
	return false, claims
}

func NewJWTUsecase() interfaces.JWTUsecase {
	return &jwtUsecase{
		SecretKey: os.Getenv("SECRET_KEY"),
	}
}
