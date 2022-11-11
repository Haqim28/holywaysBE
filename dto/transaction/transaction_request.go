package transaction

type CreateTransaction struct {
	Status       string `json:"status" form:"status" gorm:"type: varchar(255)"`
	FundID       int    `json:"fund_id" form:"fund_id" gorm:"type: int"`
	UserFundID   int    `json:"userfund_id" form:"userfund_id" gorm:"type: int"`
	UserDonateID int    `json:"userdonate_id" form:"userdonate_id" gorm:"type:int"`
}

type UpdateTransactionRequest struct {
	Status string `json:"status" form:"status" gorm:"type: varchar(255)"`
}
