package shops

import (
	"github.com/agustinrabini/go-toolkit/goutils/apierrors"
	"github.com/agustinrabini/api-items-project/internal/dto"
	"github.com/gin-gonic/gin"
)

type ClientMock struct {
	HandleGetShopByUserID func(c *gin.Context) (dto.ShopIDResponse, apierrors.ApiError)
}

func NewClientMock() ClientMock {
	return ClientMock{}
}

func (mock ClientMock) GetShopByUserID(c *gin.Context) (dto.ShopIDResponse, apierrors.ApiError) {
	if mock.HandleGetShopByUserID != nil {
		return mock.HandleGetShopByUserID(c)
	}
	return dto.ShopIDResponse{}, nil
}
