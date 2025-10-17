package repository

import (
	"errors"
	"fmt"
	"profiles/models"

	"gorm.io/gorm"
)

type ProfilesRepository struct {
	q *gorm.DB
}

// NOTE:
// NewUserRepo() is a constructor function.
// It creates a new instance of UserRepository (in memory),
// stores the given *Db.Queries value inside that struct,
// and returns a pointer to it (*UserRepository).
// Even though we don’t write "UserRepository{...}" in main.go,
// this function does it internally and gives us that ready-to-use object.

func NewProfileRepo(q *gorm.DB) *ProfilesRepository {

	return &ProfilesRepository{q: q}
}

func (r *ProfilesRepository) CreateProfile(arg *models.Profile) error {
	var user models.User
	result := r.q.Where("id=?", arg.UserID).First(&user)
	if result.Error != nil {
		return result.Error
	}
	user.Profile = arg
	result2 := r.q.Save(&user)
	if result2 != nil {
		return result2.Error
	}
	return nil
}

func (r *ProfilesRepository) GetPrfile(email string) (*models.Profile, error) {
	var User models.Profile
	result := r.q.Preload("Profile").Where("email=?", email).First(&User)

	if result.Error != nil || User.Deleted {
		if User.Deleted {
			return &models.Profile{}, fmt.Errorf("user is already deleted")
		}
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &User, nil

}

// func (r *UserRepository) GetAllProfiles(page int, Limit int, sort_by, order, filter int64) ([]models.Profile, error) {

// 	// get the specific rows
// 	offset := (page - 1) * Limit
// 	fmt.Printf("limit : %d and offset : %d", Limit, offset)
// 	rows, err := r.q.Raw("SELECT  * from users Where Deleted=false LIMIT $1 OFFSET $2", Limit, offset).Rows()
// 	if err != nil {
// 		return []models.Profile{}, err
// 	}
// 	defer rows.Close()

// 	var users []models.Profile
// 	for rows.Next() {
// 		var a models.Profile
// 		r.q.ScanRows(rows, &a)
// 		users = append(users, a)

// 	}

// 	// sort_by
// 	sort.Slice(users, func(i, j int) bool {
// 		if order == "asc" {
// 			return users[i].CreatedAt.Before(users[j].CreatedAt)
// 		} else {
// 			return users[i].CreatedAt.After(users[j].CreatedAt)
// 		}
// 	})

// 	//filter
// 	var filtered_data []models.Profile
// 	for _, u := range users {
// 		if strings.Contains(u.ID, filter) {
// 			filtered_data = append(filtered_data, u)
// 		}
// 	}

// 	return filtered_data, nil

// }

func (r *ProfilesRepository) UpdateProfiles(id int64, updates models.Profile) (models.Profile, error) {
	var user2 models.Profile
	result := r.q.Preload("Profile").Where("id=?", id).First(&user2)
	if result.Error != nil || user2.Deleted {
		if user2.Deleted {
			return models.Profile{}, fmt.Errorf("invalid user")
		}
		return models.Profile{}, result.Error
	}
	// prevent changing ID or timestamps
	if updates.CreatedAt.IsZero() {
		return models.Profile{}, fmt.Errorf("restricted field in use")
	}
	if updates.Deleted {
		return models.Profile{}, fmt.Errorf("restricted field in use")
	}
	if updates.ID != 0 {
		return models.Profile{}, fmt.Errorf("restricted field in use")
	}
	if updates.UpdatedAt != nil {
		return models.Profile{}, fmt.Errorf("restricted field in use")
	}

	if err := r.q.Model(&user2).Updates(updates).Error; err != nil {
		return models.Profile{}, err
	}

	if err := r.q.Preload("Profile").Where("id=?", id).First(&user2).Error; err != nil {
		return models.Profile{}, err
	}

	return user2, nil

}

func (r *ProfilesRepository) DeleteProfiles(id int) (models.Profile, error) {
	fmt.Println("enter into the repo")
	var user models.Profile
	result := r.q.Where("id=?", id).First(&user)
	if result.Error != nil || user.Deleted {
		if user.Deleted {
			return models.Profile{}, fmt.Errorf("user already deleted")
		}
		fmt.Println("enter into the  if state")
		return models.Profile{}, result.Error

	}

	fmt.Println("result : ", result)

	user.Deleted = true
	r.q.Save(&user)
	return user, nil
}
