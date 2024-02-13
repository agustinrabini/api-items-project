package dependencies

import (
	"github.com/agustinrabini/api-items-project/cmd/api/handlers"
	"github.com/agustinrabini/api-items-project/internal/categories"
	"github.com/agustinrabini/api-items-project/internal/clients/prices"
	"github.com/agustinrabini/api-items-project/internal/clients/shops"
	"github.com/agustinrabini/api-items-project/internal/items"
)

type Dependencies interface {
	ItemsRepository() items.Repository
	CategoriesRepository() categories.Repository
}

func GetDependecyManager() Dependencies {
	return NewDependencyManager()
}

func BuildDependencies() (HandlersStruct, error) {
	depManager := GetDependecyManager()

	// Repository
	itemsRepository := depManager.ItemsRepository()
	categoriesRepository := depManager.CategoriesRepository()

	// External Clients
	pricesClient := prices.ClientInstance
	shopsClient := shops.ClientInstance

	// Services
	itemsService := items.NewService(itemsRepository, pricesClient, shopsClient)
	categoriesService := categories.NewService(categoriesRepository)

	// Handlers
	itemsHandler := handlers.NewItemsHandler(itemsService)
	categoriesHandler := handlers.NewCategoriesHandler(categoriesService)

	handler := HandlersStruct{
		Items:      itemsHandler,
		Categories: categoriesHandler,
	}
	return handler, nil
}

type HandlersStruct struct {
	Items      handlers.ItemsHandler
	Categories handlers.CategoriesHandler
}
