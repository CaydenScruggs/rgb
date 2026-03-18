package services

import (
	"log/slog"
	"rgb/internal/repositories"
)

type ShipmentService struct {
	shipmentRepo *repositories.ShipmentRepository
	logger       *slog.Logger
}

func NewShipmentService(sr *repositories.ShipmentRepository, log *slog.Logger) *ShipmentService {
	return &ShipmentService{shipmentRepo: sr, logger: log}
}
