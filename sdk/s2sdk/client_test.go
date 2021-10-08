package s2sdk_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/avoropaev/idp-project/sdk/s2sdk"
	"github.com/avoropaev/idp-project/sdk/s2sdk/models"

	"github.com/stretchr/testify/suite"
)

type s2TestSuite struct {
	suite.Suite
	s2     *httptest.Server
	client s2sdk.S2Client
	ctx    context.Context
}

var (
	fixtures = struct {
		responseHashCalc *models.HashCalcResponse
	}{
		responseHashCalc: &models.HashCalcResponse{
			Hash: "test-hash",
		},
	}
)

func (ex *s2TestSuite) SetupSuite() {
	ex.ctx = context.Background()

	ex.s2 = httptest.NewServer(http.HandlerFunc(s2Handler))

	var err error

	ex.client, err = s2sdk.NewJSONRPC(ex.s2.URL)
	ex.Require().NoError(err)
}

func s2Handler(w http.ResponseWriter, r *http.Request) {
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
	case "hash.calc":
		result := struct {
			JSONRPC string `json:"jsonrpc"`
			ID      int64  `json:"id"`
			Result  models.HashCalcResponse
		}{
			JSONRPC: "2.0",
			ID:      time.Now().Unix(),
			Result:  *fixtures.responseHashCalc,
		}

		b, _ = json.Marshal(result)
		_, _ = fmt.Fprintf(w, "%s", b)
	}
}

func (ex *s2TestSuite) TearDownSuite() {
	ex.s2.Close()
}

func (ex *s2TestSuite) TestClient_HashCalc() {
	resp, err := ex.client.HashCalc(context.Background(), models.HashCalcRequest{})

	ex.Require().NoError(err)
	ex.Require().Equal(resp, fixtures.responseHashCalc)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(s2TestSuite))
}
