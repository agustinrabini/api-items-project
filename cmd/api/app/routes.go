package app

import (
	"github.com/agustinrabini/go-toolkit/goauth"
	"github.com/agustinrabini/api-items-project/cmd/api/dependencies"
	"github.com/agustinrabini/api-items-project/cmd/api/handlers"
	"github.com/gin-gonic/gin"
)

func MapUrlsToControllers(router *gin.Engine, h dependencies.HandlersStruct) {
	// Health
	health := handlers.NewHealthCheckerHandler()
	router.GET("/ping", health.Ping)

	// Items
	router.GET("/items", handlers.LoggerHandler("GetItemsByUserID"), h.Items.GetItemsByUserID)
	router.GET("/items/:id", handlers.LoggerHandler("GetItemByID"), h.Items.GetItemByID)
	router.GET("/items/shop/:id", handlers.LoggerHandler("GetItemsByShopID"), h.Items.GetItemsByShopID)
	router.GET("/items/shop/:id/category/:category_id", handlers.LoggerHandler("GetItemsByShopIDandCategoryID"), h.Items.GetItemsByShopCategoryID)
	router.POST("/items/list", handlers.LoggerHandler("GetItemsByIDs"), h.Items.GetItemsByIDs)
	router.POST("/items", handlers.LoggerHandler("CreateItem"), h.Items.CreateItem)
	router.DELETE("/items/:id", handlers.LoggerHandler("DeleteItem"), h.Items.DeleteItem)
	router.PUT("/items/:id", handlers.LoggerHandler("UpdateItem"), h.Items.UpdateItem)

	router.PUT("/items/category", goauth.PasswordMiddleware(), handlers.LoggerHandler("UpdateCategory"), h.Categories.Update)
	router.DELETE("/items/category/:id_category", goauth.PasswordMiddleware(), handlers.LoggerHandler("DeleteCategory"), h.Categories.Delete)
	router.POST("/items/category", goauth.PasswordMiddleware(), handlers.LoggerHandler("CreateCategory"), h.Categories.Create)
	router.GET("/items/category/:id_category", handlers.LoggerHandler("GetCategory"), h.Categories.Get)
	router.GET("/items/categories", handlers.LoggerHandler("GetAllCategories"), h.Categories.GetAllCategories)
}
