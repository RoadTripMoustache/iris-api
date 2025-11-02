package utils

import (
	"io"
	"net/http"
)

type ContextKey string

var (
	ContextEmailKey         ContextKey = "EMAIL"
	ContextTokenKey         ContextKey = "TOKEN"
	ContextPaginationKey    ContextKey = "PAGINATION"
	ContextVersionFilterKey ContextKey = "VERSION_FILTER"
)

type Context struct {
	Pagination    Pagination
	VersionFilter VersionFilter
	UserID        string
	UserEmail     string
	QueryParams   map[string][]string
	Headers       map[string]string
	Vars          map[string]string
	Body          io.ReadCloser
	Path          string
	Method        string
	RequestURI    string
	Request       *http.Request
}

func (c Context) CleanPagination() Context {
	c.Pagination = Pagination{}

	return c
}

func (c Context) CleanVersionFilter() Context {
	c.VersionFilter = VersionFilter{}

	return c
}
