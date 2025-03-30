package main

import (
	"phrasmotica/flyer-api/auth"
	docs "phrasmotica/flyer-api/docs/flyer"
	"phrasmotica/flyer-api/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// adapted from https://levelup.gitconnected.com/tutorial-generate-swagger-specification-and-swaggerui-for-gin-go-web-framework-9f0c038483b5, https://github.com/swaggo/gin-swagger

// @title Flyer API
// @version 0.1.0
// @description This is the Flyer API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	router := gin.Default()

	router.Use(auth.CORSMiddleware())

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	flyers := router.Group("/flyer", auth.TokenAuth(true))
	{
		flyers.GET("", routes.GetFlyers)
		flyers.POST("", routes.PostFlyer)
	}

	router.Run(":8000")
}
