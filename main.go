package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/omarsabri666/url_shorter/config"
	"github.com/omarsabri666/url_shorter/handler"
	"github.com/omarsabri666/url_shorter/middleware"
	"github.com/omarsabri666/url_shorter/repository"
	"github.com/omarsabri666/url_shorter/routes"
	urlService "github.com/omarsabri666/url_shorter/service/url"
	service "github.com/omarsabri666/url_shorter/service/user"
	"github.com/omarsabri666/url_shorter/validators"
)

func init() {
	
	  if err := godotenv.Load(); err != nil {
        log.Println(".env file not found, relying on system environment variables")
    }
	db := config.Connect()
	ctx := context.Background()
	rdb:= config.NewRedisClient()
		pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	validators.RegisterValidators()
	r := gin.Default()
	r.Use(middleware.RequestIdMiddleware())
	r.Use(middleware.RateLimiterMiddleware(rdb, time.Minute, 60))
	// r.Use(middleware.AuthMiddleware())
	api:= r.Group("/")
	

	fmt.Println("Redis says:", pong)
	// repositories
	urlRepo := repository.NewSqlUrlRepository(db)
	userRepo := repository.NewDbUserRepository(db)
	// old repo using mysql
	// tokenRepo := repository.NewTokenRepositoryDb(db)
	// new repo using redis 
	tokenRepo := repository.NewTokenRepositoryRedis(rdb)


	// services
	urlService:= urlService.NewUrlService(urlRepo,rdb)
	userService:= service.NewUserService(userRepo,tokenRepo)



	// handlers
	urlHandler:= handler.NewURLHandler(urlService)
	userHandler := handler.NewUserHandler(userService)


	routes.RegisterUserRoutes(api,userHandler)
	routes.RegisterURLRoutes(api,urlHandler)
	defer db.Close()
	port:= os.Getenv("PORT")
	log.Println(port)
		r.Run(":"+port)
}
func main() {

	
	


}