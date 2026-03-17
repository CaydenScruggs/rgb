package models

// Used for POST /shipments/createShipment — no ID, all fields required
type CreateShipmentRequest struct {
	ItemID      string `json:"itemID" validate:"required"`
	Origin      string `json:"origin" validate:"required"`
	Destination string `json:"destination" validate:"required"`
	Status      string `json:"status" validate:"required"`
	ShipDate    Date   `json:"shipDate" validate:"required"`
}

// Consider switching to PATCH over PUT

// Used for PUT /items/:id — ID comes from the URL, fields still required
type UpdateShipmentRequest struct {
	ItemID      string `json:"itemID" validate:"required"`
	Origin      string `json:"origin" validate:"required"`
	Destination string `json:"destination" validate:"required"`
	Status      string `json:"status" validate:"required"`
	ShipDate    Date   `json:"shipDate" validate:"required"`
}

// The actual database model — used internally, never exposed directly
type Shipment struct {
	ItemID      string
	Origin      string
	Destination string
	Status      string
	ShipDate    Date
}
