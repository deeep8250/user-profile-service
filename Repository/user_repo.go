package repository

import (
	"errors"
	"fmt"
	"sort"
	"strings"

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

	tx := r.q.Begin()
	fmt.Println("→ Transaction started")

	// 1️⃣ Insert user
	if err := tx.Create(arg).Error; err != nil {
		fmt.Println("❌ user insert failed:", err)
		tx.Rollback()
		return err
	}
	fmt.Println("✅ user inserted, ID:", arg.ID)

	// 2️⃣ Insert related profile
	profile := models.Profile{UserID: arg.ID}
	if err := tx.Create(&profile).Error; err != nil {
		fmt.Println("❌ profile insert failed:", err)
		tx.Rollback()
		return err
	}

	// 3️⃣ Commit
	if err := tx.Commit().Error; err != nil {
		fmt.Println("❌ commit failed:", err)
		return err
	}
	fmt.Println("✅ Transaction committed successfully")
	return nil

	// if err := r.q.Create(arg).Error; err != nil {
	// 	fmt.Println("❌ insert failed:", err)
	// }
	// fmt.Println("✅ inserted ID:", arg.ID)
	// return nil

}

func (r *UserRepository) GetUser(email string) (*models.User, error) {
	fmt.Println("entered in to repo")
	var User models.User
	result := r.q.Preload("Profile").Where("email=?", email).First(&User)

	if result.Error != nil || User.Deleted {
		if User.Deleted {
			return &models.User{}, fmt.Errorf("user is already deleted")
		}
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	fmt.Println("user  is ", User)
	return &User, nil

}

func (r *UserRepository) GetAllUser(page int, Limit int, sort_by, order, filter string) ([]models.User, error) {

	// get the specific rows
	offset := (page - 1) * Limit
	fmt.Printf("limit : %d and offset : %d", Limit, offset)
	rows, err := r.q.Raw("SELECT  * from users Where Deleted=false LIMIT $1 OFFSET $2", Limit, offset).Rows()
	if err != nil {
		return []models.User{}, err
	}

	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var a models.User
		var b models.User
		r.q.ScanRows(rows, &a)
		r.q.Preload("Profile").Where("id=?", a.ID).First(&b)
		users = append(users, b)

	}

	// sort_by
	sort.Slice(users, func(i, j int) bool {
		if order == "asc" {
			return users[i].CreatedAt.Before(*users[j].CreatedAt)
		} else {
			return users[i].CreatedAt.After(*users[j].CreatedAt)
		}
	})

	//filter
	var filtered_data []models.User
	for _, u := range users {
		if strings.Contains(u.Name, filter) {
			filtered_data = append(filtered_data, u)
		}
	}

	return filtered_data, nil

}

func (r *UserRepository) Update(id int64, updates models.User) (models.User, error) {
	var user2 models.User
	result := r.q.Preload("Profile").Where("id=?", id).First(&user2)
	if result.Error != nil || user2.Deleted {
		if user2.Deleted {
			return models.User{}, fmt.Errorf("invalid user")
		}
		return models.User{}, result.Error
	}
	// prevent changing ID or timestamps
	if updates.CreatedAt != nil {
		return models.User{}, fmt.Errorf("restricted field in use")
	}
	if updates.Deleted {
		return models.User{}, fmt.Errorf("restricted field in use")
	}
	if updates.ID != 0 {
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

func (r *UserRepository) Delete(id int) (models.User, error) {
	fmt.Println("enter into the repo")
	var user models.User
	result := r.q.Preload("Profile").Where("id=?", id).First(&user)
	if result.Error != nil || user.Deleted {
		if user.Deleted {
			return models.User{}, fmt.Errorf("user already deleted")
		}
		fmt.Println("enter into the  if state")
		return models.User{}, result.Error

	}

	fmt.Println("result : ", result)

	user.Deleted = true
	user.Profile.Deleted = true
	r.q.Save(&user)
	return user, nil
}
