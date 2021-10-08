package s1sdk_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/avoropaev/idp-project/sdk/s1sdk"
	"github.com/avoropaev/idp-project/sdk/s1sdk/models"

	"github.com/stretchr/testify/suite"
)

type s1TestSuite struct {
	suite.Suite
	s1     *httptest.Server
	client s1sdk.S1Client
	ctx    context.Context
}

var (
	fixtures = struct {
		responseGuidGenerate *models.GuidGenerateResponse
	}{
		responseGuidGenerate: &models.GuidGenerateResponse{
			Token: "test-token",
		},
	}
)

func (ex *s1TestSuite) SetupSuite() {
	ex.ctx = context.Background()

	ex.s1 = httptest.NewServer(http.HandlerFunc(s1Handler))

	var err error

	ex.client, err = s1sdk.NewJSONRPC(ex.s1.URL)
	ex.Require().NoError(err)
}

func s1Handler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	in := struct {
		JSONRPC string      `json:"jsonrpc"`
		ID      int64       `json:"id"`
		Method  string      `json:"method"`
		Params  interface{} `json:"params"`
	}{}

	err = json.Unmarshal(b, &in)
	if err != nil {
		return
	}

	switch in.Method {
	case "guid.generate":
		result := struct {
			JSONRPC string `json:"jsonrpc"`
			ID      int64  `json:"id"`
			Result  models.GuidGenerateResponse
		}{
			JSONRPC: "2.0",
			ID:      time.Now().Unix(),
			Result:  *fixtures.responseGuidGenerate,
		}

		b, _ = json.Marshal(result)
		_, _ = fmt.Fprintf(w, "%s", b)
	}
}

func (ex *s1TestSuite) TearDownSuite() {
	ex.s1.Close()
}

func (ex *s1TestSuite) TestClient_GuidGenerate() {
	resp, err := ex.client.GuidGenerate(context.Background(), models.GuidGenerateRequest{})

	ex.Require().NoError(err)
	ex.Require().Equal(resp, fixtures.responseGuidGenerate)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(s1TestSuite))
}
