package domain

import (
	"context"
	"time"
)

type User struct {
	ID        int       `json:"id`
	Name      string    `json:"name"`
	Surname   string    `json:"name"`
	Email     string    `json:"email"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// UserUsecases ...
type UserUsecase interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id int) (User, error)
	GetAll(ctx context.Context) ([]User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error
}

// UserRepository ...
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id int) (User, error)
	GetAll(ctx context.Context) ([]User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error
}
