package service

import (
	"context"

	"github.com/mars-terminal/mock-tests/entities/user"
)

type UserService interface {
	Create(ctx context.Context, name, password string) (*user.User, error)
}
