package code

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/avoropaev/idp-project/sdk/s1sdk"
	s1Models "github.com/avoropaev/idp-project/sdk/s1sdk/models"
	"github.com/avoropaev/idp-project/sdk/s2sdk"
	s2Models "github.com/avoropaev/idp-project/sdk/s2sdk/models"
	"github.com/google/uuid"
	"strconv"
)

var (
	ErrTokenNotFound = errors.New("token not found")
	ErrHashNotFound  = errors.New("hash not found")
)

// +kit:endpoint

// Service contract
type Service interface {
	GuidGenerate(ctx context.Context, code Code) (res GuidGenerateResponse, err error)
	HashCalc(ctx context.Context, code Code) (res HashCalcResponse, err error)
	HashCode(ctx context.Context, code Code) (*string, error)
}

// service realization
type service struct {
	S1Client s1sdk.S1Client
	S2Client s2sdk.S2Client
}

// NewService constructor
func NewService(s1Client s1sdk.S1Client, s2Client s2sdk.S2Client) Service {
	return (Service)(&service{
		S1Client: s1Client,
		S2Client: s2Client,
	})
}

func (s *service) GuidGenerate(_ context.Context, code Code) (res GuidGenerateResponse, err error) {
	if code == 1 {
		token := uuid.MustParse("00000000-0000-0000-0000-000000000001")
		res.Token = &token
	}

	if res.Token == nil {
		return res, ErrTokenNotFound
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

func (s *service) HashCode(ctx context.Context, code Code) (*string, error) {
	req1 := s1Models.GuidGenerateRequest{
		Code: int64(code),
	}

	res1, err := s.S1Client.GuidGenerate(ctx, req1)
	if err != nil {
		return nil, err
	}

	req2 := s2Models.HashCalcRequest{
		Code: int64(code),
	}

	res2, err := s.S2Client.HashCalc(ctx, req2)
	if err != nil {
		return nil, err
	}

	if res1 == nil {
		return nil, ErrTokenNotFound
	}

	if res2 == nil {
		return nil, ErrHashNotFound
	}

	result := res1.Token + "_" + res2.Hash

	return &result, nil
}
