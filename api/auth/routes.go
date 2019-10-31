package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func loginUser(c *gin.Context) {
	token := c.MustGet("token").(string)
	c.JSON(http.StatusCreated, map[string]string{
		"authToken": token,
	})
}
