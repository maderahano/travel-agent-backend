package main

import (
	"os"

	"github.com/labstack/echo/v4"

	conf "travel-agent-backend/config"
	docs "travel-agent-backend/docs"
	rest "travel-agent-backend/internal/route"

	echoSwag "github.com/swaggo/echo-swagger"
)

// @title        Travel Agent System API Documentation
// @description  This is Travel Agent System API
// @version      2.0
// @host         localhost:8888
// @BasePath
// @schemes                     http https
// @securityDefinitions.apiKey  JWT
// @in                          header
// @name                        Authorization
func main() {
	config := conf.InitConfig()
	e := echo.New()

	rest.RegisterAuthGroupAPI(e, config)
	rest.RegisterUserGroupAPI(e, config)
	rest.RegisterLocationGroupAPI(e, config)
	rest.RegisterOrderGroupAPI(e, config)
	rest.RegisterProductGroupAPI(e, config)
	rest.RegisterProductPhotoGroupAPI(e, config)
	rest.RegisterTravelAgentGroupAPI(e, config)
	rest.RegisterTravelAgentSocialMediaGroupAPI(e, config)

	e.GET("/swagger/*", echoSwag.WrapHandler)
	docs.SwaggerInfo.Host = os.Getenv("APP_HOST")

	e.Logger.Fatal(e.Start(config.SERVER_ADDRESS))
}
