package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/omarsabri666/url_shorter/global"
	"github.com/omarsabri666/url_shorter/helpers"
	"github.com/omarsabri666/url_shorter/model/token"
	"github.com/omarsabri666/url_shorter/model/user"
	"github.com/omarsabri666/url_shorter/repository"
	service "github.com/omarsabri666/url_shorter/service/user"
	"github.com/omarsabri666/url_shorter/validators"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (u *UserHandler) SignupHttp(w http.ResponseWriter, r *http.Request) {
	var req user.User
	defer r.Body.Close()
	if r.ContentLength == 0 {
		helpers.WriteJson(w, 400, user.UserResponse{Message: "empty body", Success: false})
	}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		helpers.WriteJson(w, 400, user.UserResponse{Message: "invalid json body", Success: false})
		return
	}
	if err := validators.Validate.Struct(req); err != nil {
		helpers.WriteJson(w, 400, user.UserResponse{Message: err.Error(), Success: false})
		return
	}
	err := u.service.Signup(req)
	if err != nil {
		log.Println(err)
		helpers.WriteJson(w, 400, user.UserResponse{Message: err.Error(), Success: false})
		return
	}
	helpers.WriteJson(w, 201, user.UserResponse{Message: "Signup successful", Success: true})

}
func (u *UserHandler) LoginHttp(w http.ResponseWriter, r *http.Request) {
	var req user.UserSignin
	defer r.Body.Close()
	if r.ContentLength == 0 {
		helpers.WriteJson(w, 400, user.UserResponse{Message: "empty body", Success: false})

		// http.Error(w, "empty body", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		helpers.WriteJson(w, 400, user.UserResponse{Message: "invalid json body", Success: false})
		// http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}
	if err := validators.Validate.Struct(req); err != nil {
		helpers.WriteJson(w, 400, user.UserResponse{Message: err.Error(), Success: false})
		// http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	UserToken, err := u.service.Login(req, r.Context())
	if err != nil {
		log.Println(err)
		helpers.WriteJson(w, 401, user.UserResponse{Message: "invalid credentials", Success: false})
		// http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	domain := os.Getenv("DOMAIN")

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    UserToken.RefreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // false if HTTP local dev
		SameSite: http.SameSiteStrictMode,
		MaxAge:   60 * 60 * 24 * 30, // 30 days
		Domain:   domain,
	})

	helpers.WriteJson(w, 200, user.UserResponse{Message: "Login successful", Success: true, Data: map[string]string{"access_token": UserToken.AccessToken}})
}

func (u UserHandler) LogoutHttp(w http.ResponseWriter, r *http.Request) {
	var tStruct repository.TokenStruct
	refreshToken, err := r.Cookie("refresh_token")
	if err != nil {
		log.Println(err)
		helpers.WriteJson(w, 401, token.TokenResponse{Message: "Token missing", Success: false})
		return
	}
	tStruct.RefreshToken = refreshToken.Value
	userId, _ := r.Context().Value(global.UserIdKey).(string)
	tStruct.UserId = userId
	err = u.service.Logout(tStruct, r.Context())
	if err != nil {
		log.Println(err)
		helpers.WriteJson(w, 500, token.TokenResponse{Message: err.Error(), Success: false})
		return
	}
	domain := os.Getenv("DOMAIN")

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // false if HTTP local dev
		SameSite: http.SameSiteStrictMode,
		MaxAge:   -1,
		Domain:   domain,
	})

	helpers.WriteJson(w, 200, token.TokenResponse{Message: "Logout successful", Success: true})

}

func (u *UserHandler) RefreshTokenHttp(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := r.Cookie("refresh_token")
	if err != nil {
		helpers.WriteJson(w, 401, token.TokenResponse{Message: "Token missing", Success: false})
		return

	}
	accessTokenStruct, err := u.service.RefreshToken(refreshToken.Value, r.Context())

	if err != nil {
		log.Println(err)
		helpers.WriteJson(w, 401, token.TokenResponse{Message: err.Error(), Success: false})
		return
	}

	helpers.WriteJson(w, 200, token.TokenResponse{Data: map[string]string{"access_token": accessTokenStruct.AccessToken}, Success: true})

}
