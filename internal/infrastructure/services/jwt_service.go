package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secretKey []byte
}

func NewJwtService(secret string) *JWTService {
	if secret == "" {

	}

	return &JWTService{
		secretKey: []byte(secret),
	}
}

func (s *JWTService) GenerateToken(userID uint, email string) (string, error) {
	expiration := time.Now().Add(72 * time.Hour)

	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     expiration.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(s.secretKey)
}

func (s *JWTService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return s.secretKey, nil
	})
}
