package models

type Donation struct {
	ID     int                  `json:"id"`
	FundID int                  `json:"fund_id"`
	Fund   Fund                 `json:"fund"`
	UserID int                  `json:"user_id"`
	User   UsersProfileResponse `json:"user"`
	Price  int                  `json:"price"`
}
