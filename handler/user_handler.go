package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	errs "github.com/omarsabri666/url_shorter/err"
	"github.com/omarsabri666/url_shorter/model/user"
	"github.com/omarsabri666/url_shorter/repository"
	service "github.com/omarsabri666/url_shorter/service/user"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService ) *UserHandler {
	return &UserHandler{service: service }
}
func (u *UserHandler) Signup(c *gin.Context) {
	var req user.User
	if c.Request.ContentLength == 0 {

		HandleError(c,errs.BadRequest("empty body"))
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {


		HandleError(c,errs.BadRequest(err.Error()))
		// c.JSON(400, gin.H{"error": err.Error()})
		return

	}
err:=	u.service.Signup(req)

	if err != nil {
		log.Println(err)
		HandleError(c,err)
		// c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "Signup successful"})

}
func (u *UserHandler) Login(c *gin.Context) {
	var req user.UserSignin
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		HandleError(c,errs.BadRequest(err.Error()))
		// c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	UserToken, err := u.service.Login(req,c.Request.Context())
	if err != nil {
		log.Println(err)
		HandleError(c,err)
		// c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("refresh_token", UserToken.RefreshToken, 3600 * 24 * 30, "/", "localhost", false, true)
	c.JSON(200, gin.H{"message": "Login successful", "access_token": UserToken.AccessToken,})

}
func (u *UserHandler) Logout(c *gin.Context) {
	var tStruct repository.TokenStruct

token,err:=	c.Cookie("refresh_token")
	if err != nil {
		log.Println(err)
		HandleError(c,errs.Unauthorized(err.Error()))
		return
	}
	 userId,_:= c.Get("userID")

	tStruct.RefreshToken=token
	tStruct.UserId=userId.(string)
err =	u.service.Logout(tStruct,c.Request.Context())

	if err != nil {
		log.Println(err)
		HandleError(c,err)
		// c.JSON(400, gin.H{"error": err.Error()})
		return
	}	
	    c.SetCookie("refresh_token", "", -1, "/", "localhost", true, true)

	c.JSON(200, gin.H{"message": "Logout successful"})





}
func (h *UserHandler) RefreshToken(c *gin.Context) {
	refreshToken , err:= c.Cookie("refresh_token")
	if err != nil {
		log.Println(err)
		HandleError(c,errs.Unauthorized(err.Error()))
		// c.JSON(400, gin.H{"error": "could not get refresh token"})
		return
	}
//  token , err :=	h.token.RefreshToken(refreshToken)
  token , err :=	h.service.RefreshToken(refreshToken,c.Request.Context())

	if err != nil {
		log.Println(err)
		HandleError(c,err)
		return
	}

 

 c.JSON(201,gin.H{"token": token.AccessToken})



	
}