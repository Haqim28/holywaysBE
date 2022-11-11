package transactiondto

import "time"

type CreateTransactionRequest struct {
	FundID       int       `json:"fund_id" form:"fund_id"`
	UserDonateID int       `json:"user_donate_id" form:"user_donate_id"`
	UserFundID   int       `json:"user_fund_id" form:"user_fund_id"`
	Status       string    `json:"status" form:"status"`
	DonateAmount int       `json:"donate_amount"`
	CreatedAt    time.Time `json:"create_at"`
	UpdatedAt    time.Time `json:"-"`
}

type UpdateOrderRequest struct {
	Qty         int `json:"qty" form:"qty"`
	Price_order int `json:"price_order" form:"price_order"`
}
