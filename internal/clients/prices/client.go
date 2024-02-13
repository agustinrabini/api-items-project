package prices

import (
	"context"

	"github.com/agustinrabini/api-items-project/cmd/api/config"

	"github.com/agustinrabini/go-toolkit/goutils/apierrors"
	"github.com/agustinrabini/api-items-project/internal/domain"
)

var (
	ClientInstance = newPricesApiClient()
)

type Client interface {
	GetPriceByItemID(ctx context.Context, itemID string) (domain.Price, apierrors.ApiError)
	GetItemsPrices(ctx context.Context, itemsIDs []string) (domain.Prices, apierrors.ApiError)
	CreatePrice(ctx context.Context, price *domain.Price) apierrors.ApiError
	ModifyPrice(ctx context.Context, price *domain.Price) apierrors.ApiError
}

func newPricesApiClient() Client {
	if config.IsProductionEnvironment() {
		return NewClientImpl()
	}
	return NewClientMock()
}
