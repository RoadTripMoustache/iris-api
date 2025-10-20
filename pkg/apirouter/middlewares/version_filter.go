package middlewares

import (
	"context"
	routerUtils "github.com/RoadTripMoustache/guide_nestor_api/pkg/apirouter/utils"
	"net/http"
)

// VersionFilterMiddleware - Middleware to generate the VersionFilter object based on query param values
func VersionFilterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		versionFilter := routerUtils.VersionFilter{
			CurrentVersion: routerUtils.QueryParamExtractor(r.URL.Query(), "version[current]"),
		}
		ctx := context.WithValue(r.Context(), routerUtils.ContextVersionFilterKey, versionFilter)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
