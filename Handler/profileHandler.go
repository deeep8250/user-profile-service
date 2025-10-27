package handler

import (
	"fmt"
	"net/http"
	service "profiles/Service"
	"profiles/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	r *service.ProfileService
}

/// For User

func NewProfileHandler(service *service.ProfileService) *ProfileHandler {
	return &ProfileHandler{r: service}
}

func (r *ProfileHandler) GetProfileByEmail(c *gin.Context) {
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

	result, err := r.r.GetProfileByEmail(user.Email)
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

func (r *ProfileHandler) UpdateProfileHandler(c *gin.Context) {
	id := c.Query("user_id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.Profile
	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"eror": err.Error(),
		})
		return
	}

	result, err := r.r.UpdateProfile(idInt, user)
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

func (s *ProfileHandler) SoftDelete(c *gin.Context) {
	fmt.Println("enter into delete handler ")
	id := c.DefaultQuery("id", "")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("enter into handler ")
	result, err := s.r.DeleteUser(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})

}
