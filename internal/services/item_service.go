package services

import (
	"log/slog"
	"rgb/internal/models"
	"rgb/internal/repositories"

	"github.com/gin-gonic/gin"
)

type ItemService struct {
	itemRepo *repositories.ItemRepository
	logger   *slog.Logger
}

func NewItemService(ir *repositories.ItemRepository, log *slog.Logger) *ItemService {
	return &ItemService{itemRepo: ir, logger: log}
}

func (s *ItemService) CreateItem(c *gin.Context, item *models.CreateItemRequest) (*models.Item, error) {
	// Do business rules here
	return s.itemRepo.Save(c, item)
}

func (s *ItemService) GetItemById(c *gin.Context, id *string) (*models.Item, error) {
	// Do business rules here
	return s.itemRepo.GetById(c, id)
}

// func (s *ItemService) UpdateItem(ctx *gin.Context, item models.Item) (*models.Item, error) {
//     // Business rule: validate species exists and fetch its data in one step
//     species, err := s.speciesRepo.GetByID(ctx, item.Species)
//     if err != nil {
//         return nil, ErrInvalidSpecies
//     }

//     // Business rule: check feeding schedule based on species data
//     if item.LastFed.Before(time.Now().Add(-species.MaxFeedingInterval)) {
//         s.logger.Warn("item is overdue for feeding", "item_id", item.ID)
//     }

//     // Business rule: weight sanity check against species norms
//     if item.Weight > species.MaxWeight {
//         return nil, fmt.Errorf("weight %.2f exceeds maximum for species %s", item.Weight, species.Name)
//     }

//     // Delegate the actual save to the repository
//     return s.itemRepo.Save(ctx, item)
// }
