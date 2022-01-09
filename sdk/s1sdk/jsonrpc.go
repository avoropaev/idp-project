package s1sdk

import (
	"net/http"
	"net/url"

	io "github.com/avoropaev/idp-project/sdk/s1sdk/io/jsonrpc"

	"github.com/go-kit/kit/transport/http/jsonrpc"
)

type jsonrpcclient struct {
	httpClient *http.Client
	client
}

func NewJSONRPC(url string) (S1Client, error) {
	return NewJSONRPCWithClient(url, http.DefaultClient)
}

func NewJSONRPCWithClient(serviceURL string, httpClient *http.Client) (S1Client, error) {
	c := &jsonrpcclient{httpClient: httpClient}

	parsedURL, err := url.Parse(serviceURL + "/rpc")
	if err != nil {
		return nil, err
	}

	c.getJSONRPCEndpoints(parsedURL)

	return c, nil
}

func (c *jsonrpcclient) getJSONRPCEndpoints(url *url.URL) {
	c.endpoints.GUIDGenerate = jsonrpc.NewClient(
		url,
		"guid.generate",
		jsonrpc.SetClient(c.httpClient),
		jsonrpc.ClientResponseDecoder(io.DecodeGuidGenerate),
	).Endpoint()
}
