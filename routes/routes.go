package routes

import (
	handler "profiles/Handler"

	"github.com/gin-gonic/gin"
)

func Routes(c *gin.Engine, handler *handler.UserHandler) {

	//users
	c.POST("users/create-users", handler.NewUser)
	c.GET("users/users", handler.GetUserByEmail)
	c.GET("users/List-of-users", handler.GetAllUsers)
	c.PUT("users/update-user/:id", handler.UpdateUserHandler)
	c.DELETE("users/delete-user")

	//Profiles
	c.POST("profiles/create-profiles")
	c.GET("profiles/profiles/:id")
	c.GET("profiles/List-of-profiles")
	c.PUT("profiles/update-profiles")
	c.DELETE("profiles/delete-profiles")
}
