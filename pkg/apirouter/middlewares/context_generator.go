// Package middlewares contains all the middleware functions.
package middlewares

import (
	routerUtils "github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func GenerateContext(r *http.Request, method string, path string) routerUtils.Context {
	headers := make(map[string]string)
	for key, values := range r.Header {
		convertedValue := strings.Join(values, ",") // Join values with comma (modify separator as needed)
		headers[key] = convertedValue
	}

	vars := mux.Vars(r)

	pagination := routerUtils.Pagination{}
	if value := r.Context().Value(routerUtils.ContextPaginationKey); value != nil {
		pagination = value.(routerUtils.Pagination)
	}

	versionFilter := routerUtils.VersionFilter{}
	if value := r.Context().Value(routerUtils.ContextVersionFilterKey); value != nil {
		versionFilter = value.(routerUtils.VersionFilter)
	}

	userID := ""
	if value := r.Context().Value(routerUtils.ContextTokenKey); value != nil {
		userID = value.(string)
	}

	userEmail := ""
	if value := r.Context().Value(routerUtils.ContextEmailKey); value != nil {
		userEmail = value.(string)
	}

	return routerUtils.Context{
		Pagination:    pagination,
		VersionFilter: versionFilter,
		UserID:        userID,
		UserEmail:     userEmail,
		QueryParams:   r.URL.Query(),
		Headers:       headers,
		Vars:          vars,
		Body:          r.Body,
		Method:        method,
		RequestURI:    r.RequestURI,
		Path:          path,
		Request:       r,
	}
}
