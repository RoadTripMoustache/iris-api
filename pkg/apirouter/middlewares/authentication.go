// Package middlewares contains all the middleware functions.
package middlewares

import (
	"context"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/apirouter/utils"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/enum"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/tools/auth"
	"net/http"
	"strings"
)

// AuthenticationMiddleware - Validate authentication of the requester
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/v1/admob") {
			next.ServeHTTP(w, r)
			return
		}
		authorizationHeader := r.Header.Get(enum.Authorization)

		if len(authorizationHeader) < 8 {
			// If the token is too short, it means it's an incorrect one. So we return a 401 error.
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		authorizationHeader = authorizationHeader[7:]

		// Validate firebase token
		token := auth.GetInstance().VerifyIDToken(authorizationHeader)
		if token == nil {
			// If the token is nil, it means it's an incorrect one. So we return a 401 error.
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), utils.ContextTokenKey, token.Subject)
		ctx = context.WithValue(ctx, utils.ContextEmailKey, token.Firebase.Identities["email"].([]interface{})[0])
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
