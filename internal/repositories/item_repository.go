package repositories

import (
	"fmt"
	"rgb/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

// Just fetches — no decisions made here
func (r *ItemRepository) GetById(ctx *gin.Context, id *string) (*models.Item, error) {
	var item models.Item

	r.db.First(&item, *id)
	return &item, nil
}

// Just saves — no decisions made here
func (r *ItemRepository) Save(ctx *gin.Context, req *models.CreateItemRequest) (*models.Item, error) {
	var item models.Item
	fmt.Println("Test Save")
	return &item, nil
}
