package auth

import "github.com/gin-gonic/gin"

func GroupAuthRoutes(r *gin.Engine) {
	usersRouter := r.Group("/auth")
	{
		usersRouter.POST("/login", localAuth, loginUser)
	}
}
