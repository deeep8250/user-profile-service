package models

import (
	"time"
)

type Profile struct {
	ID        *int64     `json:"id"`
	UserID    int64      `json:"user_id"`
	AvatarURL string     `json:"avatar_url"`
	Bio       string     `json:"bio"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Deleted   bool       `json:"deleted"`
}

type User struct {
	ID        *int64     `json:"id"`
	Email     string     `json:"email"`
	Name      string     `json:"name"`
	Password  string     `json:"-"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Deleted   bool       `json:"deleted"`
	Profile   *Profile   `gorm:"foreignKey:UserID;references:ID" json:"profile"` ///omitempty Prevents sending "profile": null if it’s missing.
}
