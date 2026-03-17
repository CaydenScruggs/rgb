package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setRouter() *gin.Engine {
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

	items := router.Group("/items")
	{
		// Route to handle adding new and updating existing items to the system, logic is in the handler
		items.POST("/updateItem", updateItem)

		// Route to handle getting existing items from the system, logic is in the handler
		items.GET("/getItems", getItems)
	}

	shipments := router.Group("/shipments")
	{
		// Route to handle adding new and updating existing shipments from the system, logic is in the handler
		shipments.GET("/getShipments", updateShipment)

		// Route to handle getting existing shipments from the system, logic is in the handler
		shipments.GET("/getShipments", getShipments)
	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}
