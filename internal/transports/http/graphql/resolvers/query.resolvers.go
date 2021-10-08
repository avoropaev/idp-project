package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/avoropaev/idp-project/internal/app/code"
	"github.com/avoropaev/idp-project/internal/transports/http/graphql/generated"
)

func (r *queryResolver) Hashcode(ctx context.Context, codeValue int64) (*string, error) {
	return r.CodeService.HashCode(ctx, code.Code(codeValue))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
