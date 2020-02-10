package main

import (
	"github.com/continuouslylearning/mosaic/api/album"
	"github.com/continuouslylearning/mosaic/api/database"
	"github.com/continuouslylearning/mosaic/api/users"

	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
	}

	r := gin.New()

	config := cors.DefaultConfig()
	clientOrigin := os.Getenv("CLIENT_ORIGIN")
	config.AllowOrigins = []string{clientOrigin}

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{clientOrigin},
		AllowMethods: []string{"DELETE", "GET", "OPTIONS", "POST"},
		AllowHeaders: []string{"Authorization", "Content-Length", "Content-Type", "Origin"},
	}))

	database.InitializeDB(r)
	database.InitializeS3Handler(r)

	users.GroupUserRoutes(r)
	album.GroupAlbumRoutes(r)

	r.Run()
}
