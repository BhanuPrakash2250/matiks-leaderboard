package handlers

import (
	"net/http"

	"backend/services"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	q := c.Query("q")
	c.JSON(http.StatusOK, services.SearchUsers(q))
}
