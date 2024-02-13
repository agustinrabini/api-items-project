package categories

import (
	"context"
	"errors"
	"fmt"

	"github.com/agustinrabini/go-toolkit/gonosql"
	"github.com/agustinrabini/go-toolkit/goutils/apierrors"
	"github.com/agustinrabini/api-items-project/cmd/api/config"
	"github.com/agustinrabini/api-items-project/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

const (
	errorInDB = "[%s] Error in DB"
)

var ErrItemNotFound = apierrors.NewNotFoundApiError("categories not found")

type Repository interface {
	Create(ctx context.Context, input domain.Category) (string, apierrors.ApiError)
	Update(ctx context.Context, input domain.Category) apierrors.ApiError
	Delete(ctx context.Context, categoryID string) (string, apierrors.ApiError)
	Get(ctx context.Context, categoryID string) (domain.Category, apierrors.ApiError)
	GetAllCategories(ctx context.Context) ([]domain.Category, apierrors.ApiError)
}

type categoriesRepository struct {
	Collection *mongo.Collection
}

func NewRepository(Collection *mongo.Collection) Repository {
	if config.IsProductionEnvironment() {
		return NewRepositoryImpl(Collection)
	}
	return NewCategoriesRepositoryMock()
}

func NewRepositoryImpl(Collection *mongo.Collection) Repository {
	return &categoriesRepository{Collection: Collection}
}

func (storage *categoriesRepository) Get(ctx context.Context, categoryID string) (domain.Category, apierrors.ApiError) {
	var models domain.Category
	result, err := gonosql.Get(ctx, storage.Collection, categoryID)
	if err != nil {
		return domain.Category{}, apierrors.NewBadRequestApiError(fmt.Sprintf(errorInDB, "Get") + ": " + err.Error())
	}
	if errors.Is(result.Err(), mongo.ErrNoDocuments) {
		return domain.Category{}, ErrItemNotFound
	}
	if result.Err() != nil {
		return domain.Category{}, apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "Get"), result.Err())
	}
	err = result.Decode(&models)
	if err != nil {
		return domain.Category{}, apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "Get"), err)
	}
	return models, nil

}

func (storage *categoriesRepository) Update(ctx context.Context, input domain.Category) apierrors.ApiError {

	primitiveID, err := primitive.ObjectIDFromHex(input.ID)
	if err != nil {
		return apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "Update"), err)
	}

	update := bson.M{
		"$set": bson.M{
			"name": input.Name,
		},
	}
	result, err := storage.Collection.UpdateOne(ctx, bson.M{"_id": primitiveID}, update)
	if err != nil {
		return apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "Update"), err)
	}
	if result.MatchedCount == 0 {
		return apierrors.NewNotFoundApiError(fmt.Sprintf(errorInDB, "Update Category"))
	}

	if result.ModifiedCount == 0 {
		return apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "Update Category"), fmt.Errorf("no update was made"))
	}
	return nil
}

func (storage *categoriesRepository) Delete(ctx context.Context, categoryID string) (string, apierrors.ApiError) {

	result, err := gonosql.Delete(ctx, storage.Collection, categoryID)
	if err != nil {
		return "", apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "Delete"), err)
	}

	if result.DeletedCount == 0 {
		return "", apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "Delete"), fmt.Errorf("not deletiong was made, it was expected to be so"))
	}

	return categoryID, nil
}

func (storage *categoriesRepository) Create(ctx context.Context, input domain.Category) (string, apierrors.ApiError) {

	result, err := gonosql.InsertOne(ctx, storage.Collection, input)
	if err != nil {
		return "", apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "Save"), err)
	}

	return fmt.Sprint(result.InsertedID), nil

}

func (storage *categoriesRepository) GetAllCategories(ctx context.Context) ([]domain.Category, apierrors.ApiError) {

	var categories []domain.Category
	var c domain.Category

	filter := bson.M{}

	cursor, err := storage.Collection.Find(ctx, filter)
	if err != nil {
		return []domain.Category{}, apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "GetAll")+": "+err.Error(), err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		err := cursor.Decode(&c)
		if err != nil {
			return []domain.Category{}, apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "GetAll")+": "+err.Error(), err)
		}

		categories = append(categories, c)
	}

	if err := cursor.Err(); err != nil {
		return []domain.Category{}, apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "GetAll")+": "+err.Error(), err)
	}

	return categories, nil

}
