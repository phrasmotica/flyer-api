package routes

import (
	"phrasmotica/flyer-api/database"
	"phrasmotica/flyer-api/logging"

	"github.com/gin-gonic/gin"
)

type IRoutes interface {
	GetFlyers(c *gin.Context)
	PostFlyer(c *gin.Context)

	GetPlayers(c *gin.Context)
	PostPlayer(c *gin.Context)
}

type Routes struct {
	Db     database.IDatabase
	Logger logging.ILogger
}
