package storage

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/mars-terminal/mock-tests/entities/user"
)

type Store struct {
	db    sqlx.DB
	table string
}

func (s *Store) Create(ctx context.Context, name, password string) (*user.User, error) {
	query := fmt.Sprintf(`
INSERT 
	INTO %s
VALUES 
	(name=$1, password=$2)
RETURNING id
`, s.table)

	result, err := s.db.ExecContext(ctx, query, name, password)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &user.User{
		Id:       id,
		Name:     name,
		Password: password,
	}, nil

}
