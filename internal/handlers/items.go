package handlers

import (
	"log/slog"
	"net/http"
	"rgb/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ItemHandler struct {
	validate *validator.Validate
	logger   *slog.Logger
}

func NewItemHandler(v *validator.Validate, l *slog.Logger) *ItemHandler {
	return &ItemHandler{
		validate: v,
		logger:   l,
	}
}

func (h *ItemHandler) CreateItem(c *gin.Context) {
	var item models.CreateItemRequest
	if !BindAndValidate(c, h.validate, &item) {
		return
	}

	// Do something with item

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

func (h *ItemHandler) GetItems(c *gin.Context) {
	// Return items

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
