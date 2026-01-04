package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/omarsabri666/url_shorter/config"
	"github.com/omarsabri666/url_shorter/global"
	"github.com/omarsabri666/url_shorter/handler"
	"github.com/omarsabri666/url_shorter/middleware"
	"github.com/omarsabri666/url_shorter/repository"
	urlService "github.com/omarsabri666/url_shorter/service/url"
	service "github.com/omarsabri666/url_shorter/service/user"
	"github.com/omarsabri666/url_shorter/validators"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using system env")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// DB + Redis
	db := config.Connect()
	defer db.Close()
	rdb := config.NewRedisClient()
	ctx := context.Background()
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Redis says:", pong)

	// Validators
	validators.RegisterValidators()

	// Repositories
	urlRepo := repository.NewSqlUrlRepository(db)
	userRepo := repository.NewDbUserRepository(db)
	tokenRepo := repository.NewTokenRepositoryRedis(rdb)

	// Services
	urlService := urlService.NewUrlService(urlRepo, rdb)
	userService := service.NewUserService(userRepo, tokenRepo)

	// Handlers
	urlHandler := handler.NewURLHandler(urlService)
	userHandler := handler.NewUserHandler(userService)

	// Rate limiter
	rateLimiter := func(next http.Handler) http.Handler {
		return middleware.RateLimiterMiddlewareHttp(next, rdb, time.Minute, 60)
	}

	// Global middleware stack
	stack := global.Chain(
		middleware.RecoveryMiddleware,
		middleware.LoggerMiddleware,
		middleware.RequestIdMiddlewareHttp,
		rateLimiter,
	)

	// Routes
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{url}", urlHandler.GetURLHttp)
	mux.HandleFunc("POST /", urlHandler.CreateURLHttp)
	mux.HandleFunc("POST /auth/signup", userHandler.SignupHttp)
	mux.HandleFunc("POST /auth/signin", userHandler.LoginHttp)
	mux.Handle("POST /auth/signout", middleware.AuthMiddlewareHttp(http.HandlerFunc(userHandler.LogoutHttp)))
	mux.Handle("POST /auth/refresh", middleware.AuthMiddlewareHttp(http.HandlerFunc(userHandler.RefreshTokenHttp)))

	// Server
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      stack(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start server
	go func() {
		log.Println("Server running on port", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctxShut, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(ctxShut)
	log.Println("Server stopped")
}
