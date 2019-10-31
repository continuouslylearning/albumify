package album

import (
	"github.com/continuouslylearning/mosaic/api/database"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func getAlbum(c *gin.Context) {
	s3, _ := c.MustGet("s3").(*database.S3Handler)
	objects, _ := s3.GetImages()
	c.JSON(http.StatusOK, objects)
}

func postImages(c *gin.Context) {
	s3, _ := c.MustGet("s3").(*database.S3Handler)
	form, _ := c.MultipartForm()
	for name, file := range form.File {
		fileContent, _ := file[0].Open()
		body, _ := ioutil.ReadAll(fileContent)
		s3.UploadImage(name, body)
	}

	c.AbortWithStatus(http.StatusCreated)
}
