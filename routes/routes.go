package routes

import "github.com/gin-gonic/gin"

func Routes(c *gin.Engine) {

	//users
	c.POST("/create-users")
	c.GET("/users/:id")
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
