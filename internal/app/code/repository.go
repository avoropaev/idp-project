package code

import (
	"context"
	"github.com/google/uuid"
)

// Repository contract
type Repository interface {
	GetDataByTokenAndHash(ctx context.Context, token uuid.UUID, hash string) (res string, err error)
}

// repository realization
type repository struct {
}

// NewRepository constructor
func NewRepository() Repository {
	return (Repository)(&repository{})
}

func (s *repository) GetDataByTokenAndHash(_ context.Context, token uuid.UUID, hash string) (res string, err error) {
	return token.String() + "_" + hash, nil
}
