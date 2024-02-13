package items

import (
	"context"

	"github.com/agustinrabini/api-items-project/cmd/api/config"
	"github.com/gin-gonic/gin"

	"github.com/agustinrabini/go-toolkit/goutils/apierrors"
	"github.com/agustinrabini/api-items-project/internal/clients/prices"
	"github.com/agustinrabini/api-items-project/internal/clients/shops"
	"github.com/agustinrabini/api-items-project/internal/domain"
	"github.com/agustinrabini/api-items-project/internal/dto"
	"github.com/creasty/defaults"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	Get(ctx context.Context, itemID string) (domain.Item, apierrors.ApiError)
	GetItemsByUserID(ctx context.Context, userID string) (domain.ItemsOutput, apierrors.ApiError)
	GetItemsByShopID(ctx context.Context, shopID string) (domain.ItemsOutput, apierrors.ApiError)
	GetItemsByIDs(ctx context.Context, items domain.ItemsIds) (domain.ItemsOutput, apierrors.ApiError)
	GetItemsByShopCategoryID(ctx context.Context, shopID string, categoryID string) (domain.ItemsOutput, apierrors.ApiError)
	Delete(ctx context.Context, itemID string) apierrors.ApiError
	CreateItem(c *gin.Context, itemRequest dto.ItemDTO) (interface{}, apierrors.ApiError)
	Update(ctx context.Context, itemID string, itemRequest dto.ItemDTO) apierrors.ApiError
}

type service struct {
	repo         Repository
	pricesClient prices.Client
	shopsClient  shops.Client
}

func NewService(repository Repository, pricesClient prices.Client, shopsClient shops.Client) Service {
	if config.IsProductionEnvironment() {
		return NewServiceImpl(repository, pricesClient, shopsClient)
	}
	return NewItemsServiceMock()
}

func NewServiceImpl(repository Repository, pricesClient prices.Client, shopsClient shops.Client) Service {
	return &service{repo: repository, pricesClient: pricesClient, shopsClient: shopsClient}
}

func (s *service) Get(ctx context.Context, itemID string) (domain.Item, apierrors.ApiError) {
	var (
		item  domain.Item
		err   apierrors.ApiError
		price domain.Price
		//wg    sync.WaitGroup
	)

	item, err = s.repo.Get(ctx, itemID)
	if err != nil {
		return domain.Item{}, err
	}

	price, err = s.pricesClient.GetPriceByItemID(ctx, itemID)
	if err != nil {
		return domain.Item{}, err
	}

	item.Price = price

	item.Validate()

	return item, nil
}

func (s *service) Delete(ctx context.Context, itemID string) apierrors.ApiError {
	return s.repo.Delete(ctx, itemID)
}

func (s *service) CreateItem(c *gin.Context, itemRequest dto.ItemDTO) (interface{}, apierrors.ApiError) {

	shopID, err := s.shopsClient.GetShopByUserID(c)
	if err != nil {
		return nil, apierrors.NewInternalServerApiError("", err)
	}

	item, err2 := itemRequest.ToItem()
	if err2 != nil {
		return nil, apierrors.NewBadRequestApiError("Error convert body to domain: " + err.Error())
	}
	item.ShopID = shopID.ID

	err2 = defaults.Set(&item)
	if err != nil {
		return nil, apierrors.NewInternalServerApiError("set default errors: ", err)
	}
	completeEligibleIDs(&item)
	insertedID, apiErr := s.repo.Save(c, item)
	if apiErr != nil {
		return nil, apiErr
	}
	item.Price.ItemId = insertedID.(primitive.ObjectID).Hex()
	apiErr = s.pricesClient.CreatePrice(c, &item.Price)
	if apiErr != nil {
		return nil, apiErr
	}
	return insertedID, nil
}

func (s *service) GetItemsByUserID(ctx context.Context, userID string) (domain.ItemsOutput, apierrors.ApiError) {

	var itemsIDs []string

	items, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return domain.ItemsOutput{}, err
	}

	for _, item := range items {
		itemsIDs = append(itemsIDs, item.ID)
	}

	response, err := s.pricesClient.GetItemsPrices(ctx, itemsIDs)
	if err != nil {
		return domain.ItemsOutput{}, err
	}

	itemsWithPrice := pricesResponseToItemPrice(response, items)
	items = itemsWithPrice

	return domain.ItemsOutput{Items: itemsWithPrice}, nil
}

func (s *service) GetItemsByShopID(ctx context.Context, shopID string) (domain.ItemsOutput, apierrors.ApiError) {

	var itemsIDs []string

	items, err := s.repo.GetByShopID(ctx, shopID)
	if err != nil {
		return domain.ItemsOutput{}, err
	}

	for _, item := range items {
		itemsIDs = append(itemsIDs, item.ID)
	}

	response, err := s.pricesClient.GetItemsPrices(ctx, itemsIDs)
	if err != nil {
		return domain.ItemsOutput{}, err
	}

	itemsWithPrice := pricesResponseToItemPrice(response, items)
	items = itemsWithPrice

	return domain.ItemsOutput{Items: items}, nil
}

func (s *service) GetItemsByShopCategoryID(ctx context.Context, shopID string, categoryID string) (domain.ItemsOutput, apierrors.ApiError) {

	var itemsIDs []string

	items, err := s.repo.GetByShopCategoryID(ctx, shopID, categoryID)
	if err != nil {
		return domain.ItemsOutput{}, err
	}

	for _, item := range items {
		itemsIDs = append(itemsIDs, item.ID)
	}

	response, err := s.pricesClient.GetItemsPrices(ctx, itemsIDs)
	if err != nil {
		return domain.ItemsOutput{}, err
	}

	itemsWithPrice := pricesResponseToItemPrice(response, items)
	items = itemsWithPrice

	return domain.ItemsOutput{Items: items}, nil
}

func (s *service) GetItemsByIDs(ctx context.Context, itemsIds domain.ItemsIds) (domain.ItemsOutput, apierrors.ApiError) {

	var itemsIDs []string

	items, err := s.repo.GetByIDs(ctx, itemsIds.Items)
	if err != nil {
		return domain.ItemsOutput{}, err
	}

	for _, item := range items {
		itemsIDs = append(itemsIDs, item.ID)
	}

	response, err := s.pricesClient.GetItemsPrices(ctx, itemsIDs)
	if err != nil {
		return domain.ItemsOutput{}, err
	}

	itemsWithPrice := pricesResponseToItemPrice(response, items)
	items = itemsWithPrice

	return domain.ItemsOutput{Items: items}, nil
}

func (s *service) Update(ctx context.Context, itemID string, itemRequest dto.ItemDTO) apierrors.ApiError {
	updateItem, err := itemRequest.ToItem()
	if err != nil {
		return apierrors.NewBadRequestApiError("Error convert body to domain: " + err.Error())
	}
	apiErr := s.repo.Update(ctx, itemID, &updateItem)
	if apiErr != nil {
		return apiErr
	}
	return s.pricesClient.ModifyPrice(ctx, &updateItem.Price)
}

func completeEligibleIDs(item *domain.Item) {
	for i := range item.Eligible {
		item.Eligible[i].ID = primitive.NewObjectID().Hex()
	}
}
