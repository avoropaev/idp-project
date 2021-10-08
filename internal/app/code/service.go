package code

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/google/uuid"
	"strconv"
)

var (
	TokenNotFound = errors.New("token not found")
)

// +kit:endpoint

// Service contract
type Service interface {
	GuidGenerate(ctx context.Context, code Code) (Res GuidGenerateResponse, err error)
	HashCalc(ctx context.Context, code Code) (Res HashCalcResponse, err error)
}

// service realization
type service struct{}

// NewService constructor
func NewService() Service {
	return (Service)(&service{})
}

func (s *service) GuidGenerate(_ context.Context, code Code) (res GuidGenerateResponse, err error) {
	if code == 1 {
		token := uuid.MustParse("00000000-0000-0000-0000-000000000001")
		res.Token = &token
	}

	if res.Token == nil {
		return res, TokenNotFound
	}

	return res, nil
}

func (s *service) HashCalc(_ context.Context, code Code) (res HashCalcResponse, err error) {
	data := strconv.Itoa(int(code))
	h := sha256.New()
	h.Write([]byte(data))
	hash := hex.EncodeToString(h.Sum([]byte("secret")))

	res.Hash = &hash

	return res, nil
}
