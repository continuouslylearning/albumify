package users

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		parts := strings.Split(auth, " ")
		if len(parts) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		scheme := parts[0]
		token := parts[1]
		if scheme != "Bearer" || len(token) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		_, e := verifyToken(token)
		if e != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

func localAuth(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var req LoginRequest
	var user User

	e := c.ShouldBind(&req)
	if e != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	e = db.Where("username = ?", req.Username).Find(&user).First(&user).Error
	if e != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	valid := verifyPassword(user.Password, req.Password)
	if !valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, _ := createToken(&user)
	c.Set("token", token)
	c.Next()
}
