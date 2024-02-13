package dependencies

import (
	"github.com/agustinrabini/api-items-project/cmd/api/config"

	"github.com/agustinrabini/go-toolkit/gonosql"
	"github.com/agustinrabini/api-items-project/internal/categories"
	"github.com/agustinrabini/api-items-project/internal/items"
	"github.com/agustinrabini/api-items-project/internal/platform/storage"
)

const (
	KvsItemsCollection      = "items"
	KvsCategoriesCollection = "categories"
)

type DependencyManager struct {
	*gonosql.Data
}

func NewDependencyManager() DependencyManager {
	var db *gonosql.Data
	if config.IsProductionEnvironment() {
		db = storage.NewNoSQL()
	} else {
		db = storage.NewMock()
	}
	if db.Error != nil {
		panic(db.Error)
	}
	return DependencyManager{
		db,
	}
}

func (m DependencyManager) ItemsRepository() items.Repository {
	return items.NewRepository(m.NewCollection(KvsItemsCollection))
}

func (m DependencyManager) CategoriesRepository() categories.Repository {
	return categories.NewRepository(m.NewCollection(KvsCategoriesCollection))
}
