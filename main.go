package main

import (
	"log"
	"os"
	"phrasmotica/flyer-api/auth"
	"phrasmotica/flyer-api/database"
	docs "phrasmotica/flyer-api/docs/flyer"
	"phrasmotica/flyer-api/logging"
	"phrasmotica/flyer-api/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/dig"
)

func createLogger() logging.ILogger {
	return &logging.ConsoleLogger{
		InfoLevel:  log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile),
		ErrorLevel: log.New(os.Stdout, "ERROR: ", log.LstdFlags|log.Lshortfile),
	}
}

func createDb(logger logging.ILogger) database.IDatabase {
	factory := database.DatabaseFactory{
		Logger: logger,
	}

	return factory.CreateDb()
}

func createMiddleware(logger logging.ILogger) *auth.Middleware {
	return &auth.Middleware{
		Logger: logger,
	}
}

func createRoutes(db database.IDatabase, logger logging.ILogger) routes.IRoutes {
	return &routes.Routes{
		Db:     db,
		Logger: logger,
	}
}

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
	digContainer := dig.New()
	digContainer.Provide(createLogger)
	digContainer.Provide(createDb)
	digContainer.Provide(createMiddleware)
	digContainer.Provide(createRoutes)
	digContainer.Invoke(runRouter)
}

func runRouter(r routes.IRoutes, m *auth.Middleware) {
	router := gin.Default()

	router.Use(m.CORSMiddleware())

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	flyers := router.Group("/flyer", m.TokenAuth(true))
	{
		flyers.GET("", r.GetFlyers)
		flyers.POST("", r.PostFlyer)
	}

	router.Run(":8000")
}
