package types

// Request types
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=255"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type CreateFarmRequest struct {
	Name        string  `json:"name" binding:"required,max=255"`
	CropID      uint    `json:"crop_id" binding:"required,gt=0"`
	Description string  `json:"description" binding:"required,max=1000"`
	Location    string  `json:"location" binding:"required,max=255"`
	Area        float64 `json:"area" binding:"required,gt=0"`
}

type UpdateFarmRequest struct {
	Name        string  `json:"name,omitempty" binding:"omitempty,max=255"`
	Status      string  `json:"status,omitempty"`
	Health      int     `json:"health,omitempty" binding:"omitempty,gt=0,lte=100"`
	Area        float64 `json:"area,omitempty" binding:"omitempty,gt=0"`
	Description string  `json:"description,omitempty" binding:"omitempty,max=1000"`
	Location    string  `json:"location,omitempty" binding:"omitempty,max=255"`
}

type CreateOrderRequest struct {
	BuyerID  uint    `json:"buyer_id" binding:"required,gt=0"`
	FarmID   uint    `json:"farm_id" binding:"required,gt=0"`
	Quantity float64 `json:"quantity" binding:"required,gt=0"`
}

type Query struct {
	Page   int    `json:"page" binding:"gte=1"`
	Size   int    `json:"size" binding:"gte=1,lte=100"`
	Search string `json:"search"`
}

type UpdateProfileRequest struct {
	Username  string  `json:"username"`
	FullName  string  `json:"full_name"`
	AvatarURL string  `json:"avatar_url"`
	Balance   float64 `json:"balance"`
}
