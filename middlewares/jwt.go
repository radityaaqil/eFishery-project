package middlewares

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type JWTService interface {
	GenerateToken(userId int) (string, error)
	ValidateToken(TokenGenerated string) (*jwt.Token, error)
}

type jwt_service struct{}

func NewJwtService() *jwt_service {
	return &jwt_service{}
}

func LoadSecret() []byte {
	err := godotenv.Load(".env")

	if err != nil {
		log.Panic("Failed to load env")
	}

	SECRET := os.Getenv("JWT_SECRET")

	return []byte(SECRET)
}

func (s *jwt_service) GenerateToken(userId int) (string, error) {
	JWT_SECRET := LoadSecret()

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(JWT_SECRET)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (s *jwt_service) ValidateToken(TokenGenerated string) (*jwt.Token, error) {
	JWT_SECRET := LoadSecret()

	token1, err := jwt.Parse(TokenGenerated, func(token1 *jwt.Token) (interface{}, error) {
		_, ok := token1.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		return token1, err
	}
	return token1, nil
}
