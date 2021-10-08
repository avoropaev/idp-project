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
	GuidGenerate(context.Context, models.GuidGenerateRequest) (*models.GuidGenerateResponse, error)
}

type client struct {
	endpoints endpoints
}

type endpoints struct {
	GuidGenerate endpoint.Endpoint
}

func (c *client) GuidGenerate(ctx context.Context, req models.GuidGenerateRequest) (*models.GuidGenerateResponse, error) {
	res, err := c.endpoints.GuidGenerate(ctx, req)

	if err != nil {
		return nil, err
	}

	resp, ok := res.(models.GuidGenerateResponse)
	if !ok {
		return nil, fmt.Errorf("%w to %T", ErrCast, models.GuidGenerateResponse{})
	}

	return &resp, nil
}
