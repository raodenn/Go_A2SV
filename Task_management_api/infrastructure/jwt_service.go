package infrastructure

import (
	domain "task_manager/domain"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtSvc struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewJWTService(secretKey string, duration time.Duration) domain.JwtSvc {
	return &JwtSvc{
		secretKey:     secretKey,
		tokenDuration: duration,
	}
}

func (s *JwtSvc) GenerateToken(userID string, userType string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":   userID,
		"user_type": userType,
		"exp":       time.Now().Add(s.tokenDuration).Unix(),
		"iat":       time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *JwtSvc) ValidateToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.secretKey), nil
	})
}
