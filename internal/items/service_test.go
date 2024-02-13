package items

import (
	"context"
	"testing"

	"github.com/agustinrabini/api-items-project/internal/clients/prices"
	"github.com/agustinrabini/api-items-project/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestService_NewService(t *testing.T) {
	repoMock := NewItemsRepositoryMock()
	itemsService := NewService(repoMock, prices.ClientInstance)
	assert.NotNil(t, itemsService)
}

func TestService_Get(t *testing.T) {
	repoMock := NewItemsRepositoryMock()
	itemsService := NewServiceImpl(repoMock, prices.ClientInstance)
	_, err := itemsService.Get(context.Background(), "JPT-1234")
	assert.NoError(t, err)
}

func TestService_Save(t *testing.T) {
	repoMock := NewItemsRepositoryMock()
	itemsService := NewServiceImpl(repoMock, prices.ClientInstance)
	_, err := itemsService.CreateItem(context.Background(), domain.Item{Name: "Test Item"})
	assert.NoError(t, err)
}

func TestService_Delete(t *testing.T) {
	repoMock := NewItemsRepositoryMock()
	itemsService := NewServiceImpl(repoMock, prices.ClientInstance)
	err := itemsService.Delete(context.Background(), "JPT-1234")
	assert.NoError(t, err)
}

func TestService_GetItemsByShopID(t *testing.T) {
	repoMock := NewItemsRepositoryMock()
	itemsService := NewServiceImpl(repoMock, prices.ClientInstance)
	_, err := itemsService.GetItemsByShopID(context.Background(), "JPT-1234")
	assert.NoError(t, err)
}
