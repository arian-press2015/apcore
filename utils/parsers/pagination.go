package parsers

import "strconv"

const (
	DEFAULT_OFFSET = 0
	DEFAULT_LIMIT  = 10
)

func ParsePaginationParams(offsetStr, limitStr string) (int, int) {
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = DEFAULT_OFFSET
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = DEFAULT_LIMIT
	}

	return offset, limit
}
