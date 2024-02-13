package items

import (
	"context"

	"github.com/agustinrabini/go-toolkit/goutils/apierrors"
	"github.com/agustinrabini/api-items-project/internal/domain"
)

type RepositoryMock struct {
	HandleGet                 func(ctx context.Context, itemID string) (domain.Item, apierrors.ApiError)
	HandleDelete              func(ctx context.Context, itemID string) apierrors.ApiError
	HandleSave                func(ctx context.Context, item domain.Item) (interface{}, apierrors.ApiError)
	HandleGetByUserID         func(ctx context.Context, userID string) ([]domain.Item, apierrors.ApiError)
	HandleGetByShopID         func(ctx context.Context, shopID string) ([]domain.Item, apierrors.ApiError)
	HandleGetByShopCategoryID func(ctx context.Context, shopID, categoryID string) ([]domain.Item, apierrors.ApiError)
	HandleGetByIDs            func(ctx context.Context, itemsIDs []string) ([]domain.Item, apierrors.ApiError)
	HandleUpdate              func(ctx context.Context, itemID string, updateItem *domain.Item) apierrors.ApiError
}

func NewItemsRepositoryMock() RepositoryMock {
	return RepositoryMock{}
}

func (mock RepositoryMock) Get(ctx context.Context, itemID string) (domain.Item, apierrors.ApiError) {
	if mock.HandleGet != nil {
		return mock.HandleGet(ctx, itemID)
	}
	return domain.Item{}, nil
}

func (mock RepositoryMock) Delete(ctx context.Context, itemID string) apierrors.ApiError {
	if mock.HandleDelete != nil {
		return mock.HandleDelete(ctx, itemID)
	}
	return nil
}

func (mock RepositoryMock) Save(ctx context.Context, item domain.Item) (interface{}, apierrors.ApiError) {
	if mock.HandleSave != nil {
		return mock.HandleSave(ctx, item)
	}
	return nil, nil
}

func (mock RepositoryMock) GetByUserID(ctx context.Context, userID string) ([]domain.Item, apierrors.ApiError) {
	if mock.HandleGetByUserID != nil {
		return mock.HandleGetByUserID(ctx, userID)
	}
	return nil, nil
}

func (mock RepositoryMock) GetByShopID(ctx context.Context, shopID string) ([]domain.Item, apierrors.ApiError) {
	if mock.HandleGetByShopID != nil {
		return mock.HandleGetByShopID(ctx, shopID)
	}
	return nil, nil
}

func (mock RepositoryMock) GetByShopCategoryID(ctx context.Context, shopID, categoryID string) ([]domain.Item, apierrors.ApiError) {
	if mock.HandleGetByShopCategoryID != nil {
		return mock.HandleGetByShopCategoryID(ctx, shopID, categoryID)
	}
	return nil, nil
}

func (mock RepositoryMock) GetByIDs(ctx context.Context, itemsIDs []string) ([]domain.Item, apierrors.ApiError) {
	if mock.HandleGetByIDs != nil {
		return mock.HandleGetByIDs(ctx, itemsIDs)
	}
	return nil, nil
}

func (mock RepositoryMock) Update(ctx context.Context, itemID string, updateItem *domain.Item) apierrors.ApiError {
	if mock.HandleUpdate != nil {
		return mock.HandleUpdate(ctx, itemID, updateItem)
	}
	return nil
}
