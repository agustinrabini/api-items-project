package dto

import "github.com/agustinrabini/api-items-project/internal/domain"

type CategoryDTO struct {
	ID   string `json:"id", omitempty`
	Name string `json:"name"`
}

type CategoriesDTO struct {
	CategoryDTO []domain.Category `json:"categories"`
}
