package models

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ID        int64     `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	AvatarURL string    `json:"avatar_url"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDelete  bool      `json:"is_delete"`
}

type User struct {
	ID        int64
	Email     string
	Name      string
	Password  string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDelete  bool      `json:"is_delete"`
	Profile   *Profile  `json:"profile,omitempty"` ///omitempty Prevents sending "profile": null if it’s missing.
}
