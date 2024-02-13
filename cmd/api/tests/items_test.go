package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/agustinrabini/go-toolkit/goutils/apierrors"
	"github.com/agustinrabini/api-items-project/cmd/api/handlers"
	"github.com/agustinrabini/api-items-project/internal/domain"
	"github.com/agustinrabini/api-items-project/internal/items"
	"github.com/agustinrabini/api-items-project/internal/mocks"
	"github.com/agustinrabini/api-items-project/internal/platform/tests"
	"github.com/stretchr/testify/assert"
)

const (
	mockItemBody = `{
		"name": "Item mock",
		"shop_id": "1",
		"description": "Test description"
	}`
)

func TestGetItemByIDOk(t *testing.T) {
	teardownTestCase := tests.CustomSetupTestCase(beforeTestCase, afterTestCase)
	defer teardownTestCase()

	mockService := items.NewItemsServiceMock()
	itemsHandler := handlers.NewItemsHandler(mockService)
	depend.Items = itemsHandler

	response := executeRequest(buildRouter(), "GET", "/items/1", nil, "")
	assert.Equal(t, http.StatusOK, response.Code)
	var responseBody domain.Item
	err := json.Unmarshal(response.Body.Bytes(), &responseBody)
	assert.Nil(t, err)
	assert.Equal(t, mocks.ItemMock, responseBody)
}

func TestGetItemByIDErr(t *testing.T) {
	teardownTestCase := tests.CustomSetupTestCase(beforeTestCase, afterTestCase)
	defer teardownTestCase()

	mockService := items.NewItemsServiceMock()
	mockService.HandleGet = func(ctx context.Context, shopID string) (domain.Item, apierrors.ApiError) {
		return domain.Item{}, apierrors.NewInternalServerApiError("mock error", nil)
	}
	itemsHandler := handlers.NewItemsHandler(mockService)
	depend.Items = itemsHandler

	response := executeRequest(buildRouter(), "GET", "/items/1", nil, "")
	assert.Equal(t, http.StatusInternalServerError, response.Code)
}

func TestGetItemsByShopIDOk(t *testing.T) {
	teardownTestCase := tests.CustomSetupTestCase(beforeTestCase, afterTestCase)
	defer teardownTestCase()

	mockService := items.NewItemsServiceMock()
	itemsHandler := handlers.NewItemsHandler(mockService)
	depend.Items = itemsHandler

	response := executeRequest(buildRouter(), "GET", "/items/shop/1", nil, "")
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetItemsByShopIDErr(t *testing.T) {
	teardownTestCase := tests.CustomSetupTestCase(beforeTestCase, afterTestCase)
	defer teardownTestCase()

	mockService := items.NewItemsServiceMock()
	mockService.HandleGetItemsByShopID = func(ctx context.Context, shopID string) (domain.ItemsOutput, apierrors.ApiError) {
		return domain.ItemsOutput{}, apierrors.NewInternalServerApiError("mock error", nil)
	}
	itemsHandler := handlers.NewItemsHandler(mockService)
	depend.Items = itemsHandler

	response := executeRequest(buildRouter(), "GET", "/items/shop/1", nil, "")
	assert.Equal(t, http.StatusInternalServerError, response.Code)
}

func TestCreateItemOk(t *testing.T) {
	teardownTestCase := tests.CustomSetupTestCase(beforeTestCase, afterTestCase)
	defer teardownTestCase()

	mockService := items.NewItemsServiceMock()
	itemsHandler := handlers.NewItemsHandler(mockService)
	depend.Items = itemsHandler

	response := executeRequest(buildRouter(), "POST", "/items", nil, mockItemBody)
	assert.Equal(t, http.StatusCreated, response.Code)
}

func TestCreateItemErr(t *testing.T) {
	teardownTestCase := tests.CustomSetupTestCase(beforeTestCase, afterTestCase)
	defer teardownTestCase()

	mockService := items.NewItemsServiceMock()
	mockService.HandleCreateItem = func(ctx context.Context, item domain.Item) (interface{}, apierrors.ApiError) {
		return nil, apierrors.NewInternalServerApiError("mock error", nil)
	}
	itemsHandler := handlers.NewItemsHandler(mockService)
	depend.Items = itemsHandler

	response := executeRequest(buildRouter(), "POST", "/items", nil, mockItemBody)
	assert.Equal(t, http.StatusInternalServerError, response.Code)
}
