package server

import (
	"rgb/internal/handlers"
)

func Start(itemHandler *handlers.ItemHandler, shipmentHandler *handlers.ShipmentHandler) {
	router := setRouter(itemHandler, shipmentHandler)

	// Start listening and serving requests
	router.Run(":8080")
}
