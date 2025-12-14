package service

// import (
// 	"github.com/omarsabri666/url_shorter/helpers"
// 	"github.com/omarsabri666/url_shorter/repository"
// )

// type TokenService struct {
// 	repo repository.TokenRepository
// }

// func NewTokenService(repo repository.TokenRepository) *TokenService {
// 	return &TokenService{repo: repo}
// }


// func (t *TokenService) InsertToken(token repository.TokenStruct) (*repository.TokenStruct ,error) {
// 	var to repository.TokenStruct




// 	err:= t.repo.InsertToken(token)


// 	if err != nil {
// 		return nil,err
// 	}

// 	return &to,nil


// }


// func (t *TokenService) GetToken(token repository.TokenStruct) (*repository.TokenStruct, error) {
// 	// return t.repo.GetToken(token)
// 	 refreshToken,err:= t.repo.GetToken(token)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return refreshToken, nil
// }
// func (t *TokenService) RefreshToken(token string) (*repository.TokenStruct, error) {
// 	userId , err := helpers.VerifyRefreshToken(token)

// 	if err != nil {
// 		return nil, err
// 	}
// 	refreshTokenStruct := repository.TokenStruct{RefreshToken: token, UserId: userId}
//  _,err = 	t.repo.GetToken(refreshTokenStruct)

// 	if err != nil {
// 		return nil, err
// 	}
//  accessToken, err :=	helpers.GenerateToken(userId,helpers.AccessToken)
//  if err != nil {
// 	 return nil, err
//  }
//  refreshTokenStruct.AccessToken = accessToken




// 	return  &refreshTokenStruct , nil
// }
// 	// return t.repo.GetToken(token)
