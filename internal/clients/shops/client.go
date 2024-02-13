package shops

import (
	"github.com/agustinrabini/go-toolkit/goutils/apierrors"
	"github.com/agustinrabini/api-items-project/cmd/api/config"
	"github.com/agustinrabini/api-items-project/internal/dto"
	"github.com/gin-gonic/gin"
)

var (
	ClientInstance = newShopsApiClient()
)

type Client interface {
	GetShopByUserID(c *gin.Context) (dto.ShopIDResponse, apierrors.ApiError)
}

func newShopsApiClient() Client {
	if config.IsProductionEnvironment() {
		return NewClientImpl()
	}
	return NewClientMock()
}
