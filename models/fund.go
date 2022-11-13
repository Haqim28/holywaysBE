package models

type Fund struct {
	ID          int                  `json:"id"`
	Title       string               `json:"title" gorm:"type: varchar(255)"`
	Image       string               `json:"image" gorm:"type: varchar(255)"`
	Goal        int                  `json:"goal" gorm:"type: int"`
	Status      string               `json:"status" gorm:"type: varchar(255)"`
	Description string               `json:"description" gorm:"type: varchar(255)"`
	UserID      int                  `json:"user_id" gorm:"type: int"`
	User        UsersProfileResponse `json:"user"`
	UserDonate  []Transaction        `json:"donation"`
}
