package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func loginUser(c *gin.Context) {
	token := c.MustGet("token").(string)
	c.JSON(http.StatusCreated, gin.H{
		"authToken": token,
	})
}

func createUser(c *gin.Context) {
	var newUser User
	db := c.MustGet("db").(*gorm.DB)

	e := c.ShouldBindJSON(&newUser)
	if e != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Missing username or password",
		})
		return
	}

	e = db.Where("username = ?", newUser.Username).First(&User{}).Error
	if e == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "This username is taken",
		})
		return
	}

	digest, e := hashPassword(newUser.Password)
	if e != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	newUser.Password = digest
	db.Create(&newUser)

	var token string
	token, e = createToken(&newUser)
	if e != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"authToken": token,
	})
}
