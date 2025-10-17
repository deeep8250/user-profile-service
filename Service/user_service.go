package service

import (
	"fmt"
	repository "profiles/Repository"
	"profiles/models"

	"golang.org/x/crypto/bcrypt"
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

func (s *UserService) RegisterUser(user *models.User) (string, error) {

	result, err := s.r.GetUser(user.Email)
	if err != nil {
		fmt.Println("error 1: ", err.Error())
		return "", fmt.Errorf("error 2: %w", err)
	}

	if result != nil {
		fmt.Println("User already exist")
		return "User already exist", nil
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("password hashing failed")
	}

	user.Password = string(hashPass)

	err = s.r.CreateUser(user)
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
	if result == nil {
		return models.User{}, fmt.Errorf("no result found")
	}
	fmt.Println("service layer email :", result)
	return *result, nil
}

func (s *UserService) GetAllusers(page int, pageLimit int, sort_by string, order, filter string) ([]models.User, error) {

	result, err := s.r.GetAllUser(page, pageLimit, sort_by, order, filter)
	if err != nil {
		return []models.User{}, err
	}

	return result, nil
}

func (s *UserService) UpdateUser(id int64, user models.User) (models.User, error) {
	result, err := s.r.Update(id, user)
	if err != nil {
		return models.User{}, err
	}
	return result, nil

}

func (s *UserService) DeleteUser(id int) (models.User, error) {
	fmt.Println("enter into the service")
	result, err := s.r.Delete(id)
	if err != nil {
		return models.User{}, err
	}
	return result, nil
}
