//Package graphql contains graphql app endpoints
package graphql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	codeModule "github.com/avoropaev/idp-project/internal/app/code"
	"github.com/avoropaev/idp-project/internal/transports/http/graphql/generated"
	"github.com/avoropaev/idp-project/internal/transports/http/graphql/resolvers"
)

//go:generate go run github.com/99designs/gqlgen

func NewGraphqlHandler(codeService codeModule.Service) *handler.Server {
	return handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolvers.Resolver{
			CodeService: codeService,
		},
	}))
}

func NewPlaygroundHandler() http.Handler {
	return playground.Handler("IDP playground", "/graphql")
}
