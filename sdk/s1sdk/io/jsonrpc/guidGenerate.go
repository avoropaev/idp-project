package jsonrpc

import (
	"context"
	"encoding/json"

	"github.com/avoropaev/idp-project/sdk/s1sdk/models"

	"github.com/go-kit/kit/transport/http/jsonrpc"
)

func DecodeGuidGenerate(_ context.Context, response jsonrpc.Response) (interface{}, error) {
	if response.Error != nil {
		return nil, response.Error
	}

	var res models.GUIDGenerateResponse
	if err := json.Unmarshal(response.Result, &res); err != nil {
		return nil, err
	}

	return res, nil
}
