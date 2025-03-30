package routes

import (
	"context"
	"net/http"
	"phrasmotica/flyer-api/database"
	"phrasmotica/flyer-api/logging"
	"phrasmotica/flyer-api/models"

	"github.com/gin-gonic/gin"
)

var db = database.CreateDb()

// PostFlyer     godoc
// @Summary      Gets all existing flyers
// @Description  Gets all existing flyers
// @Tags         Flyers
// @Produce      json
// @Security     BearerAuth
// @Success      200
// @Failure      401
// @Failure      503
// @Router       /flyer [get]
func GetFlyers(c *gin.Context) {
	ctx := context.TODO()

	success, flyers := db.GetFlyers(ctx)

	if !success {
		logging.Error.Println("Could not get flyers")
		c.AbortWithStatus(http.StatusServiceUnavailable)
		return
	}

	c.IndentedJSON(http.StatusOK, flyers)
}

// PostFlyer     godoc
// @Summary      Creates a new flyer
// @Description  Creates a new flyer
// @Tags         Flyers
// @Produce      json
// @Param        flyer body contracts.CreateFlyerRequest true "The flyer"
// @Security     BearerAuth
// @Success      204
// @Failure      400
// @Failure      503
// @Failure      401
// @Router       /flyer [post]
func PostFlyer(c *gin.Context) {
	var newFlyer models.Flyer

	if err := c.BindJSON(&newFlyer); err != nil {
		logging.Error.Println("Invalid body format")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx := context.TODO()

	if success := db.AddFlyer(ctx, &newFlyer); !success {
		logging.Error.Println("Could not add flyer")
		c.AbortWithStatus(http.StatusServiceUnavailable)
		return
	}

	logging.Info.Printf("Added flyer %s\n", newFlyer.ID)

	c.IndentedJSON(http.StatusNoContent, newFlyer)
}
