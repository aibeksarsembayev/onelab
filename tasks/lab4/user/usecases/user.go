package usecase

import (
	"context"
	"time"

	"github.com/aibeksarsembayev/onelab/tasks/lab4/domain"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

// NewUserUsecase will create new userUsecase object representation of domain.UserUsecase interface
func NewUserUsecase(u domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepo:       u,
		contextTimeout: timeout,
	}
}

func (u *userUsecase) Create(ctx context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return nil
}

func (u *userUsecase) GetByID(ctx context.Context, id int) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return domain.User{}, nil
}

func (u *userUsecase) GetAll(ctx context.Context) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return nil, nil
}

func (u *userUsecase) Update(ctx context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return nil
}

func (u *userUsecase) Delete(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return nil
}
