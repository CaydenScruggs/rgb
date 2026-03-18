package handlers

import (
	"log/slog"
	"net/http"
	"rgb/internal/models"
	"rgb/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ItemHandler struct {
	itemService *services.ItemService
	validate    *validator.Validate
	logger      *slog.Logger
}

func NewItemHandler(is *services.ItemService, v *validator.Validate, l *slog.Logger) *ItemHandler {
	return &ItemHandler{
		itemService: is,
		validate:    v,
		logger:      l,
	}
}

func (h *ItemHandler) CreateItem(c *gin.Context) {
	var item models.CreateItemRequest
	if !BindAndValidate(c, h.validate, &item) {
		return
	}

	// Do something with item
	h.itemService.CreateItem(c, &item)

	// TODO: item id is spoofed currently
	h.logger.Info("item created", "item_id", 0)
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (h *ItemHandler) UpdateItem(c *gin.Context) {
	var item models.UpdateItemRequest

	if !BindAndValidate(c, h.validate, &item) {
		return
	}

	// Do something with item

	// TODO: item id is spoofed currently
	h.logger.Info("item updated", "item_id", 0)
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (h *ItemHandler) GetItemById(c *gin.Context) {
	// Return item by ID

	id := c.Param("id")
	h.itemService.GetItemById(c, &id)
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
