package s2sdk

import (
	"net/http"
	"net/url"

	io "github.com/avoropaev/idp-project/sdk/s2sdk/io/jsonrpc"

	"github.com/go-kit/kit/transport/http/jsonrpc"
)

type jsonrpcclient struct {
	httpClient *http.Client
	client
}

func NewJSONRPC(url string) (S2Client, error) {
	return NewJSONRPCWithClient(url, http.DefaultClient)
}

func NewJSONRPCWithClient(serviceURL string, httpClient *http.Client) (S2Client, error) {
	c := &jsonrpcclient{httpClient: httpClient}

	parsedURL, err := url.Parse(serviceURL + "/rpc")
	if err != nil {
		return nil, err
	}

	c.getJSONRPCEndpoints(parsedURL)

	return c, nil
}

func (c *jsonrpcclient) getJSONRPCEndpoints(url *url.URL) {
	c.endpoints.HashCalc = jsonrpc.NewClient(
		url,
		"hash.calc",
		jsonrpc.SetClient(c.httpClient),
		jsonrpc.ClientResponseDecoder(io.DecodeHashCalc),
	).Endpoint()
}
