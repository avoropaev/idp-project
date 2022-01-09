package code

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

// Repository contract
type Repository interface {
	GetDataByTokenAndHash(ctx context.Context, token uuid.UUID, hash string) (res string, err error)
}

// repository realization
type repository struct {
	db *sql.DB
}

// NewRepository constructor
func NewRepository(db *sql.DB) Repository {
	return (Repository)(&repository{db})
}

func (s *repository) GetDataByTokenAndHash(ctx context.Context, token uuid.UUID, hash string) (res string, err error) {
	data := new(string)

	err = s.db.QueryRowContext(
		ctx,
		"SELECT data FROM hdata WHERE guid::text = $1 and hash = $2;",
		token.String(),
		hash,
	).Scan(data)

	if err != nil {
		return "", err
	}

	return *data, nil
}
