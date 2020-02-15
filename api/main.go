package main

import (
	"albumify/album"
	"albumify/db"
	"albumify/users"

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

	db.InitializeDB(r)
	db.InitializeS3Handler(r)
	redis := db.InitializeRedis(r)
	defer redis.Pool.Close()

	users.GroupUserRoutes(r)
	album.GroupAlbumRoutes(r)

	r.Run()
}
