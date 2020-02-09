package album

import (
	"io/ioutil"
	"net/http"

	"github.com/continuouslylearning/mosaic/api/database"
	"github.com/gin-gonic/gin"
)

func getAlbum(c *gin.Context) {
	username := c.MustGet("username").(string)
	s3 := c.MustGet("s3").(*database.S3Handler)
	objects, e := s3.GetImages(username)
	if e != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, objects)
}

func postImages(c *gin.Context) {
	username := c.MustGet("username").(string)
	s3 := c.MustGet("s3").(*database.S3Handler)
	form, _ := c.MultipartForm()
	for name, file := range form.File {
		fileContent, _ := file[0].Open()
		body, _ := ioutil.ReadAll(fileContent)
		e := s3.UploadImage(name, body, username)
		if e != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}

	c.AbortWithStatus(http.StatusCreated)
}
