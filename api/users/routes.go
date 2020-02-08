package users

import "github.com/gin-gonic/gin"

func GroupUserRoutes(r *gin.Engine) {
	r.POST("/login", localAuth, loginUser)
	r.POST("/users", createUser)
}
