package routes

import (
	"profiles/handler"

	"github.com/gin-gonic/gin"
)

func Routes(c *gin.Engine, handler *handler.UserHandler, profileHandler *handler.ProfileHandler) {

	//users
	c.POST("users/create-users", handler.NewUser)
	c.GET("users/users", handler.GetUserByEmail)
	c.GET("users/List-of-users", handler.GetAllUsers)
	c.PUT("users/update-user/:id", handler.UpdateUserHandler)
	c.DELETE("users/delete-user", handler.SoftDelete)

	//Profiles
	// c.POST("profiles/create-profiles",profileHandler.)
	c.GET("profiles/Getprofiles", profileHandler.GetProfileByEmail)
	// c.GET("profiles/List-of-profiles", profileHandler)
	c.PUT("profiles/update-profiles", profileHandler.UpdateProfileHandler)
	c.DELETE("profiles/delete-profiles", profileHandler.SoftDelete)
}
