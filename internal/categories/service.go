package categories

import (
	"context"
	"fmt"
	"strings"

	"github.com/agustinrabini/go-toolkit/goutils/apierrors"
	"github.com/agustinrabini/api-items-project/cmd/api/config"
	"github.com/agustinrabini/api-items-project/internal/domain"
)

type Service interface {
	Create(ctx context.Context, input domain.Category) (string, apierrors.ApiError)
	Update(ctx context.Context, input domain.Category) apierrors.ApiError
	Delete(ctx context.Context, categoryID string) (string, apierrors.ApiError)
	Get(ctx context.Context, categoryID string) (domain.Category, apierrors.ApiError)
	GetAllCategories(ctx context.Context) ([]domain.Category, apierrors.ApiError)
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	if config.IsProductionEnvironment() {
		return NewServiceImpl(repository)
	}
	return NewItemsServiceMock()
}

func NewServiceImpl(repository Repository) Service {
	return &service{repo: repository}
}

func (s *service) Get(ctx context.Context, categoryID string) (domain.Category, apierrors.ApiError) {

	category, err := s.repo.Get(ctx, categoryID)
	if err != nil {
		return domain.Category{}, err
	}

	return category, nil
}

func (s *service) Update(ctx context.Context, input domain.Category) apierrors.ApiError {

	caretogires, err := s.repo.GetAllCategories(ctx)
	if err != nil {
		return err
	}

	err = validateCategoryExistence(input.Name, caretogires)
	if err != nil {
		return err
	}

	err = s.repo.Update(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Delete(ctx context.Context, categoryID string) (string, apierrors.ApiError) {
	categoryID, err := s.repo.Delete(ctx, categoryID)
	if err != nil {
		return "", err
	}

	return categoryID, nil
}

func (s *service) Create(ctx context.Context, input domain.Category) (string, apierrors.ApiError) {

	caretogires, err := s.repo.GetAllCategories(ctx)
	if err != nil {
		return "", err
	}

	err = validateCategoryExistence(input.Name, caretogires)
	if err != nil {
		return "", err
	}

	id, err := s.repo.Create(ctx, input)
	if err != nil {
		return "", err
	}

	if id == "" {
		return "", apierrors.NewInternalServerApiError("the id of the created item category is nil, should not", fmt.Errorf("an unexpected error occurred when creating the item category"))
	}

	return id, nil
}

func (s *service) GetAllCategories(ctx context.Context) ([]domain.Category, apierrors.ApiError) {

	categories, err := s.repo.GetAllCategories(ctx)
	if err != nil {
		return []domain.Category{}, err
	}

	if len(categories) == 0 {
		return []domain.Category{}, nil
	}

	return categories, nil
}

func validateCategoryExistence(inputCategoryName string, categories []domain.Category) apierrors.ApiError {

	for _, c := range categories {

		if strings.ToLower(inputCategoryName) == strings.ToLower(c.Name) {
			return apierrors.NewValidationApiError("category item alredy exists", "this category item alredy exists", apierrors.CauseList{})
		}
	}

	return nil
}
