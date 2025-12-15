package repository

import (
	"context"
	"fmt"
	"log"

	errs "github.com/omarsabri666/url_shorter/err"
	"github.com/omarsabri666/url_shorter/global"
	"github.com/redis/go-redis/v9"
)

type TokenRepositoryRedis struct {
	redis *redis.Client
}

func NewTokenRepositoryRedis(redis *redis.Client) *TokenRepositoryRedis {
	return &TokenRepositoryRedis{redis: redis}
}

func (r *TokenRepositoryRedis) InsertToken(token TokenStruct, c context.Context) error {
	key := fmt.Sprintf("refresh_token:%s:%s", token.UserId, token.RefreshToken)
	err := r.redis.Set(c, key, token.RefreshToken, global.RefreshTokenExp).Err()

	if err != nil {
		log.Println(err)

		return errs.InternalServerError("failed to insert token")
	}

	return nil
}
func (r *TokenRepositoryRedis) GetToken(token TokenStruct, c context.Context) (*TokenStruct, error) {
	key := fmt.Sprintf("refresh_token:%s:%s", token.UserId, token.RefreshToken)
	refreshToken, err := r.redis.Get(c, key).Result()
	if err == redis.Nil {
		return nil, errs.Unauthorized("token expired")
	}

	if err != nil {
		log.Println(err)
		return nil, errs.InternalServerError("failed to get token")
	}

	return &TokenStruct{RefreshToken: refreshToken}, nil

}
func (r *TokenRepositoryRedis) DeleteToken(token TokenStruct, c context.Context) error {
	key := fmt.Sprintf("refresh_token:%s:%s", token.UserId, token.RefreshToken)
	err := r.redis.Del(c, key).Err()

	if err != nil {
		log.Println(err)
		return errs.InternalServerError("failed to delete token")
	}
	return nil
}
