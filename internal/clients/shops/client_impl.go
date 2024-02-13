package shops

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/agustinrabini/go-toolkit/goutils/apierrors"
	"github.com/agustinrabini/go-toolkit/rest"
	"github.com/agustinrabini/api-items-project/cmd/api/config"
	"github.com/agustinrabini/api-items-project/internal/dto"
	"github.com/gin-gonic/gin"
)

const (
	shopsBaseEndpoint = "/shops"
)

type ClientImpl struct {
	Client *rest.RequestBuilder
}

func NewClientImpl() *ClientImpl {
	customPool := &rest.CustomPool{
		MaxIdleConnsPerHost: 100,
	}
	restClientShop := &rest.RequestBuilder{
		BaseURL:        config.InternalBaseShopsClient,
		Timeout:        5 * time.Second,
		ContentType:    rest.JSON,
		EnableCache:    false,
		DisableTimeout: false,
		CustomPool:     customPool,
		FollowRedirect: true,
		//RetryStrategy:  retry.NewSimpleRetryStrategy(3, 30*time.Millisecond),
		MetricsConfig: rest.MetricsReportConfig{TargetId: "shops-api"},
	}
	return &ClientImpl{Client: restClientShop}
}

func (cl ClientImpl) GetShopByUserID(c *gin.Context) (dto.ShopIDResponse, apierrors.ApiError) {
	var responseBytes []byte

	response := cl.Client.Get(shopsBaseEndpoint, rest.Headers(c.Request.Header))

	if response.Response == nil || (response.StatusCode != http.StatusNotFound && response.StatusCode != http.StatusOK) {
		return dto.ShopIDResponse{}, apierrors.NewInternalServerApiError(fmt.Sprint("Unexpected error getting shop, url: "+shopsBaseEndpoint), response.Err)
	}

	responseBytes = response.Bytes()

	if response.StatusCode != http.StatusNotFound {
		responseBytes = response.Bytes()
	}

	if responseBytes != nil {
		shop := dto.ShopIDResponse{}

		if unmarshallError := json.Unmarshal(responseBytes, &shop); unmarshallError != nil {
			return dto.ShopIDResponse{}, apierrors.NewInternalServerApiError("Unexpected error unmarshalling shop json response. value: "+string(responseBytes), unmarshallError)
		}

		return shop, nil
	}

	return dto.ShopIDResponse{}, apierrors.NewNotFoundApiError("shop not found")
}
