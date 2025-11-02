package middlewares

import (
	"context"
	routerUtils "github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/utils"
	"net/http"
)

// PaginationMiddleware - Middleware to generate the Pagination object based on query param values
func PaginationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pagination := routerUtils.Pagination{
			PageNumber: utils.StrToInt(routerUtils.QueryParamExtractor(r.URL.Query(), "page[number]")),
			PageSize:   utils.StrToInt(routerUtils.QueryParamExtractor(r.URL.Query(), "page[size]")),
		}

		ctx := context.WithValue(r.Context(), routerUtils.ContextPaginationKey, pagination)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
