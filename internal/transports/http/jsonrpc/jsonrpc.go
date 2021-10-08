package jsonrpc

import (
	"github.com/avoropaev/idp-project/internal/app/code/codedriver"
	"github.com/go-kit/kit/transport/http/jsonrpc"
)

func MakeHandlers(gc codedriver.Endpoints) jsonrpc.EndpointCodecMap {
	return jsonrpc.EndpointCodecMap{
		"guid.generate": {
			Endpoint: gc.GuidGenerate,
			Decode:   guidGenerateRequestDecoder,
			Encode:   guidGenerateResponseEncoder,
		},
		"hash.calc": {
			Endpoint: gc.HashCalc,
			Decode:   hashCalcRequestDecoder,
			Encode:   hashCalcResponseEncoder,
		},
	}
}
