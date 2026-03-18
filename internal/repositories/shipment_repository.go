package repositories

import (
	"gorm.io/gorm"
)

type ShipmentRepository struct {
	db *gorm.DB
}

func NewShipmentRepository(db *gorm.DB) *ShipmentRepository {
	return &ShipmentRepository{db: db}
}
