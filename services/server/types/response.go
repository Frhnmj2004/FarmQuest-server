package types

import (
	"time"
)

// auth
type LoginResponse struct {
	Token string `json:"token"`
}
type RegisterResponse struct {
	Token string `json:"token"`
}

// crop
type GetSimpleCropResponse struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	CroppedImageURL string `json:"cropped_image_url"`
	FullImageURL    string `json:"full_image_url"`
}
type GetCropResponse struct {
	ID           uint     `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	FullImageURL string   `json:"full_image_url"`
	BasicNeeds   []string `json:"basic_needs"`
	Tags         []string `json:"tags"`
	Price        int64    `json:"price"`
	Rating       int64    `json:"rating"`
}

// farm
type GetSimpleFarmResponse struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Image    string  `json:"image"`
	Status   string  `json:"status"`
	Location string  `json:"location"`
	Area     float64 `json:"area"`
}

type GetFarmResponse struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Image        string    `json:"image"`
	Status       string    `json:"status"`
	Description  string    `json:"description"`
	Location     string    `json:"location"`
	Health       int       `json:"health"`
	Area         float64   `json:"area"`
	PlantedAt    time.Time `json:"planted_at"`
	GrowingAt    time.Time `json:"growing_at,omitempty"`
	HarvestAt    time.Time `json:"harvest_at,omitempty"`
	GrowthStatus struct {
		Description string `json:"description"`
		Image       string `json:"image_url"`
	} `json:"growth_status"`
	RelatedNews []struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Image       string `json:"image_url"`
		Link        string `json:"link"`
	} `json:"related_news"`
}

// growth status
type GetGrowthStatusResponse struct {
	Stages []struct {
		Description string    `json:"description"`
		Status      string    `json:"status"`
		Date        time.Time `json:"date,omitempty"`
	} `json:"stages"`
	//GrowthTip struct {
	//Message string `json:"message"`
	//Image   string `json:"image_url,omitempty"`
	//} `json:"growth_tip"`
}

// questions
type GetQuestionsResponse struct {
	Text    string `json:"text"`
	Answers []struct {
		Text string `json:"text"`
	} `json:"answers"`
}

// coins
type CoinsResponse struct {
	Coins int `json:"coins"`
}
type RewardResponse struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	TaskType     string    `json:"task_type"`
	PointsReward int       `json:"points_reward"`
	TimeSlot     string    `json:"time_slot"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	CompletedAt  time.Time `json:"completed_at,omitempty"`
}

// NewsResponse represents a single news item
type GetSimpleNewsResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	AuthorID    uint      `json:"author_id"`
	PublishedAt time.Time `json:"published_at"`
}

