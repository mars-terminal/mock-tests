package service

import (
	"context"

	"github.com/mars-terminal/mock-tests/entities/user"
)

//go:generate mockgen -source=user.go -package=service -destination=user_mock.go

type UserService interface {
	Create(ctx context.Context, name, password string) (*user.User, error)
}
