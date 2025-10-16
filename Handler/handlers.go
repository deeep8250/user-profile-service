package handler

import (
	"fmt"
	"net/http"
	service "profiles/Service"
	"profiles/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	r *service.UserService
}

func (s *UserHandler) NewProfileHandler(profileService *service.ProfileService) any {
	panic("unimplemented")
}

/// For User

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

	fmt.Println("enter into the get all  user handler")

	pageStr := c.DefaultQuery("page", "1")
	pageLimitStr := c.DefaultQuery("page_size", "10")
	sort_by := c.DefaultQuery("sort_by", "created_at")
	order := c.DefaultQuery("order", "asc")
	filter := c.DefaultQuery("name", "")

	page, _ := strconv.Atoi(pageStr)
	pageLimit, _ := strconv.Atoi(pageLimitStr)

	result, err := s.r.GetAllusers(page, pageLimit, sort_by, order, filter)
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

func (s *UserHandler) SoftDelete(c *gin.Context) {
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
