package prices

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/agustinrabini/go-toolkit/goutils/apierrors"
	"github.com/agustinrabini/go-toolkit/rest"
	"github.com/agustinrabini/api-items-project/cmd/api/config"
	"github.com/agustinrabini/api-items-project/internal/domain"
)

const (
	pricesBaseEndpoint = "/prices"
	pricesItemsPrices  = "/items"
)

var errorPricingService error = fmt.Errorf("error at Prices Services. URL: ")
var errorPricingServiceUnmarshal error = fmt.Errorf("Unexpected error unmarshalling price json response. Body: ")

type ClientImpl struct {
	Client *rest.RequestBuilder
}

func NewClientImpl() *ClientImpl {
	customPool := &rest.CustomPool{
		MaxIdleConnsPerHost: 100,
	}
	restClientPrices := &rest.RequestBuilder{
		BaseURL:        config.InternalBasePricesClient,
		Timeout:        5 * time.Second,
		ContentType:    rest.JSON,
		EnableCache:    false,
		DisableTimeout: false,
		CustomPool:     customPool,
		FollowRedirect: true,
		//RetryStrategy:  retry.NewSimpleRetryStrategy(3, 30*time.Millisecond),
		MetricsConfig: rest.MetricsReportConfig{TargetId: "prices-api"},
	}
	return &ClientImpl{Client: restClientPrices}
}

func (c ClientImpl) GetPriceByItemID(ctx context.Context, itemID string) (domain.Price, apierrors.ApiError) {
	var rawPrice []byte

	endpoint := fmt.Sprintf("%s/item/%s", pricesBaseEndpoint, itemID)
	response := c.Client.Get(endpoint, rest.Context(ctx))

	if response.Response == nil || (response.StatusCode != http.StatusNotFound && response.StatusCode != http.StatusOK) {
		return domain.Price{}, apierrors.NewInternalServerApiError(fmt.Sprint("Unexpected error getting price, url: "+endpoint), response.Err)
	}

	rawPrice = response.Bytes()
	if response.StatusCode != http.StatusNotFound {
		rawPrice = response.Bytes()
	}

	if rawPrice != nil {
		price := domain.Price{}

		if unmarshallError := json.Unmarshal(rawPrice, &price); unmarshallError != nil {
			return domain.Price{}, apierrors.NewInternalServerApiError("Unexpected error unmarshalling price json response. value: "+string(rawPrice), unmarshallError)
		}

		return price, nil
	}

	return domain.Price{}, apierrors.NewNotFoundApiError("price not found")
}

func (c ClientImpl) GetItemsPrices(ctx context.Context, itemsIDs []string) (domain.Prices, apierrors.ApiError) {
	var rawPrice []byte
	prices := domain.Prices{}

	type request struct {
		ItemsIDs []string `json:"items_ids"`
	}

	req := request{
		ItemsIDs: itemsIDs,
	}

	endpoint := fmt.Sprintf("%s%s", pricesBaseEndpoint, pricesItemsPrices)
	response := c.Client.Post(endpoint, req)

	if response.Response == nil || (response.StatusCode != http.StatusNotFound && response.StatusCode != http.StatusOK) {
		return domain.Prices{}, apierrors.NewInternalServerApiError(fmt.Sprint(errorPricingService, endpoint), response.Err)
	}

	rawPrice = response.Bytes()
	if response.StatusCode != http.StatusNotFound {
		rawPrice = response.Bytes()
	}

	if rawPrice != nil {

		if unmarshallError := json.Unmarshal(rawPrice, &prices); unmarshallError != nil {
			return domain.Prices{}, apierrors.NewInternalServerApiError(errorPricingServiceUnmarshal.Error()+string(rawPrice), unmarshallError)
		}

		return prices, nil
	}

	return domain.Prices{}, apierrors.NewNotFoundApiError("price not found")
}

func (c ClientImpl) CreatePrice(ctx context.Context, price *domain.Price) apierrors.ApiError {
	var response *rest.Response

	fmt.Println("DEBUG price body:", price)

	bb, err := json.Marshal(price)
	if err != nil {
		fmt.Println("DEBUG ERROR MARHSALING", err.Error())
	}

	fmt.Println("DEBUG price body json:", string(bb))

	response = c.Client.Post(pricesBaseEndpoint, price, rest.Context(ctx))

	if response.Err != nil || response.Response == nil || response.StatusCode != http.StatusCreated {
		return apierrors.NewInternalServerApiError(fmt.Sprint("Unexpected error create price, url: "+pricesBaseEndpoint), response.Err)
	}

	return nil
}

func (c ClientImpl) ModifyPrice(ctx context.Context, price *domain.Price) apierrors.ApiError {
	var response *rest.Response
	endpoint := fmt.Sprintf("%s/%s", pricesBaseEndpoint, price.ID)
	response = c.Client.Put(endpoint, price, rest.Context(ctx))

	if response.Err != nil || response.Response == nil || response.StatusCode != http.StatusOK {
		return apierrors.NewInternalServerApiError(fmt.Sprint("Unexpected error create price, url: "+pricesBaseEndpoint), response.Err)
	}

	return nil
}
