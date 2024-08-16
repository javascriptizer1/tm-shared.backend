package pagination

import (
	"errors"
)

type PaginationParams struct {
	Limit  int64
	Offset int64
}

func New(limit *int64, page *int64) (PaginationParams, error) {
	var zero PaginationParams
	var pagination PaginationParams

	if limit != nil && *limit > 0 {
		pagination.Limit = *limit
	} else {
		pagination.Limit = 10
	}

	if page != nil && *page > 0 {
		pagination.Offset = (*page - 1) * pagination.Limit
	} else {
		pagination.Offset = 0
	}

	if pagination.Limit > 1000 {
		return zero, errors.New("limit too large")
	}

	return pagination, nil
}
