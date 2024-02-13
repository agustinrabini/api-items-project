package items

import (
	"context"
	"errors"
	"fmt"

	"github.com/agustinrabini/api-items-project/cmd/api/config"

	"github.com/agustinrabini/go-toolkit/gonosql"
	"github.com/agustinrabini/go-toolkit/goutils/apierrors"
	"github.com/agustinrabini/api-items-project/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

const (
	errorInDB = "[%s] Error in DB"
)

var ErrItemNotFound = apierrors.NewNotFoundApiError("item not found")

type Repository interface {
	Get(ctx context.Context, itemID string) (domain.Item, apierrors.ApiError)
	Delete(ctx context.Context, itemID string) apierrors.ApiError
	Save(ctx context.Context, item domain.Item) (interface{}, apierrors.ApiError)
	GetByUserID(ctx context.Context, userID string) ([]domain.Item, apierrors.ApiError)
	GetByShopID(ctx context.Context, shopID string) ([]domain.Item, apierrors.ApiError)
	GetByShopCategoryID(ctx context.Context, shopID, categoryID string) ([]domain.Item, apierrors.ApiError)
	GetByIDs(ctx context.Context, itemsIDs []string) ([]domain.Item, apierrors.ApiError)
	Update(ctx context.Context, itemID string, updateItem *domain.Item) apierrors.ApiError
}

type itemRepository struct {
	Collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) Repository {
	if config.IsProductionEnvironment() {
		return NewRepositoryImpl(collection)
	}
	return NewItemsRepositoryMock()
}

func NewRepositoryImpl(collection *mongo.Collection) Repository {
	return &itemRepository{Collection: collection}
}

func (storage *itemRepository) Get(ctx context.Context, itemID string) (domain.Item, apierrors.ApiError) {
	var models domain.Item
	result, err := gonosql.Get(ctx, storage.Collection, itemID)
	if err != nil {
		return domain.Item{}, apierrors.NewBadRequestApiError(fmt.Sprintf(errorInDB, "Get") + ": " + err.Error())
	}
	if errors.Is(result.Err(), mongo.ErrNoDocuments) {
		return domain.Item{}, ErrItemNotFound
	}
	if result.Err() != nil {
		return domain.Item{}, apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "Get"), result.Err())
	}
	err = result.Decode(&models)
	if err != nil {
		return domain.Item{}, apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "Get"), err)
	}
	return models, nil
}

func (storage *itemRepository) Save(ctx context.Context, item domain.Item) (interface{}, apierrors.ApiError) {
	result, err := gonosql.InsertOne(ctx, storage.Collection, item)
	if err != nil {
		return nil, apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "Save"), err)
	}
	return result.InsertedID, nil
}

func (storage *itemRepository) Delete(ctx context.Context, itemID string) apierrors.ApiError {

	_, err := gonosql.Delete(ctx, storage.Collection, itemID)
	if err != nil {
		return apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "Delete"), err)
	}
	return nil
}

func (storage *itemRepository) GetByUserID(ctx context.Context, userID string) ([]domain.Item, apierrors.ApiError) {
	var models []domain.Item
	var i domain.Item

	filter := bson.M{
		"user_id": userID,
	}

	cursor, err := storage.Collection.Find(ctx, filter)
	if err != nil {
		return []domain.Item{}, apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "GetAll")+": "+err.Error(), err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		err := cursor.Decode(&i)
		if err != nil {
			return []domain.Item{}, apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "GetAll")+": "+err.Error(), err)
		}

		models = append(models, i)
	}

	if err := cursor.Err(); err != nil {
		return []domain.Item{}, apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "GetAll")+": "+err.Error(), err)
	}

	return models, nil
}

func (storage *itemRepository) GetByShopID(ctx context.Context, shopID string) ([]domain.Item, apierrors.ApiError) {
	var models []domain.Item
	cursorResult, err := gonosql.GetByKey(ctx, storage.Collection, "shop_id", shopID)
	if err != nil {
		return nil, apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "GetByShopID"), err)
	}
	if err = cursorResult.All(ctx, &models); err != nil {
		return models, apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "GetByShopID"), err)
	}
	return models, nil
}

func (storage *itemRepository) GetByShopCategoryID(ctx context.Context, shopID, categoryID string) ([]domain.Item, apierrors.ApiError) {
	var models []domain.Item
	filter := bson.M{"shop_id": shopID, "category._id": categoryID}
	cursor, err := gonosql.GetByFilter(ctx, storage.Collection, filter)
	if err != nil {
		return nil, apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "GetByShopCategoryID"), err)
	}
	if cursor == nil {
		return nil, nil
	}
	if err = cursor.All(ctx, &models); err != nil {
		return nil, apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "GetByShopCategoryID"), err)
	}
	return models, nil
}

func (storage *itemRepository) GetByIDs(ctx context.Context, itemsIDs []string) ([]domain.Item, apierrors.ApiError) {
	var models []domain.Item
	cursor, err := gonosql.GetByIDs(ctx, storage.Collection, itemsIDs)
	if err != nil {
		return nil, apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "GetByIDs"), err)
	}
	if cursor == nil {
		return nil, nil
	}
	if err = cursor.All(ctx, &models); err != nil {
		return nil, apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "GetByIDs"), err)
	}
	/* TODO que no tenga que realizar toda la consulta para fallar si un ID no existe, buscar la forma que mongo db falle,
	*  en node.js existe algo llamado strict, tratar de implementarlo o buscar esa solucion.
	 */
	if len(models) != len(itemsIDs) {
		return nil, apierrors.NewNotFoundApiError(fmt.Sprintf(errorInDB, "GetByIDs") + ": Some items not found")
	}
	return models, nil
}

func (storage *itemRepository) Update(ctx context.Context, itemID string, updateItem *domain.Item) apierrors.ApiError {
	// TODO review result for Update?
	_, err := gonosql.Update(ctx, storage.Collection, itemID, updateItem)
	if err != nil {
		return apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "Update"), err)
	}
	return nil
}
