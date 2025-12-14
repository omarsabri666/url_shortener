package service

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	errs "github.com/omarsabri666/url_shorter/err"
	"github.com/omarsabri666/url_shorter/helpers"
	"github.com/omarsabri666/url_shorter/model/url"
	"github.com/omarsabri666/url_shorter/repository"
	"github.com/redis/go-redis/v9"
)

type URLService struct {
	repo  repository.URLRepository
	redis *redis.Client
}
func NewUrlService(repo repository.URLRepository , redisCleint *redis.Client) *URLService {
	return &URLService{repo: repo, redis: redisCleint}
}



func (s *URLService) CreateURL(req url.CreateURLRequest, c context.Context) (string ,error) {
	log.Println(req.Alias)
	  if err := godotenv.Load(); err != nil {
        log.Println(".env file not found, relying on system environment variables")
    }
	    domain := os.Getenv("DOMAIN")

	var shortUrl string
	if req.Alias == nil {
		count,err:= s.repo.IncrementCounter()
		if err != nil {
			return "",errs.InternalServerError(err.Error())
		}
	shortUrl =	helpers.EncodeBase62(count)
		// to do generate the long url using the base 62 function

	} else {
		shortUrl = *req.Alias

		// to do validate the alias 
		// create the url with the provided alias 
		// the duplication of the alias will be handled by the database
	}

	u:= url.URL{
		LongUrl: req.LongUrl,
		ShortUrl: shortUrl,
	}
	
	err := s.repo.CreateURL(u,c)
	if err != nil {

		if strings.HasPrefix(err.Error(), "duplicate entry for") {
			return  "" , errs.Conflict(err.Error())
		}
		
		return "",err
	}
	return domain+"/"+shortUrl, nil
}

func (s *URLService) GetURL(shortURL string , c context.Context) (*url.URL, error) {
	cacheKey:= "URL_"+shortURL
    cached, err := s.redis.Get(c, cacheKey).Result()
	if err == nil {
		log.Print("from redis")

		return &url.URL{LongUrl: cached}, nil
	}
	if err != redis.Nil {
    log.Println("Redis get error:", err)

	}

	u,err := s.repo.GetURL(shortURL)
	if err != nil {
		return nil, err
	}
err  = 	s.redis.Set(c,cacheKey, u.LongUrl, time.Minute*10).Err()
if err != nil {
    log.Println("Redis set error:", err)
}
	return u, nil
}