package handlers

import (
	"net/http"

	"github.com/agustinrabini/go-toolkit/goutils/apierrors"
	"github.com/agustinrabini/api-items-project/internal/categories"
	"github.com/agustinrabini/api-items-project/internal/domain"
	"github.com/agustinrabini/api-items-project/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type CategoriesHandler struct {
	Service categories.Service
}

func NewCategoriesHandler(service categories.Service) CategoriesHandler {
	return CategoriesHandler{Service: service}
}

// GetAllCategories godoc
// @Summary Get All Categories Items
// @Description Get All Categories Items
// @Tags Get All Categories Items
// @Accept  json
// @Produce  json
// @Success 200 {object} []domain.Category
// @Router /items/categories [get]
func (h CategoriesHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.Service.GetAllCategories(c)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, dto.CategoriesDTO{
		CategoryDTO: categories,
	},
	)
}

// CreateCategory godoc
// @Summary Create Category
// @Description Create Category Item in db
// @Tags Items
// @Accept  json
// @Produce  json
// @Param item body domain.Category true "Add Category Item"
// @Success 200
// @Router /items/categories [post]
func (h CategoriesHandler) Create(c *gin.Context) {
	var input domain.Category

	if err := binding.JSON.Bind(c.Request, &input); err != nil {
		apiErr := apierrors.NewBadRequestApiError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	if input.Name == "" {
		apiErr := apierrors.NewBadRequestApiError("category name must not be nil")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	id, err := h.Service.Create(c, input)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, id)
}

// GetCategory godoc
// @Summary Get Category
// @Description Get Category
// @Tags Get Category
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.Category
// @Router /items/category/:id_category [get]
func (h CategoriesHandler) Get(c *gin.Context) {

	categoryID := c.Param("id_category")
	if categoryID == "" {
		apiErr := apierrors.NewBadRequestApiError("id category is nil and should not")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	category, err := h.Service.Get(c, categoryID)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, category)
}

// DeleteCategory godoc
// @Summary Delete Category
// @Description Delete Category
// @Produce  json
// @Param id_category path string true "Category ID"
// @Success 200
// @Router /items/category/:id_category [delete]
func (h CategoriesHandler) Delete(c *gin.Context) {

	categoryID := c.Param("id_category")

	if categoryID == "" {
		apiErr := apierrors.NewBadRequestApiError("id category is nil and should not")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	categoryID, err := h.Service.Delete(c, categoryID)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, categoryID)
}

// UpdateCategory godoc
// @Summary Update Category Item
// @Description Update Category Item by ID
// @Tags Category
// @Accept  json
// @Produce  json
// @Param id_category path string true "Category ID"
// @Param item body domain.Category true "Update item"
// @Success 200
// @Router /items/category [put]
func (h CategoriesHandler) Update(c *gin.Context) {

	var input domain.Category

	if err := binding.JSON.Bind(c.Request, &input); err != nil {
		apiErr := apierrors.NewBadRequestApiError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	if input.Name == "" {
		apiErr := apierrors.NewBadRequestApiError("category name must not be nil")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	err := h.Service.Update(c, input)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
