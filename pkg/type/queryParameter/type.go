package queryParameter

import (
	"ol-ilyassov/clean_arch/pkg/type/pagination"
	"ol-ilyassov/clean_arch/pkg/type/sort"
)

type QueryParameter struct {
	Sorts      sort.Sort
	Pagination pagination.Pagination
	// тут можно добавить фильтр.
}
