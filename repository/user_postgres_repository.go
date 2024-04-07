package repository

import (
	"context"

	"golang-clean-architecture/target/ent"
)

type UserPostgresRepository struct {
	db *ent.Client
}

func NewUserPostgresRepository(db *ent.Client) UserRepository {
	return &UserPostgresRepository{db: db}
}

func (repository *UserPostgresRepository) ListUsers(ctx context.Context) ([]*ent.User, error) {
	return repository.db.User.Query().All(ctx)
}
