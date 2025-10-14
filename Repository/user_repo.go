package repository

import (
	"errors"
	"profiles/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	q *gorm.DB
}

// NOTE:
// NewUserRepo() is a constructor function.
// It creates a new instance of UserRepository (in memory),
// stores the given *Db.Queries value inside that struct,
// and returns a pointer to it (*UserRepository).
// Even though we don’t write "UserRepository{...}" in main.go,
// this function does it internally and gives us that ready-to-use object.

func NewUserRepo(q *gorm.DB) *UserRepository {

	return &UserRepository{q: q}
}

func (r *UserRepository) CreateUser(arg *models.User) error {
	return r.q.Create(arg).Error
}

func (r *UserRepository) GetUser(email string) (*models.User, error) {
	var User models.User
	result := r.q.Preload("Profile").Where("email=?", email).First(&User)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &User, nil

}

// func (r *UserRepository) Update(id int64) error {
// var User models.

// }

// func (r *UserRepository) Delete (id int64) error {}
