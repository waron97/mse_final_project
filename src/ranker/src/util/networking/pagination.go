package networking

import (
	"net/http"
	"strconv"
)

func Pagination(r *http.Request) (int, int) {
	query := r.URL.Query()
	page_ := query.Get("page")
	limit_ := query.Get("limit")
	skip := 0
	limit := 10
	if limit_ != "" {
		limitInt, err := strconv.Atoi(limit_)
		if err != nil {
			limit = 10
		} else {
			limit = limitInt
		}
	}
	if page_ != "" {
		page, err := strconv.Atoi(page_)
		if err != nil {
			page = 1
		}
		skip = (page - 1) * limit
	}
	return skip, limit
}
