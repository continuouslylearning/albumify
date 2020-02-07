package users

import "github.com/gin-gonic/gin"

func GroupUserRoutes(r *gin.Engine) {
	r.POST("/login", loginUser)
	r.POST("/users", createUser)
}
