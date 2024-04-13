package main

import (
	"context"
	"github.com/gin-gonic/gin"
	filesSwag "github.com/swaggo/files"
	ginSwag "github.com/swaggo/gin-swagger"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use()
	url := ginSwag.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwag.WrapHandler(filesSwag.Handler, url))

	srv := http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
