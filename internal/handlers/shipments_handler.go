package handlers

import (
	"log/slog"
	"net/http"
	"rgb/internal/models"
	"rgb/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ShipmentHandler struct {
	shipmentService *services.ShipmentService
	validate        *validator.Validate
	logger          *slog.Logger
}

func NewShipmentHandler(ss *services.ShipmentService, v *validator.Validate, l *slog.Logger) *ShipmentHandler {
	return &ShipmentHandler{
		shipmentService: ss,
		validate:        v,
		logger:          l,
	}
}

func (h *ShipmentHandler) CreateShipment(c *gin.Context) {
	var shipment models.CreateShipmentRequest

	if !BindAndValidate(c, h.validate, &shipment) {
		return
	}

	// Do something with shipment

	// TODO: shipment id is spoofed currently
	h.logger.Info("shipment created", "shipment_id", 0)
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (h *ShipmentHandler) UpdateShipment(c *gin.Context) {
	var shipment models.UpdateShipmentRequest

	if !BindAndValidate(c, h.validate, &shipment) {
		return
	}

	// Do something with shipment

	// TODO: shipment id is spoofed currently
	h.logger.Info("shipment updated", "shipment_id", 0)
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (h *ShipmentHandler) GetShipments(c *gin.Context) {
	// Return shipments

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
