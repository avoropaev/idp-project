package main

import (
	"context"

	"github.com/avoropaev/idp-project/cmd/server/middleware"
	"github.com/avoropaev/idp-project/cmd/server/providers"
	"github.com/avoropaev/idp-project/internal/app/code/codedriver"
	appHttp "github.com/avoropaev/idp-project/internal/transports/http/jsonrpc"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/tracing/opencensus"
	"github.com/go-kit/kit/transport/http/jsonrpc"
	"github.com/gorilla/mux"
	"github.com/sagikazarmark/kitx/correlation"
	"github.com/streadway/handy/encoding"
	"github.com/vseinstrumentiru/lego/v2/log"
	legoHttp "github.com/vseinstrumentiru/lego/v2/transport/http"
	legoMiddleware "github.com/vseinstrumentiru/lego/v2/transport/http/middleware"

	"net/http"

	codeModule "github.com/avoropaev/idp-project/internal/app/code"
	"github.com/avoropaev/idp-project/internal/transports/http/graphql"
	kitXEndpoint "github.com/sagikazarmark/kitx/endpoint"
)

type app struct {
	HTTP *legoHttp.Config
}

func (app app) Providers() []interface{} {
	return []interface{}{
		providers.ProvideCodeRepository,
		providers.ProvideCodeService,
		providers.ProvideS1Client,
		providers.ProvideS2Client,
	}
}

func (app app) ConfigureHTTP(router *mux.Router, cmService *codeModule.Service, logger log.Logger) {
	router.StrictSlash(true)

	endpointMiddleware := []endpoint.Middleware{
		correlation.Middleware(),
		opencensus.TraceEndpoint("", opencensus.WithSpanName(func(ctx context.Context, _ string) string {
			name, _ := kitXEndpoint.OperationName(ctx)

			return name
		})),
		middleware.LoggingMiddleware(logger),
	}

	mw := kitXEndpoint.Combine(endpointMiddleware...)

	gcEndpoints := codedriver.MakeEndpoints(*cmService, mw)

	handler := appHttp.MakeHandlers(gcEndpoints)
	jsonRPCHandler := jsonrpc.NewServer(handler)

	router.Path("/rpc").
		Methods(http.MethodPost).
		Handler(jsonRPCHandler)

	router.Path("/graphql").
		Handler(graphql.NewGraphqlHandler(*cmService))

	router.Path("/graphql/playground").
		Handler(graphql.NewPlaygroundHandler())

	router.Use(
		encoding.Gzipper(5),
		legoMiddleware.LogRequestWithMaxLenMiddleware(5*1024),
		legoMiddleware.LogResponseWithMaxLenMiddleware(5*1024),
		legoMiddleware.TraceTagsMiddleware(legoMiddleware.TraceTagsMiddlewareConfig{
			"x-trace-request-dc":  "request.dc",
			"x-trace-request-app": "request.app",
		}),
	)
}
