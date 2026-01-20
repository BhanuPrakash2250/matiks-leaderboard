package handlers

import (
	"net/http"
	"strconv"

	"backend/services"

	"github.com/gin-gonic/gin"
)

func GetLeaderboard(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	c.JSON(http.StatusOK, services.GetUsers(limit))
}
