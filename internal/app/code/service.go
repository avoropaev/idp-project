package code

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"sync"

	"github.com/avoropaev/idp-project/sdk/s1sdk"
	s1Models "github.com/avoropaev/idp-project/sdk/s1sdk/models"
	"github.com/avoropaev/idp-project/sdk/s2sdk"
	s2Models "github.com/avoropaev/idp-project/sdk/s2sdk/models"
	"github.com/google/uuid"
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
	Rep      Repository
}

// NewService constructor
func NewService(s1Client s1sdk.S1Client, s2Client s2sdk.S2Client, rep Repository) Service {
	return (Service)(&service{
		S1Client: s1Client,
		S2Client: s2Client,
		Rep:      rep,
	})
}

func (s *service) GuidGenerate(_ context.Context, code Code) (res GuidGenerateResponse, err error) {
	if code >= 0 && code <= 9 {
		token := uuid.MustParse("00000000-0000-0000-0000-00000000000" + strconv.FormatInt(int64(code), 10))
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
	var wg sync.WaitGroup
	wg.Add(2)

	res1, err1 := func() (*s1Models.GUIDGenerateResponse, error) {
		defer wg.Done()

		req := s1Models.GUIDGenerateRequest{
			Code: int64(code),
		}

		return s.S1Client.GUIDGenerate(ctx, req)
	}()

	res2, err2 := func() (*s2Models.HashCalcResponse, error) {
		defer wg.Done()

		req := s2Models.HashCalcRequest{
			Code: int64(code),
		}

		return s.S2Client.HashCalc(ctx, req)
	}()

	wg.Wait()

	if err1 != nil {
		return nil, err1
	}

	if err2 != nil {
		return nil, err2
	}

	if res1 == nil {
		return nil, ErrTokenNotFound
	}

	if res2 == nil {
		return nil, ErrHashNotFound
	}

	token, err := uuid.Parse(res1.Token)
	if err != nil {
		return nil, err
	}

	result, err := s.Rep.GetDataByTokenAndHash(ctx, token, res2.Hash)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
