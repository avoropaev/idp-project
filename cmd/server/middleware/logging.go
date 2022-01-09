package middleware

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-kit/kit/endpoint"
	appKitEndpoint "github.com/sagikazarmark/appkit/endpoint"
	kitXEndpoint "github.com/sagikazarmark/kitx/endpoint"
)

func LoggingMiddleware(logger appKitEndpoint.Logger) endpoint.Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.TraceContext(ctx, "processing request")

			defer func(begin time.Time) {
				operationName, _ := kitXEndpoint.OperationName(ctx)
				body, _ := json.Marshal(request)

				logger.TraceContext(ctx, "processing request finished", map[string]interface{}{
					"elapsed_time":   time.Since(begin),
					"operation_name": operationName,
					"body":           body,
				})
			}(time.Now())

			return e(ctx, request)
		}
	}
}
