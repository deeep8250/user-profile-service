package routes

import (
	handler "profiles/Handler"

	"github.com/gin-gonic/gin"
)

func Routes(c *gin.Engine, handler *handler.UserHandler) {

	//users
	c.POST("/create-users", handler.NewUser)
	c.GET("/users", handler.GetUserByEmail)
	c.GET("/List-of-users")
	c.PUT("/update-user")
	c.DELETE("/delete-user")

	//Profiles
	c.POST("/create-profiles")
	c.GET("/profiles/:id")
	c.GET("/List-of-profiles")
	c.PUT("/update-profiles")
	c.DELETE("/delete-profiles")
}
