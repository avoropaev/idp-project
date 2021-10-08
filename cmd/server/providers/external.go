package providers

import (
	"github.com/avoropaev/idp-project/config"
	"github.com/avoropaev/idp-project/sdk/s1sdk"
	"github.com/avoropaev/idp-project/sdk/s2sdk"
	"net/http"
)

func ProvideS1Client(cfg *config.External, client *http.Client) (s1sdk.S1Client, error) {
	cl, err := s1sdk.NewJSONRPCWithClient(cfg.S1, client)
	if err != nil {
		return nil, err
	}

	return cl, nil
}

func ProvideS2Client(cfg *config.External, client *http.Client) (s2sdk.S2Client, error) {
	cl, err := s2sdk.NewJSONRPCWithClient(cfg.S2, client)
	if err != nil {
		return nil, err
	}

	return cl, nil
}
