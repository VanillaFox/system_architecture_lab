package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/VanillaFox/system_architecture_lab/api-gateway/pkg/circuitbreaker"
	"github.com/VanillaFox/system_architecture_lab/api-gateway/restapi"
	"github.com/VanillaFox/system_architecture_lab/api-gateway/restapi/proxy"

	_ "github.com/VanillaFox/system_architecture_lab/docs/apigateway"
	"github.com/gin-gonic/gin"
	filesSwag "github.com/swaggo/files"
	ginSwag "github.com/swaggo/gin-swagger"
)

func main() {
	usersHost := os.Getenv("USERS_HOST")
	usersPort := os.Getenv("USERS_PORT")

	conferencesHost := os.Getenv("CONFERENCES_HOST")
	conferencesPort := os.Getenv("CONFERENCES_PORT")

	users, err := url.Parse(fmt.Sprintf("http://%s:%s", usersHost, usersPort))
	if err != nil {
		panic(err)
	}

	usersProxy := httputil.NewSingleHostReverseProxy(users)

	conferences, err := url.Parse(fmt.Sprintf("http://%s:%s", conferencesHost, conferencesPort))
	if err != nil {
		panic(err)
	}

	conferencesProxy := httputil.NewSingleHostReverseProxy(conferences)

	usersCB := circuitbreaker.NewCircuitBreaker(2, 2, time.Second*10, time.Second*60)
	conferencesCB := circuitbreaker.NewCircuitBreaker(2, 2, time.Second*10, time.Second*50)

	routerHandler := restapi.NewRouter(
		proxy.NewConferencesHandler(conferencesProxy, conferencesCB),
		proxy.NewUsersHandler(usersProxy, usersCB),
	)

	r := gin.Default()

	url := ginSwag.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwag.WrapHandler(filesSwag.Handler, url))

	routerHandler.SetRoutes(r)

	srv := http.Server{
		Addr:    ":8082",
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
