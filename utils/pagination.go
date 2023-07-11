package utils

type Pagination struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
}

func GetOffset(page int, size int) int {

	if page == 0 || size == 0 {
		page, size = -1, -1
	}

	if page <= 0 {
		return 0
	}
	return (page - 1) * size
}
