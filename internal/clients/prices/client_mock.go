package prices

import (
	"context"

	"github.com/agustinrabini/go-toolkit/goutils/apierrors"
	"github.com/agustinrabini/api-items-project/internal/domain"
)

type ClientMock struct {
	HandleGetPriceByItemID func(ctx context.Context, itemID string) (domain.Price, apierrors.ApiError)
	HandleCreatePrice      func(ctx context.Context, price *domain.Price) apierrors.ApiError
	HandleModifyPrice      func(ctx context.Context, price *domain.Price) apierrors.ApiError
	HandleGetItemsPrices   func(ctx context.Context, itemsIDs []string) (domain.Prices, apierrors.ApiError)
}

func NewClientMock() ClientMock {
	return ClientMock{}
}

func (mock ClientMock) GetPriceByItemID(ctx context.Context, itemID string) (domain.Price, apierrors.ApiError) {
	if mock.HandleGetPriceByItemID != nil {
		return mock.HandleGetPriceByItemID(ctx, itemID)
	}
	return domain.Price{
		Amount: 100,
		Currency: domain.Currency{
			ID: "ARS",
		},
	}, nil
}

func (mock ClientMock) CreatePrice(ctx context.Context, price *domain.Price) apierrors.ApiError {
	if mock.HandleCreatePrice != nil {
		return mock.HandleCreatePrice(ctx, price)
	}
	return nil
}

func (mock ClientMock) ModifyPrice(ctx context.Context, price *domain.Price) apierrors.ApiError {
	if mock.HandleModifyPrice != nil {
		return mock.HandleModifyPrice(ctx, price)
	}
	return nil
}

func (mock ClientMock) GetItemsPrices(ctx context.Context, itemsIDs []string) (domain.Prices, apierrors.ApiError) {
	if mock.HandleGetItemsPrices != nil {
		return mock.HandleGetItemsPrices(ctx, itemsIDs)
	}
	return domain.Prices{}, nil
}
