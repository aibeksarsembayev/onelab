package db

import (
	"context"

	"github.com/aibeksarsembayev/onelab/tasks/lab4/domain"
)

type dbUserRepository struct {
	// Conn
}

// NewDBUserRepository ...
func NewDBUserRepository() domain.UserRepository {
	return &dbUserRepository{}
}

func (db *dbUserRepository) Create(ctx context.Context, user *domain.User) error {
	return nil
}

func (db *dbUserRepository) GetByID(ctx context.Context, id int) (domain.User, error) {
	return domain.User{}, nil
}

func (db *dbUserRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	return nil, nil
}

func (db *dbUserRepository) Update(ctx context.Context, user *domain.User) error {
	return nil
}

func (db *dbUserRepository) Delete(ctx context.Context, id int) error {
	return nil
}
