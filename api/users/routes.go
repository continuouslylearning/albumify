package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func loginUser(c *gin.Context) {
	token := c.MustGet("token").(string)
	c.JSON(http.StatusCreated, map[string]string{
		"authToken": token,
	})
}

func createUser(c *gin.Context) {
	var newUser User
	db := c.MustGet("db").(*gorm.DB)

	e := c.ShouldBindJSON(&newUser)
	if e != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	e = db.Where("username = ?", newUser.Username).First(&User{}).Error
	if e != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	hash, _ := hashPassword(newUser.Password)
	newUser.Password = hash
	db.Create(&newUser)
	c.JSON(http.StatusCreated, newUser.Normalize())
}
