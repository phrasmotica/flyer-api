package routes

import (
	"context"
	"net/http"
	"phrasmotica/flyer-api/models"

	"github.com/gin-gonic/gin"
)

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
func (r *Routes) GetFlyers(c *gin.Context) {
	ctx := context.TODO()

	success, flyers := r.Db.GetFlyers(ctx)

	if !success {
		r.Logger.Error("Could not get flyers")
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
func (r *Routes) PostFlyer(c *gin.Context) {
	var newFlyer models.Flyer

	if err := c.BindJSON(&newFlyer); err != nil {
		r.Logger.Error("Invalid body format")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx := context.TODO()

	if success := r.Db.AddFlyer(ctx, &newFlyer); !success {
		r.Logger.Error("Could not add flyer")
		c.AbortWithStatus(http.StatusServiceUnavailable)
		return
	}

	r.Logger.Info("Added flyer", "ID", newFlyer.ID)

	c.IndentedJSON(http.StatusNoContent, newFlyer)
}
