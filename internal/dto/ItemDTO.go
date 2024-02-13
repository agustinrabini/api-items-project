package dto

import (
	"github.com/agustinrabini/api-items-project/internal/domain"
	"github.com/mitchellh/mapstructure"
)

type ItemDTO struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	UserID      string        `json:"user_id"`
	Status      string        `json:"status" `
	Category    CategoryDTO   `json:"category"`
	Price       PriceDTO      `json:"price"`
	Images      []ImageDTO    `json:"images"`
	Attributes  AttributesDTO `json:"attributes"`
	Eligible    []EligibleDTO `json:"eligible"`
}

type ImageDTO string
type AttributesDTO map[string]string

func (i ItemDTO) ToItem() (domain.Item, error) {
	item := domain.Item{}
	err := mapstructure.Decode(i, &item)
	if err != nil {
		return domain.Item{}, err
	}
	return item, nil
}
