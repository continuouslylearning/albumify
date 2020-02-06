package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func getUsers(c *gin.Context) {
	var users []User
	db := c.MustGet("db").(*gorm.DB)
	db.Find(&users)
	normalizedUsers := make([]map[string]interface{}, len(users), len(users))
	for i, u := range users {
		normalizedUsers[i] = u.Normalize()
	}
	c.JSON(http.StatusOK, normalizedUsers)
}

func getUserByID(c *gin.Context) {
	var user User
	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)

	e := db.Where("id = ?", id).First(&user).Error
	if e != nil {
		if gorm.IsRecordNotFoundError(e) {
			c.Status(http.StatusNotFound)
		} else {
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, user.Normalize())
}

func deleteUser(c *gin.Context) {
	userID := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)

	e := db.Where("id = ?", userID).Delete(&User{}).Error
	if e != nil {
		if gorm.IsRecordNotFoundError(e) {
			c.Status(http.StatusNotFound)
		} else {
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	c.Status(http.StatusNoContent)
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

func updateUser(c *gin.Context) {
	var updatedUser User
	userID := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)

	e := c.ShouldBindJSON(&updatedUser)
	if e != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if len(updatedUser.Password) != 0 {
		hash, _ := hashPassword(updatedUser.Password)
		updatedUser.Password = hash
	}

	e = db.Model(&updatedUser).Where("id = ?", userID).Update(&updatedUser).First(&updatedUser).Error
	if e != nil {
		if gorm.IsRecordNotFoundError(e) {
			c.Status(http.StatusNotFound)
		} else {
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, updatedUser.Normalize())
}
