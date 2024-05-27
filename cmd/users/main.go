package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
	filesSwag "github.com/swaggo/files"
	ginSwag "github.com/swaggo/gin-swagger"

	"github.com/VanillaFox/system_architecture_lab/users/adaptres/cache"
	"github.com/VanillaFox/system_architecture_lab/users/adaptres/postgres"
	"github.com/VanillaFox/system_architecture_lab/users/app/services"
	"github.com/VanillaFox/system_architecture_lab/users/restapi"
	"github.com/VanillaFox/system_architecture_lab/users/restapi/auth"
	v1 "github.com/VanillaFox/system_architecture_lab/users/restapi/v1"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"

	_ "github.com/VanillaFox/system_architecture_lab/docs/users"
)

func main() {
	ctx := context.Background()

	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	pgConn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?target_session_attrs=read-write",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)

	cfg, err := pgxpool.ParseConfig(pgConn)
	if err != nil {
		panic(err)
	}

	cfg.ConnConfig.PreferSimpleProtocol = true

	pool, err := pgxpool.ConnectConfig(ctx, cfg)

	if err != nil {
		panic(err)
	}

	rdbHost := os.Getenv("REDIS_HOST")
	rdbPort := os.Getenv("REDIS_PORT")
	rdbPass := os.Getenv("REDIS_PASSWORD")

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rdbHost, rdbPort),
		Password: rdbPass,
		DB:       0,
	})

	jwtSecretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	v1.InitJwtSecretKey(jwtSecretKey)

	repository := postgres.NewRepository(pool)

	cache := cache.NewCache(rdb, repository)

	userService := services.NewUserService(repository, cache)

	authService := services.NewAuthService(repository, jwtSecretKey)

	routerHandler := restapi.NewRouter(
		v1.NewUserHandler(userService),
		auth.NewAuthHandler(authService),
	)

	r := gin.Default()
	r.Use()
	url := ginSwag.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwag.WrapHandler(filesSwag.Handler, url))
	r.GET("/healthcheck", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{}) })
	routerHandler.SetRoutes(r)

	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
