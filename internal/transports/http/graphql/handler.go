//Package graphql contains graphql app endpoints
package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/avoropaev/idp-project/internal/transports/http/graphql/generated"
	"github.com/avoropaev/idp-project/internal/transports/http/graphql/resolvers"
	"net/http"
)

//go:generate go run github.com/99designs/gqlgen

func NewGraphqlHandler() *handler.Server {
	return handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolvers.Resolver{},
	}))
}

func NewPlaygroundHandler() http.Handler {
	return playground.Handler("IDP playground", "/graphql")
}
