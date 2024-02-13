package categories

import (
	"context"

	"github.com/agustinrabini/go-toolkit/goutils/apierrors"
	"github.com/agustinrabini/api-items-project/internal/domain"
)

type RepositoryMock struct {
	HandleCreate           func(ctx context.Context, input domain.Category) (string, apierrors.ApiError)
	HandleUpdate           func(ctx context.Context, input domain.Category) apierrors.ApiError
	HandleDelete           func(ctx context.Context, categoryID string) (string, apierrors.ApiError)
	HandleGet              func(ctx context.Context, categoryID string) (domain.Category, apierrors.ApiError)
	HandleGetAllCategories func(ctx context.Context) ([]domain.Category, apierrors.ApiError)
}

func NewCategoriesRepositoryMock() RepositoryMock {
	return RepositoryMock{}
}
func (mock RepositoryMock) Update(ctx context.Context, input domain.Category) apierrors.ApiError {

	if mock.HandleUpdate != nil {
		return mock.HandleUpdate(ctx, input)
	}
	return nil

}

func (mock RepositoryMock) Delete(ctx context.Context, categoryID string) (string, apierrors.ApiError) {

	if mock.HandleDelete != nil {
		return mock.HandleDelete(ctx, categoryID)
	}
	return "", nil

}

func (mock RepositoryMock) Get(ctx context.Context, categoryID string) (domain.Category, apierrors.ApiError) {

	if mock.HandleGet != nil {
		return mock.HandleGet(ctx, categoryID)
	}
	return domain.Category{}, nil

}

func (mock RepositoryMock) Create(ctx context.Context, input domain.Category) (string, apierrors.ApiError) {

	if mock.HandleCreate != nil {
		return mock.HandleCreate(ctx, input)
	}
	return "", nil

}

func (mock RepositoryMock) GetAllCategories(ctx context.Context) ([]domain.Category, apierrors.ApiError) {
	if mock.HandleGetAllCategories != nil {
		return mock.HandleGetAllCategories(ctx)
	}
	return []domain.Category{}, nil
}
