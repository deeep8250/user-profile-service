package service

import (
	"fmt"
	repository "profiles/Repository"
	"profiles/models"
)

// this struct is for accessing the repository layer's function with db connection
type UserService struct {
	r *repository.UserRepository
}

// NOTE:
// NewUserRepo() is a constructor function.
// It creates a new instance of UserRepository (in memory),
// stores the given *Db.Queries value inside that struct,
// and returns a pointer to it (*UserRepository).
// Even though we don’t write "UserRepository{...}" in main.go,
// this function does it internally and gives us that ready-to-use object.

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{r: repo}
}

func (s *UserService) RegisterUser(user models.User) (string, error) {

	result, err := s.r.GetUser(user.Email)
	if err != nil {
		fmt.Println("error 1: ", err.Error())
		return "", fmt.Errorf("error 2: %w", err)
	}

	if result != nil {
		fmt.Println("User already exist")
		return "User already exist", nil
	}

	err = s.r.CreateUser(&user)
	if err != nil {
		return "", fmt.Errorf("error 2 : %w", err)
	}
	return "User Register successfully", nil

}

func (s *UserService) GetUserByEmail(email string) (models.User, error) {
	result, err := s.r.GetUser(email)
	if err != nil {
		return models.User{}, err
	}
	return *result, nil
}
