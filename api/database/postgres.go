package database

import (
	"fmt"
	"os"

	"github.com/continuouslylearning/mosaic/api/users"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitializeDB(r *gin.Engine) *gorm.DB {
	pgConfig := os.Getenv("POSTGRES_CONFIG")
	db, err := gorm.Open("postgres", pgConfig)
	if err != nil {
		fmt.Println(err)
		panic("Could not connect to the database")
	}

	db.AutoMigrate(&users.User{})
	r.Use(AddDB(db))

	return db
}

func AddDB(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
