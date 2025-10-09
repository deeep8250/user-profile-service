package repository

import (
	"context"
	Db "profiles/internal/db"
)

type UserRepository struct {
	q *Db.Queries
}

// NOTE:
// NewUserRepo() is a constructor function.
// It creates a new instance of UserRepository (in memory),
// stores the given *Db.Queries value inside that struct,
// and returns a pointer to it (*UserRepository).
// Even though we don’t write "UserRepository{...}" in main.go,
// this function does it internally and gives us that ready-to-use object.

func NewUserRepo(q *Db.Queries) *UserRepository {

	return &UserRepository{q: q}
}

func (r *UserRepository) CreateUser(ctx context.Context, arg Db.CreateUserParams) (Db.User, error) {
	return r.q.CreateUser(ctx, arg)
}

func (r *UserRepository) GetUser(ctx context.Context, id int64) (Db.User, error) {
	return r.q.GetUserByIdNew(ctx, id)
}
