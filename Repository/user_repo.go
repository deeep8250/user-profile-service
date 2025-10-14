package repository

import (
	"errors"
	"fmt"

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

func (r *UserRepository) GetAllUser() ([]models.User, error) {

	var user []models.User
	result := r.q.Find(&user)
	if result.Error != nil {
		return []models.User{}, fmt.Errorf("error from get all users: %w", result.Error)
	}

	return user, nil

}

func (r *UserRepository) Update(id int64, updates models.User) (models.User, error) {
	var user2 models.User
	result := r.q.Preload("Profile").Where("id=?", id).First(&user2)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	// prevent changing ID or timestamps
	if updates.CreatedAt != nil {
		return models.User{}, fmt.Errorf("restricted field in use")
	}
	if updates.Deleted != nil {
		return models.User{}, fmt.Errorf("restricted field in use")
	}
	if updates.ID != nil {
		return models.User{}, fmt.Errorf("restricted field in use")
	}
	if updates.UpdatedAt != nil {
		return models.User{}, fmt.Errorf("restricted field in use")
	}

	if err := r.q.Model(&user2).Updates(updates).Error; err != nil {
		return models.User{}, err
	}

	if err := r.q.Preload("Profile").Where("id=?", id).First(&user2).Error; err != nil {
		return models.User{}, err
	}

	return user2, nil

}

// func (r *UserRepository) Delete (id int64) error {}
