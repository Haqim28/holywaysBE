package models

import "time"

type Transaction struct {
	ID           int                  `json:"id" gorm:"primary_key:auto_increment"`
	UserFundID   int                  `json:"userfund_id"`
	UserFund     UsersProfileResponse `json:"userfund" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserDonateID int                  `json:"userdonate_id"`
	UserDonate   UsersProfileResponse `json:"userdonate" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status       string               `json:"status"`
	Subtotal     int                  `json:"subtotal"`
	FundID       int                  `json:"fund_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Fund         Fund                 `json:"fund"`
	CreatedAt    time.Time            `json:"create_at"`
	UpdatedAt    time.Time            `json:"-"`
}
