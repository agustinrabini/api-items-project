package items

import (
	"context"

	"github.com/agustinrabini/go-toolkit/goutils/apierrors"
	"github.com/agustinrabini/api-items-project/internal/domain"
	"github.com/agustinrabini/api-items-project/internal/dto"
	"github.com/agustinrabini/api-items-project/internal/mocks"
	"github.com/gin-gonic/gin"
)

type ServiceMock struct {
	HandleGet                      func(ctx context.Context, itemID string) (domain.Item, apierrors.ApiError)
	HandleDelete                   func(ctx context.Context, itemID string) apierrors.ApiError
	HandleCreateItem               func(c *gin.Context, itemRequest dto.ItemDTO) (interface{}, apierrors.ApiError)
	HandleGetAll                   func(ctx context.Context) ([]domain.Item, apierrors.ApiError)
	HandleGetItemsByUserID         func(ctx context.Context, userID string) (domain.ItemsOutput, apierrors.ApiError)
	HandleGetItemsByShopID         func(ctx context.Context, shopID string) (domain.ItemsOutput, apierrors.ApiError)
	HandleGetItemsByShopCategoryID func(ctx context.Context, shopID string, categoryID string) (domain.ItemsOutput, apierrors.ApiError)
	HandleGetItemsByIDs            func(ctx context.Context, items domain.ItemsIds) (domain.ItemsOutput, apierrors.ApiError)
	HandleUpdate                   func(ctx context.Context, itemID string, itemRequest dto.ItemDTO) apierrors.ApiError
}

func NewItemsServiceMock() ServiceMock {
	return ServiceMock{}
}

func (mock ServiceMock) Get(ctx context.Context, itemID string) (domain.Item, apierrors.ApiError) {
	if mock.HandleGet != nil {
		return mock.HandleGet(ctx, itemID)
	}
	return mocks.ItemMock, nil
}

func (mock ServiceMock) Delete(ctx context.Context, itemID string) apierrors.ApiError {
	if mock.HandleDelete != nil {
		return mock.HandleDelete(ctx, itemID)
	}
	return nil
}

func (mock ServiceMock) CreateItem(c *gin.Context, itemRequest dto.ItemDTO) (interface{}, apierrors.ApiError) {
	if mock.HandleCreateItem != nil {
		return mock.HandleCreateItem(c, itemRequest)
	}
	return nil, nil
}

func (mock ServiceMock) GetAll(ctx context.Context) ([]domain.Item, apierrors.ApiError) {
	if mock.HandleGetAll != nil {
		return mock.HandleGetAll(ctx)
	}
	return nil, nil
}

func (mock ServiceMock) GetItemsByUserID(ctx context.Context, userID string) (domain.ItemsOutput, apierrors.ApiError) {
	if mock.HandleGetItemsByUserID != nil {
		return mock.HandleGetItemsByUserID(ctx, userID)
	}
	return domain.ItemsOutput{}, nil
}

func (mock ServiceMock) GetItemsByShopID(ctx context.Context, shopID string) (domain.ItemsOutput, apierrors.ApiError) {
	if mock.HandleGetItemsByShopID != nil {
		return mock.HandleGetItemsByShopID(ctx, shopID)
	}
	return domain.ItemsOutput{}, nil
}

func (mock ServiceMock) GetItemsByShopCategoryID(ctx context.Context, shopID string, categoryID string) (domain.ItemsOutput, apierrors.ApiError) {
	if mock.HandleGetItemsByShopCategoryID != nil {
		return mock.HandleGetItemsByShopCategoryID(ctx, shopID, categoryID)
	}
	return domain.ItemsOutput{}, nil
}

func (mock ServiceMock) GetItemsByIDs(ctx context.Context, items domain.ItemsIds) (domain.ItemsOutput, apierrors.ApiError) {
	if mock.HandleGetItemsByIDs != nil {
		return mock.HandleGetItemsByIDs(ctx, items)
	}
	return domain.ItemsOutput{}, nil
}

func (mock ServiceMock) Update(ctx context.Context, itemID string, itemRequest dto.ItemDTO) apierrors.ApiError {
	if mock.HandleUpdate != nil {
		return mock.HandleUpdate(ctx, itemID, itemRequest)
	}
	return nil
}
