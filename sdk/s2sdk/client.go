// Package s2sdk is client for request to courier service
package s2sdk

import (
	"context"
	"fmt"

	"github.com/avoropaev/idp-project/sdk/s2sdk/models"

	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
)

var ErrCast = errors.New("cast error")

type S2Client interface {
	HashCalc(context.Context, models.HashCalcRequest) (*models.HashCalcResponse, error)
}

type client struct {
	endpoints endpoints
}

type endpoints struct {
	HashCalc endpoint.Endpoint
}

func (c *client) HashCalc(ctx context.Context, req models.HashCalcRequest) (*models.HashCalcResponse, error) {
	res, err := c.endpoints.HashCalc(ctx, req)

	if err != nil {
		return nil, err
	}

	resp, ok := res.(models.HashCalcResponse)
	if !ok {
		return nil, fmt.Errorf("%w to %T", ErrCast, models.HashCalcResponse{})
	}

	return &resp, nil
}
