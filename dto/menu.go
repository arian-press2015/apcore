package dto

type CreateMenuBody struct {
	Categories []string `json:"categories" binding:"required"`
}

type UpdateMenuBody struct {
	Categories []string `json:"categories" binding:"required"`
}
