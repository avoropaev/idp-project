package jsonrpc

import (
	"context"
	"encoding/json"

	"github.com/avoropaev/idp-project/sdk/s2sdk/models"

	"github.com/go-kit/kit/transport/http/jsonrpc"
)

func DecodeHashCalc(_ context.Context, response jsonrpc.Response) (interface{}, error) {
	if response.Error != nil {
		return nil, response.Error
	}

	var res models.HashCalcResponse
	if err := json.Unmarshal(response.Result, &res); err != nil {
		return nil, err
	}

	return res, nil
}
