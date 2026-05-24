package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTManager interface {
	Generate(UserID int, Name string) (string, error)
	Validate(tokenString string) (*JWTClaims, error)
}

type jwtManager struct {
	secretKey string
	issuer    string
	duration  time.Duration
}

type JWTClaims struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

func NewJWTManager(secretKey string, issuer string, duration time.Duration) JWTManager {
	return &jwtManager{
		secretKey: secretKey,
		issuer:    issuer,
		duration:  duration,
	}
}

func (j *jwtManager) Generate(UserID int, Name string) (string, error) {
	now := time.Now()

	claims := &JWTClaims{
		UserID: UserID,
		Name:   Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(j.duration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtManager) Validate(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(j.secretKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
