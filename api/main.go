package main

import (
	"github.com/continuouslylearning/albumify/api/album"
	"github.com/continuouslylearning/albumify/api/database"
	"github.com/continuouslylearning/albumify/api/users"

	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	clientOrigin := os.Getenv("CLIENT_ORIGIN")
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{clientOrigin},
		AllowMethods: []string{"DELETE", "GET", "OPTIONS", "POST"},
		AllowHeaders: []string{"Authorization", "Content-Length", "Content-Type", "Origin"},
	}))

	database.InitializeDB(r)
	database.InitializeS3Handler(r)
	redis := database.InitializeRedis()
	defer redis.Pool.Close()

	users.GroupUserRoutes(r)
	album.GroupAlbumRoutes(r)

	r.Run()
}
