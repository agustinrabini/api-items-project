package items

import (
	"github.com/agustinrabini/api-items-project/internal/domain"
)

func pricesResponseToItemPrice(pricesResponse domain.Prices, items []domain.Item) []domain.Item {

	buildedItems := []domain.Item{}

	for _, pr := range pricesResponse.Prices {

		for _, item := range items {

			if pr.ItemId == item.ID {
				item.Price = pr

				buildedItems = append(buildedItems, item)

				break
			}
		}
	}

	return buildedItems
}
