// Package s1sdk is client for request to courier service
package s1sdk

import (
	"context"
	"fmt"

	"github.com/avoropaev/idp-project/sdk/s1sdk/models"

	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
)

var ErrCast = errors.New("cast error")

type S1Client interface {
	GUIDGenerate(context.Context, models.GUIDGenerateRequest) (*models.GUIDGenerateResponse, error)
}

type client struct {
	endpoints endpoints
}

type endpoints struct {
	GUIDGenerate endpoint.Endpoint
}

func (c *client) GUIDGenerate(ctx context.Context, req models.GUIDGenerateRequest) (*models.GUIDGenerateResponse, error) {
	res, err := c.endpoints.GUIDGenerate(ctx, req)

	if err != nil {
		return nil, err
	}

	resp, ok := res.(models.GUIDGenerateResponse)
	if !ok {
		return nil, fmt.Errorf("%w to %T", ErrCast, models.GUIDGenerateResponse{})
	}

	return &resp, nil
}
