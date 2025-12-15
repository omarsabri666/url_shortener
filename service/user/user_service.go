package service

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	errs "github.com/omarsabri666/url_shorter/err"
	"github.com/omarsabri666/url_shorter/helpers"
	"github.com/omarsabri666/url_shorter/model/user"
	"github.com/omarsabri666/url_shorter/repository"
)

type UserService struct {
	userRepo  repository.UserRepository
	tokenRepo repository.TokenRepository
}

func NewUserService(userRepo repository.UserRepository, tokenRepo repository.TokenRepository) *UserService {
	return &UserService{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
	}
}

func (u *UserService) Signup(user user.User) error {
	hashedPassword, hashPasswordErr := helpers.HashPassword(user.Password, 10)
	if hashPasswordErr != nil {

		return errs.InternalServerError(hashPasswordErr.Error())
	}
	user.Password = hashedPassword
	id := uuid.New().String()
	user.Id = &id

	err := u.userRepo.CreateUser(user)

	if err != nil {
		// errs.Conflict()
		if strings.HasPrefix(err.Error(), "duplicate entry for") {
			return errs.Conflict(err.Error())
		}

		return err
	}

	return nil
}
func (u *UserService) Login(req user.UserSignin, context context.Context) (*repository.TokenStruct, error) {
	var UserToken repository.TokenStruct

	user, err := u.userRepo.GetUser(req)
	if err != nil {
		log.Printf("err getting user : %v", err)

		return nil, err
	}
	if user == nil || user.Id == nil {
		log.Println("No user found with this email")

		return nil, errs.Unauthorized("invalid email or password")
	}
	UserToken.UserId = *user.Id
	comparePassErr := helpers.ComparePassword(req.Password, user.Password)
	if comparePassErr != nil {
		log.Printf("Could not compare password %v", comparePassErr)
		return nil, errs.Unauthorized("invalid email or password")
	}
	accessToken, accessTokenErr := helpers.GenerateToken(*user.Id, helpers.AccessToken)
	if accessTokenErr != nil {
		log.Printf("Could not generate access token %v", accessTokenErr)

		return nil, errs.InternalServerError("could not generate access token")
	}
	refreshToken, refreshTokenErr := helpers.GenerateToken(*user.Id, helpers.RefreshToken)
	if refreshTokenErr != nil {
		log.Printf("Could not generate refresh token %v", refreshTokenErr)
		return nil, errs.InternalServerError("could not generate refresh token")
	}
	UserToken.AccessToken = accessToken
	UserToken.RefreshToken = refreshToken
	expiry := time.Now().Add(30 * 24 * time.Hour)

	UserToken.Exp = expiry
	insertTokenErr := u.tokenRepo.InsertToken(UserToken, context)
	if insertTokenErr != nil {
		log.Printf("failed to insert token: %v", insertTokenErr)

		return nil, errs.InternalServerError("failed to insert token")
	}
	return &UserToken, nil

}
func (u *UserService) Logout(token repository.TokenStruct, context context.Context) error {
	err := u.tokenRepo.DeleteToken(token, context)

	if err != nil {
		return err
	}
	return nil
}
func (t *UserService) RefreshToken(token string, context context.Context) (*repository.TokenStruct, error) {
	userId, err := helpers.VerifyRefreshToken(token)

	if err != nil {
		return nil, err
	}
	refreshTokenStruct := repository.TokenStruct{RefreshToken: token, UserId: userId}
	_, err = t.tokenRepo.GetToken(refreshTokenStruct, context)

	if err != nil {
		return nil, err
	}
	accessToken, err := helpers.GenerateToken(userId, helpers.AccessToken)
	if err != nil {
		return nil, err
	}
	refreshTokenStruct.AccessToken = accessToken

	return &refreshTokenStruct, nil
}
