package routes

import (
	"context"
	"net/http"
	"phrasmotica/flyer-api/contracts"
	"phrasmotica/flyer-api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// PostPlayers   godoc
// @Summary      Gets all existing players
// @Description  Gets all existing players
// @Tags         Players
// @Produce      json
// @Security     BearerAuth
// @Success      200
// @Failure      401
// @Failure      503
// @Router       /player [get]
func (r *Routes) GetPlayers(c *gin.Context) {
	ctx := context.TODO()

	success, players := r.Db.GetPlayers(ctx)

	if !success {
		r.Logger.Error("Could not get players")
		c.AbortWithStatus(http.StatusServiceUnavailable)
		return
	}

	c.IndentedJSON(http.StatusOK, players)
}

// PostPlayers   godoc
// @Summary      Creates a new player
// @Description  Creates a new player
// @Tags         Players
// @Produce      json
// @Param        player body contracts.CreatePlayerRequest true "The player"
// @Security     BearerAuth
// @Success      204
// @Failure      400
// @Failure      503
// @Failure      401
// @Router       /player [post]
func (r *Routes) PostPlayer(c *gin.Context) {
	var request contracts.CreatePlayerRequest

	if err := c.BindJSON(&request); err != nil {
		r.Logger.Error("Invalid body format")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx := context.TODO()

	newPlayer := models.Player{
		ID:   uuid.NewString(),
		Name: request.Name,
	}

	if success := r.Db.AddPlayer(ctx, &newPlayer); !success {
		r.Logger.Error("Could not add player")
		c.AbortWithStatus(http.StatusServiceUnavailable)
		return
	}

	r.Logger.Info("Added player", "ID", newPlayer.ID)

	c.IndentedJSON(http.StatusNoContent, newPlayer)
}
