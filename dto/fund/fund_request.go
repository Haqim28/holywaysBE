package funddto

type CreateFundRequest struct {
	Title       string `json:"title" form:"title"  gorm:"type: varchar(255)"`
	Image       string `json:"image" form:"image"  gorm:"type: varchar(255)"`
	Goal        int    `json:"goal" form:"goal"`
	Description string `json:"description" form:"description"  gorm:"type: varchar(255)"`
	UserID      int    `json:"user_id" form:"user_id"  gorm:"type: varchar(255)"`
}

type UpdateFundRequest struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Image       string `json:"image" form:"image" validate:"required"`
	Goal        int    `json:"goal" form:"goal"`
	Description string `json:"description" form:"description" validate:"required"`
	UserID      int    `json:"user_id" form:"user_id" validate:"required"`
}
