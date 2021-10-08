package jsonrpc

import (
	"context"
	"encoding/json"
	"github.com/avoropaev/idp-project/internal/app/code/codedriver"
)

func guidGenerateRequestDecoder(_ context.Context, msg json.RawMessage) (interface{}, error) {
	var req codedriver.GuidGenerateRequest

	if err := json.Unmarshal(msg, &req); err != nil {
		return nil, err
	}

	return req, nil
}

func guidGenerateResponseEncoder(_ context.Context, msg interface{}) (response json.RawMessage, err error) {
	data := msg.(codedriver.GuidGenerateResponse)

	if err = data.Failed(); err != nil {
		return nil, err
	}

	return json.Marshal(data.Res)
}

func hashCalcRequestDecoder(_ context.Context, msg json.RawMessage) (interface{}, error) {
	var req codedriver.HashCalcRequest

	if err := json.Unmarshal(msg, &req); err != nil {
		return nil, err
	}

	return req, nil
}

func hashCalcResponseEncoder(_ context.Context, msg interface{}) (response json.RawMessage, err error) {
	data := msg.(codedriver.HashCalcResponse)

	if err = data.Failed(); err != nil {
		return nil, err
	}

	return json.Marshal(data.Res)
}
