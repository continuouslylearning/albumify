package album

import (
	"io/ioutil"
	"net/http"

	"github.com/continuouslylearning/albumify/api/database"
	"github.com/gin-gonic/gin"
)

func deleteImage(c *gin.Context) {
	username := c.MustGet("username").(string)
	s3 := c.MustGet("s3").(*database.S3Handler)

	var req DeleteRequest
	e := c.ShouldBind(&req)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Must specify an object key",
		})
		return
	}

	key := req.Key
	e = s3.DeleteImage(key, username)
	if e != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}

func getAlbum(c *gin.Context) {
	username := c.MustGet("username").(string)
	s3 := c.MustGet("s3").(*database.S3Handler)
	objects, e := s3.GetImages(username)
	if e != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
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
			return
		}
	}

	c.AbortWithStatus(http.StatusCreated)
}
