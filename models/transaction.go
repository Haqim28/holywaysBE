package models

import "time"

type Transaction struct {
	ID           int                  `json:"id"`
	FundID       int                  `json:"fund_id"`
	Fund         Fund                 `json:"fund"`
	UserFundID   int                  `json:"user_fund_id"`
	UserFund     UsersProfileResponse `json:"user_fund" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserDonateID int                  `json:"user_donate_id"`
	UserDonate   UsersProfileResponse `json:"user_donate" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status       string               `json:"status"`
	DonateAmount int                  `json:"donate_amount"`
	CreatedAt    time.Time            `json:"create_at"`
	UpdatedAt    time.Time            `json:"-"`
}