type UserProfile struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	AvatarURL string    `json:"avatar_url"`
	Role      string    `json:"role"`
	Points    int       `json:"points"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

// NewsListResponse lists multiple news items
// // GeneralResponse is a standard response wrapper for all API endpoints
// type GeneralResponse struct {
// 	Message string      `json:"message"`
// 	Data    interface{} `json:"data,omitempty"`
// 	Error   string      `json:"error,omitempty"`
// }

// // UserResponse represents the response for user-related endpoints
// type UserResponse struct {
// 	ID        uint      `json:"id"`
// 	Username  string    `json:"username"`
// 	Email     string    `json:"email"`
// 	Role      string    `json:"role"`
// 	Points    int       `json:"points"`
// 	Balance   float64   `json:"balance"`
// 	CreatedAt time.Time `json:"created_at"`
// }

// // RegisterResponse includes the token for the newly registered user
// type RegisterResponse struct {
// 	GeneralResponse
// 	Data UserResponse `json:"data"`
// }

// // LoginResponse includes the JWT token for the authenticated user
// type LoginResponse struct {
// 	GeneralResponse
// 	Data struct {
// 		Token string       `json:"token"`
// 		User  UserResponse `json:"user"`
// 	} `json:"data"`
// }

// // ProfileResponse includes the user's profile details
// type ProfileResponse struct {
// 	GeneralResponse
// 	Data struct {
// 		FullName  string `json:"full_name"`
// 		Address   string `json:"address"`
// 		Phone     string `json:"phone"`
// 		AvatarURL string `json:"avatar_url"`
// 		Bio       string `json:"bio"`
// 	} `json:"data"`
// }

// // CropResponse represents a crop in the list view (home page)
// type CropResponse struct {
// 	ID         uint   `json:"id"`
// 	Name       string `json:"name"`
// 	Type       string `json:"type"`     // e.g., vegetable, fruit, herb
// 	Category   string `json:"category"` // e.g., indoor, outdoor, popular
// 	ImageURL   string `json:"image_url"`
// 	IsFavorite bool   `json:"is_favorite"` // Whether the user favorited this crop
// }

// // OfferResponse represents a personalized offer (banner on home page)
// type OfferResponse struct {
// 	CropID     uint    `json:"crop_id"`
// 	CropName   string  `json:"crop_name"`
// 	Discount   float64 `json:"discount"`    // e.g., 40 for 40% off
// 	StartDate  string  `json:"start_date"`  // e.g., "2 Jul"
// 	EndDate    string  `json:"end_date"`    // e.g., "20 July"
// 	BannerText string  `json:"banner_text"` // e.g., "40% off"
// }

// // CropDetailsResponse represents detailed crop information (details page)
// type CropDetailsResponse struct {
// 	ID           uint    `json:"id"`
// 	Name         string  `json:"name"`
// 	Type         string  `json:"type"`
// 	ImageURL     string  `json:"image_url"`
// 	Rating       float64 `json:"rating"`       // e.g., 4.8
// 	ReviewCount  int     `json:"review_count"` // e.g., 200
// 	Description  string  `json:"description"`
// 	WaterNeed    string  `json:"water_need"`    // e.g., "400ml"
// 	SunlightNeed string  `json:"sunlight_need"` // e.g., "Direct"
// 	Price        float64 `json:"price"`         // e.g., 129.99
// 	IsFavorite   bool    `json:"is_favorite"`
// }

// // FarmResponse represents a farm plot response
// type FarmResponse struct {
// 	ID          uint      `json:"id"`
// 	UserID      uint      `json:"user_id"`
// 	CropID      uint      `json:"crop_id"`
// 	Name        string    `json:"name"`
// 	Status      string    `json:"status"` // planted, growing, harvesting, completed
// 	PlantedAt   time.Time `json:"planted_at"`
// 	HarvestedAt time.Time `json:"harvested_at"`
// 	Yield       float64   `json:"yield"`
// }

// // FarmsResponse lists multiple farm plots
// type FarmsResponse struct {
// 	GeneralResponse
// 	Data []FarmResponse `json:"data"`
// }

// // StartFarmingResponse confirms the start of a farming plot
// type StartFarmingResponse struct {
// 	GeneralResponse
// 	Data FarmResponse `json:"data"`
// }

// // OrderResponse represents an order response
// type OrderResponse struct {
// 	ID        uint      `json:"id"`
// 	UserID    uint      `json:"user_id"` // Seller
// 	BuyerID   uint      `json:"buyer_id"`
// 	FarmID    uint      `json:"farm_id"`
// 	Quantity  float64   `json:"quantity"`
// 	Price     float64   `json:"price"`
// 	Status    string    `json:"status"` // pending, shipped, delivered, cancelled
// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// }

// // OrdersResponse lists multiple orders
// type OrdersResponse struct {
// 	GeneralResponse
// 	Data []OrderResponse `json:"data"`
// }

// // RewardResponse represents a reward response
// type RewardResponse struct {
// 	ID          uint      `json:"id"`
// 	UserID      uint      `json:"user_id"`
// 	Amount      int       `json:"amount"`
// 	Description string    `json:"description"`
// 	RewardType  string    `json:"reward_type"`
// 	CreatedAt   time.Time `json:"created_at"`
// }

// // RewardsResponse lists multiple rewards
// type RewardsResponse struct {
// 	GeneralResponse
// 	Data []RewardResponse `json:"data"`
// }

// // NewsResponse represents a news item response
// type NewsResponse struct {
// 	ID          uint      `json:"id"`
// 	Title       string    `json:"title"`
// 	Content     string    `json:"content"`
// 	AuthorID    uint      `json:"author_id"`
// 	PublishedAt time.Time `json:"published_at"`
// }

// // NewsListResponse lists multiple news items
// type NewsListResponse struct {
// 	GeneralResponse
// 	Data []NewsResponse `json:"data"`
// }

// // AlertResponse represents an alert response
// type AlertResponse struct {
// 	ID        uint      `json:"id"`
// 	UserID    uint      `json:"user_id"`
// 	Message   string    `json:"message"`
// 	Type      string    `json:"type"` // info, warning, success
// 	IsRead    bool      `json:"is_read"`
// 	CreatedAt time.Time `json:"created_at"`
// }

// // AlertsResponse lists multiple alerts
// type AlertsResponse struct {
// 	GeneralResponse
// 	Data []AlertResponse `json:"data"`
// }

// // BadgeResponse represents a badge response
// type BadgeResponse struct {
// 	ID             uint   `json:"id"`
// 	Name           string `json:"name"`
// 	Description    string `json:"description"`
// 	IconURL        string `json:"icon_url"`
// 	PointsRequired int    `json:"points_required"`
// }

// // BadgesResponse lists multiple badges
// type BadgesResponse struct {
// 	GeneralResponse
// 	Data []BadgeResponse `json:"data"`
// }

// // UserBadgesResponse lists badges earned by a user
// type UserBadgesResponse struct {
// 	GeneralResponse
// 	Data []struct {
// 		BadgeID  uint          `json:"badge_id"`
// 		EarnedAt time.Time     `json:"earned_at"`
// 		Badge    BadgeResponse `json:"badge"`
// 	} `json:"data"`
// }

// // TaskResponse represents a task response
// type TaskResponse struct {
// 	ID           uint      `json:"id"`
// 	FarmID       uint      `json:"farm_id"`
// 	TaskType     string    `json:"task_type"`
// 	PointsReward int       `json:"points_reward"`
// 	Status       string    `json:"status"`
// 	DueAt        time.Time `json:"due_at"`
// 	CompletedAt  time.Time `json:"completed_at"`
// }

// TasksResponse lists multiple tasks
// type TasksResponse struct {
// 	GeneralResponse
// 	Data []TaskResponse `json:"data"`
// }

// type RegisterRequest struct {
// 	Username string `json:"username"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }
