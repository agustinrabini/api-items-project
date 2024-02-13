package handlers

import (
	"net/http"

	"github.com/agustinrabini/go-toolkit/goauth"
	"github.com/agustinrabini/go-toolkit/goutils/apierrors"
	"github.com/agustinrabini/go-toolkit/goutils/logger"
	"github.com/agustinrabini/api-items-project/internal/dto"

	"github.com/agustinrabini/api-items-project/internal/domain"
	"github.com/agustinrabini/api-items-project/internal/items"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var ErrorNoItemID = apierrors.NewApiError("missing item_id in the requests", "", 401, nil)
var ErrorNoShopid = apierrors.NewApiError("missing shop_id in the parameters", "", 422, nil)
var ErrorNoCategory = apierrors.NewApiError("missing category_id in the parameters", "", 422, nil)

type ItemsHandler struct {
	Service items.Service
}

func NewItemsHandler(service items.Service) ItemsHandler {
	return ItemsHandler{Service: service}
}

// GetItemByID godoc
// @Summary Get details of item id
// @Description Get details of item
// @Tags Items
// @Accept  json
// @Produce  json
// @Param id path string true "Item ID"
// @Success 200 {object} domain.Item
// @Router /items/{id} [get]
func (h ItemsHandler) GetItemByID(c *gin.Context) {
	itemID := c.Param("id")
	if itemID == "" {
		c.JSON(ErrorNoItemID.Status(), ErrorNoItemID)
		return
	}
	item, err := h.Service.Get(c, itemID)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, item)
}

// GetItemsByUserID godoc
// @Summary Get Items by User ID
// @Description Get Items by User ID in headers
// @Tags Items
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.ItemsOutput
// @Router /items [get]
func (h ItemsHandler) GetItemsByUserID(c *gin.Context) {
	userID, err1 := goauth.GetUserId(c)
	if err1 != nil {
		apiErr := apierrors.NewUnauthorizedApiError(err1.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	itemsResponse, err := h.Service.GetItemsByUserID(c, userID)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, itemsResponse)
}

// GetItemsByShopID godoc
// @Summary Get details of items by shop ID
// @Description Get details of items
// @Tags Items
// @Accept  json
// @Produce  json
// @Param id path string true "Shop ID"
// @Success 200 {object} domain.ItemsOutput
// @Router /items/shop/{id} [get]
func (h ItemsHandler) GetItemsByShopID(c *gin.Context) {
	shopID := c.Param("id")
	if shopID == "" {
		c.JSON(ErrorNoShopid.Status(), ErrorNoShopid)
		return
	}
	itemsResponse, err := h.Service.GetItemsByShopID(c, shopID)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, itemsResponse)
}

// GetItemsByShopCategoryID godoc
// @Summary Get details of items by shop ID and category ID
// @Description Get details of items
// @Tags Items
// @Accept  json
// @Produce  json
// @Param id path string true "Shop ID"
// @Param category_id path string true "Category ID"
// @Success 200 {object} domain.ItemsOutput
// @Router /items/shop/{id}/category/{category_id} [get]
func (h ItemsHandler) GetItemsByShopCategoryID(c *gin.Context) {
	shopID := c.Param("id")
	if shopID == "" {
		c.JSON(ErrorNoShopid.Status(), ErrorNoShopid)
		return
	}
	categoryID := c.Param("category_id")
	if categoryID == "" {
		c.JSON(ErrorNoShopid.Status(), ErrorNoShopid)
		return
	}
	itemsResponse, err := h.Service.GetItemsByShopCategoryID(c, shopID, categoryID)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	if len(itemsResponse.Items) == 0 {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(http.StatusOK, itemsResponse)
}

// GetItemsByIDs godoc
// @Summary Get items by ids
// @Description Get item by IDs in body
// @Tags Items
// @Accept  json
// @Produce  json
// @Param items body domain.ItemsIds true "Add items"
// @Success 200 {object} domain.ItemsOutput
// @Router /items/list [post]
func (h ItemsHandler) GetItemsByIDs(c *gin.Context) {
	var itemsInput domain.ItemsIds
	if err := binding.JSON.Bind(c.Request, &itemsInput); err != nil {
		apiErr := apierrors.NewGenericErrorMessageDecoder(err)
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	itemsResponse, err := h.Service.GetItemsByIDs(c, itemsInput)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, itemsResponse)
}

// CreateItem godoc
// @Summary Create Item
// @Description Create item in db
// @Tags Items
// @Accept  json
// @Produce  json
// @Param item body dto.ItemDTO true "Add item"
// @Success 201
// @Router /items [post]
func (h ItemsHandler) CreateItem(c *gin.Context) {
	var itemInput dto.ItemDTO
	if err := binding.JSON.Bind(c.Request, &itemInput); err != nil {
		apiErr := apierrors.NewGenericErrorMessageDecoder(err)
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	userID, err1 := goauth.GetUserId(c)
	if err1 != nil {
		apiErr := apierrors.NewUnauthorizedApiError(err1.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	itemInput.UserID = userID
	insertedID, err := h.Service.CreateItem(c, itemInput)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, insertedID)
}

// DeleteItem godoc
// @Summary Delete item
// @Description Delete item by ID
// @Tags Items
// @Accept  json
// @Produce  json
// @Param id path string true "Item ID"
// @Success 204
// @Router /items/{id} [delete]
func (h ItemsHandler) DeleteItem(c *gin.Context) {
	itemID := c.Param("id")
	if itemID == "" {
		c.JSON(ErrorNoItemID.Status(), ErrorNoItemID)
		return
	}
	err := h.Service.Delete(c, itemID)
	if err != nil {
		logger.Error("error delete items by id ", err)
		c.JSON(err.Status(), err)
		return
	}
	c.Status(http.StatusNoContent)
}

// UpdateItem godoc
// @Summary Update item
// @Description Update item by ID
// @Tags Items
// @Accept  json
// @Produce  json
// @Param id path string true "Item ID"
// @Param item body dto.ItemDTO true "Add item"
// @Success 200
// @Router /items/{id} [put]
func (h ItemsHandler) UpdateItem(c *gin.Context) {

	var itemInput dto.ItemDTO

	itemID := c.Param("id")
	if itemID == "" {
		c.JSON(ErrorNoItemID.Status(), ErrorNoItemID)
		return
	}

	if err := binding.JSON.Bind(c.Request, &itemInput); err != nil {
		apiErr := apierrors.NewGenericErrorMessageDecoder(err)
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	userID, err1 := goauth.GetUserId(c)
	if err1 != nil {
		apiErr := apierrors.NewUnauthorizedApiError(err1.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	itemInput.UserID = userID
	err := h.Service.Update(c, itemID, itemInput)
	if err != nil {
		logger.Error("error update items by id ", err)
		c.JSON(err.Status(), err)
		return
	}
	c.Status(http.StatusOK)
}
