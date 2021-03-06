package album

import (
	"albumify/users"

	"github.com/gin-gonic/gin"
)

func GroupAlbumRoutes(r *gin.Engine) {
	albumRoutes := r.Group("/album")
	{
		albumRoutes.Use(users.JwtAuth())
		albumRoutes.GET("/", getAlbum)
		albumRoutes.DELETE("/", deleteImage)
		albumRoutes.POST("/", postImages)
	}
}
