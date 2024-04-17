package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	mongoRepo "github.com/VanillaFox/system_architecture_lab/conferences/adapters/mongo"
	"github.com/VanillaFox/system_architecture_lab/conferences/app/services"
	"github.com/VanillaFox/system_architecture_lab/conferences/restapi"
	conferenceV1 "github.com/VanillaFox/system_architecture_lab/conferences/restapi/conferneces/v1"
	reportsV1 "github.com/VanillaFox/system_architecture_lab/conferences/restapi/reports/v1"

	"github.com/gin-gonic/gin"
	filesSwag "github.com/swaggo/files"
	ginSwag "github.com/swaggo/gin-swagger"

	_ "github.com/VanillaFox/system_architecture_lab/docs/conferences"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mongoHost := os.Getenv("MONGO_HOST")
	mongoPort := os.Getenv("MONGO_PORT")

	URI := fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	conferencesCollection := client.Database("conf").Collection("conferences")
	reportsCollections := client.Database("conf").Collection("reports")

	repository := mongoRepo.NewRepository(conferencesCollection, reportsCollections)

	conferenceService := services.NewConferenceService(repository)
	reportService := services.NewReportService(repository)

	routerHandler := restapi.NewRouter(
		conferenceV1.NewConferenceHandler(conferenceService),
		reportsV1.NewReportHandler(reportService),
	)

	r := gin.Default()
	r.Use()
	url := ginSwag.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwag.WrapHandler(filesSwag.Handler, url))
	routerHandler.SetRoutes(r)

	srv := http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
