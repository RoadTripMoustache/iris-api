// Package utils contains all utils methods for middlewares.
package utils

type Pagination struct {
	PageNumber *int
	PageSize   *int
}

// GetOffset - Calculate the offset to use based on the page number and size.
func (p Pagination) GetOffset() *int {
	if p.PageNumber == nil || p.PageSize == nil {
		return nil
	}
	offset := (*p.PageNumber - 1) * *p.PageSize
	return &offset
}
