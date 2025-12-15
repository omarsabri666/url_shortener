package helpers

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/omarsabri666/url_shorter/global"
)

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

func GenerateToken(id string, tokenType TokenType) (string, error) {
	now := time.Now()
	var (
		expiration time.Duration
		secret     string
	)
	switch tokenType {
	case AccessToken:
		expiration = global.AccessTokenExp //  1 hour access token go crazy users
		secret = os.Getenv("ACCESS_TOKEN")
	case RefreshToken:
		expiration = global.RefreshTokenExp // 30 days for refresh tokens
		secret = os.Getenv("REFRESH_TOKEN")
	default:
		return "", jwt.ErrInvalidType
	}

	// Validate secret
	if secret == "" {
		return "", jwt.ErrInvalidKey
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":  id,
			"iat":  now.Unix(),
			"exp":  now.Add(expiration).Unix(),
			"type": string(tokenType), // Helps identify token type during validation
		},
	)
	return token.SignedString([]byte(secret))

}
func VerifyRefreshToken(refreshToken string) (string, error) {
	tokenKey := os.Getenv("REFRESH_TOKEN")
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil || !token.Valid {

		return "", err
	}
	userID, ok := claims["sub"].(string)
	if !ok || userID == "" {

		return "", errors.New("invalid token claims")
	}

	return userID, nil

}
