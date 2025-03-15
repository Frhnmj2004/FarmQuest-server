package types

// Request types
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateFarmRequest struct {
	Name   string `json:"name" binding:"required"`
	CropID uint   `json:"crop_id" binding:"required"`
}

type UpdateFarmRequest struct {
	Name   string `json:"name"`
	Status string `json:"status"` // e.g., planted, growing, harvesting
}

type CreateOrderRequest struct {
	BuyerID  uint    `json:"buyer_id" binding:"required"`
	FarmID   uint    `json:"farm_id" binding:"required"`
	Quantity float64 `json:"quantity" binding:"required"`
}

type Query struct {
	Page   int    `json:"page"`
	Size   int    `json:"size"`
	Search string `json:"search"`
}
