package service

import (
	repository "profiles/Repository"
	Db "profiles/internal/db"
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

func (s *UserService) RegisterUser(user Db.User) {

}

func (s *UserService) GetUserByID(userId int64) {

}
