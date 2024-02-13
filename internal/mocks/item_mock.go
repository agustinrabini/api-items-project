package mocks

import "github.com/agustinrabini/api-items-project/internal/domain"

var ItemMock = domain.Item{
	ID:     "123",
	Name:   "Example Item",
	ShopID: "456",
	Price: domain.Price{
		Currency: domain.Currency{
			ID: "USD",
		},
		Amount: 19.99,
	},
	Description: "This is an example item for testing purposes.",
	Status:      "active",
	Images: []domain.Image{
		"https://example.com/image1.jpg",
		"https://example.com/image2.jpg",
	},
	Attributes: domain.Attributes{
		"color": "blue",
		"size":  "M",
		"brand": "Example Brand",
	},
	Eligible: []domain.Eligible{
		{
			ID:         "1",
			Title:      "Example Elegible 1",
			Type:       "checkbox",
			IsRequired: true,
			Options: []domain.Option{
				"Option 1",
				"Option 2",
				"Option 3",
			},
		},
		{
			ID:         "2",
			Title:      "Example Elegible 2",
			Type:       "select",
			IsRequired: false,
			Options: []domain.Option{
				"Option A",
				"Option B",
				"Option C",
			},
		},
	},
}
