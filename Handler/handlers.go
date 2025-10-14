package handler

import (
	"net/http"
	service "profiles/Service"
	"profiles/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	r *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{r: service}
}

func (s *UserHandler) NewUser(c *gin.Context) {

	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response": "invalid request",
		})
		return
	}

	result, err := s.r.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": result,
	})

}

func (s *UserHandler) GetUserByEmail(c *gin.Context) {
	type User struct {
		Email string `json:"email" binding:"required,email"`
	}
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := s.r.GetUserByEmail(user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": result,
	})

}
func (s *UserHandler) GetAllUsers(c *gin.Context) {

	result, err := s.r.GetAllusers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": result,
	})
}

func (s *UserHandler) UpdateUserHandler(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User
	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"eror": err.Error(),
		})
		return
	}

	result, err := s.r.UpdateUser(idInt, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": result,
	})

}
