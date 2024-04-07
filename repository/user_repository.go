package repository

import (
	"context"

	"golang-clean-architecture/target/ent"
)

type (
	UserRepository interface {
		ListUsers(ctx context.Context) ([]*ent.User, error)
	}
)
