package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"rgb/internal/handlers"
)

func setRouter(itemHandler *handlers.ItemHandler, shipmentHandler *handlers.ShipmentHandler) *gin.Engine {
	// Creates default gin router with Logger and Recovery middleware already attached
	router := gin.Default()

	router.RedirectTrailingSlash = true

	// Create API route group
	api := router.Group("/api")
	{
		// Test route to check if app is up and running
		api.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "world",
			})
		})
	}

	// Items route group logic is handled in internal/server/handlers/items
	items := router.Group("/items")
	{
		items.POST("/", itemHandler.CreateItem)
		items.PUT("/:id", itemHandler.UpdateItem)
		items.GET("/", itemHandler.GetItems)
	}

	// Shipments route group logic is handled in internal/server/handlers/shipments
	shipments := router.Group("/shipments")
	{
		shipments.POST("/", shipmentHandler.CreateShipment)
		shipments.PUT("/:id", shipmentHandler.UpdateShipment)
		shipments.GET("/", shipmentHandler.GetShipments)
	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}
