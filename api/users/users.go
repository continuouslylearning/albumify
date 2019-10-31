package users

import "github.com/gin-gonic/gin"

func GroupUserRoutes(r *gin.Engine) {
	usersRouter := r.Group("/users")
	{
		usersRouter.GET("/", getUsers)
		usersRouter.POST("/", createUser)
		usersRouter.GET(":id", getUserByID)
		usersRouter.DELETE("/:id", deleteUser)
		usersRouter.PUT("/:id", updateUser)
	}
}
