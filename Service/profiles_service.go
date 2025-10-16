package service

import (
	"fmt"
	repository "profiles/Repository"
	"profiles/models"
)

type ProfileService struct {
	r *repository.ProfilesRepository
}

// NOTE:
// NewUserRepo() is a constructor function.
// It creates a new instance of UserRepository (in memory),
// stores the given *Db.Queries value inside that struct,
// and returns a pointer to it (*UserRepository).
// Even though we don’t write "UserRepository{...}" in main.go,
// this function does it internally and gives us that ready-to-use object.

func NewProfileService(repo *repository.ProfilesRepository) *ProfileService {
	return &ProfileService{r: repo}
}

// func (s *ProfileService) RegisterUser(email string) (string, error) {

// }

func (s *ProfileService) GetProfileByEmail(email string) (models.Profile, error) {
	result, err := s.r.GetPrfile(email)
	if err != nil {
		return models.Profile{}, err
	}
	return *result, nil
}

// func (s *ProfileService) GetAllProfiles(page int, pageLimit int, sort_by string, order, filter string) ([]models.Profile, error) {

// 	result, err := s.r.GetPrfile(page, pageLimit, sort_by, order, filter)
// 	if err != nil {
// 		return []models.User{}, err
// 	}

// 	return result, nil
// }

func (s *ProfileService) UpdateProfile(id int64, user models.Profile) (models.Profile, error) {
	result, err := s.r.UpdateProfiles(id, user)
	if err != nil {
		return models.Profile{}, err
	}
	return result, nil

}

func (s *ProfileService) DeleteUser(id int) (models.Profile, error) {
	fmt.Println("enter into the service")
	result, err := s.r.DeleteProfiles(id)
	if err != nil {
		return models.Profile{}, err
	}
	return result, nil
}
