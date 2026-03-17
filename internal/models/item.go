package models

// In the future, consider adding last shed, last drink, etc

// Used for POST /items — no ID, all fields required
type CreateItemRequest struct {
	Name    string  `json:"name"    validate:"required"`
	Species *uint8  `json:"species" validate:"required"`
	Age     *uint8  `json:"age"     validate:"required"`
	Weight  float32 `json:"weight"  validate:"required,gt=0"`
	LastFed *Date   `json:"lastFed" validate:"required"`
}

// Consider switching to PATCH over PUT

// Used for PUT /items/:id — ID comes from the URL, fields still required
type UpdateItemRequest struct {
	Name    string  `json:"name"    validate:"required"`
	Species *uint8  `json:"species" validate:"required"`
	Age     *uint8  `json:"age"     validate:"required"`
	Weight  float32 `json:"weight"  validate:"required,gt=0"`
	LastFed *Date   `json:"lastFed" validate:"required"`
}

// The actual database model — used internally, never exposed directly
type Item struct {
	ID      string
	Name    string
	Species uint8
	Age     uint8
	Weight  float32
	LastFed Date
}
