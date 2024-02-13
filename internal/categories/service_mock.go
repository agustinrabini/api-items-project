package categories

import (
	"context"

	"github.com/agustinrabini/go-toolkit/goutils/apierrors"
	"github.com/agustinrabini/api-items-project/internal/domain"
)

type ServiceMock struct {
	HandleUpdate func(ctx context.Context, input domain.Category) apierrors.ApiError
	HandleDelete func(ctx context.Context, categoryID string) (string, apierrors.ApiError)
	HandleGet    func(ctx context.Context, categoryID string) (domain.Category, apierrors.ApiError)
	HandleCreate func(ctx context.Context, input domain.Category) (string, apierrors.ApiError)

	HandleGetAllCategories func(ctx context.Context) ([]domain.Category, apierrors.ApiError)
}

func NewItemsServiceMock() ServiceMock {
	return ServiceMock{}
}

func (mock ServiceMock) Update(ctx context.Context, input domain.Category) apierrors.ApiError {
	if mock.HandleUpdate != nil {
		return mock.Update(ctx, input)
	}
	return nil
}

func (mock ServiceMock) Delete(ctx context.Context, categoryID string) (string, apierrors.ApiError) {
	if mock.HandleDelete != nil {
		return mock.Delete(ctx, categoryID)
	}
	return "", nil
}

func (mock ServiceMock) Get(ctx context.Context, categoryID string) (domain.Category, apierrors.ApiError) {
	if mock.HandleGet != nil {
		return mock.Get(ctx, categoryID)
	}
	return domain.Category{}, nil
}

func (mock ServiceMock) Create(ctx context.Context, input domain.Category) (string, apierrors.ApiError) {
	if mock.HandleGetAllCategories != nil {
		return mock.Create(ctx, domain.Category{})
	}
	return "", nil
}

func (mock ServiceMock) GetAllCategories(ctx context.Context) ([]domain.Category, apierrors.ApiError) {
	if mock.HandleGetAllCategories != nil {
		return mock.GetAllCategories(ctx)
	}
	return []domain.Category{}, nil
}
