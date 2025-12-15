package repository

import (
	"context"
	"time"
)

type TokenStruct struct {
	AccessToken  string
	RefreshToken string
	UserId       string
	Exp          time.Time
}
type TokenRepository interface {
	InsertToken(token TokenStruct, context context.Context) error
	GetToken(token TokenStruct, context context.Context) (*TokenStruct, error)
	DeleteToken(token TokenStruct, context context.Context) error
	// RefreshToken(token string) (*TokenStruct, error)
}
