package donationdto

type CreateDonationRequest struct {
	FundID int `json:"fund_id" form:"fund_id"`
	UserID int `json:"user_id" form:"user_id"`
	Price  int `json:"price"`
}

type UpdateOrderRequest struct {
	Qty         int `json:"qty" form:"qty"`
	Price_order int `json:"price_order" form:"price_order"`
}
