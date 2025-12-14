package handler

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	errs "github.com/omarsabri666/url_shorter/err"
	"github.com/omarsabri666/url_shorter/model/url"
	service "github.com/omarsabri666/url_shorter/service/url"
)

type URLHandler struct {
	service *service.URLService
}

func NewURLHandler(service *service.URLService) *URLHandler {
	return &URLHandler{service: service}
}


func (h *URLHandler) CreateURL(c *gin.Context) {
	var req url.CreateURLRequest
	var res url.CreateUrlResponse
	
	ctx,cancel := context.WithTimeout(c.Request.Context(),time.Second*2)
	defer cancel()
	if c.Request.ContentLength == 0 {
		HandleError(c,errs.BadRequest("empty body"))
		return
	}



	if err := c.ShouldBindJSON(&req); err != nil {
		HandleError(c,errs.BadRequest(err.Error()))
	
		return
	}
	shortUrl, err := h.service.CreateURL(req,ctx)
	if err != nil {
		log.Println(err)
		HandleError(c,err)
			

		return
	}
	res.Success = true
	res.ShortUrl = shortUrl
	res.Message = "URL created successfully"
	c.JSON(201, res)

	// c.JSON(200, gin.H{"short_url": shortUrl})

	

} 
func (h *URLHandler) GetURL(c *gin.Context) {
// var req url.GetUrlRequest
var res url.GetUrlResponse
shortUrl := c.Param("short_url")
if shortUrl == "" {
	HandleError(c,errs.BadRequest("short url is required"))
	return
}




	u, err := h.service.GetURL(shortUrl,c.Request.Context())
	if err != nil {
		log.Println(err)
		HandleError(c,err)
		return
	}
	res.LongUrl = u.LongUrl
	c.Redirect(302, res.LongUrl)




}
