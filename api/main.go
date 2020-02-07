package main

import (
	"github.com/continuouslylearning/mosaic/api/album"
	"github.com/continuouslylearning/mosaic/api/database"
	"github.com/continuouslylearning/mosiac/api/users"

	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Could not load Environmental Variables")
	}

	r := gin.New()

	config := cors.DefaultConfig()
	clientOrigin := os.Getenv("CLIENT_ORIGIN")
	config.AllowOrigins = []string{clientOrigin}
	r.Use(cors.New(config))

	database.InitializeDB(r)
	database.InitializeS3Handler(r)

	auth.GroupAuthRoutes(r)
	users.GroupUserRoutes(r)
	album.GroupAlbumRoutes(r)

	r.Run()
}
