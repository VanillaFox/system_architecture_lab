package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	filesSwag "github.com/swaggo/files"
	ginSwag "github.com/swaggo/gin-swagger"

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

	jwtSecretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	v1.InitJwtSecretKey(jwtSecretKey)

	repository := postgres.NewRepository(pool)

	userService := services.NewUserService(repository)

	authService := services.NewAuthService(repository, jwtSecretKey)

	routerHandler := restapi.NewRouter(
		v1.NewUserHandler(userService),
		auth.NewAuthHandler(authService),
	)

	r := gin.Default()
	r.Use()
	url := ginSwag.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwag.WrapHandler(filesSwag.Handler, url))
	routerHandler.SetRoutes(r)

	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
