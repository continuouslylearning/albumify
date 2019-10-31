package album

import (
	"github.com/gin-gonic/gin"
)

func GroupAlbumRoutes(r *gin.Engine) {
	r.GET("/album", getAlbum)
	r.POST("/album", postImages)
}
