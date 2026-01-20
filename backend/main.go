package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"backend/handlers"
	"backend/seed"
	"backend/services"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(cors.Default())

	services.SeedUsers(seed.GenerateUsers(10000))

	go func() {
		for {
			time.Sleep(3 * time.Second)
			services.RandomScoreUpdate()
		}
	}()

	r.GET("/leaderboard", handlers.GetLeaderboard)
	r.GET("/search", handlers.Search)

	r.Run("127.0.0.1:9090")
}
