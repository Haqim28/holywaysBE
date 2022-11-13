package repositories

import (
	"holyways/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	GetTransactionByFund(ID int) ([]models.Transaction, error)
	GetTransactionByFundPending(ID int) ([]models.Transaction, error)
	GetTransactionByUser(ID int) ([]models.Transaction, error)
	DeleteTransaction(ID int, order models.Transaction) (models.Transaction, error)
	UpdateTransaction(status string, ID string) error
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error
	return transaction, err
}
func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction

	err := r.db.First(&transaction, ID).Error

	return transaction, err
}

func (r *repository) GetTransactionByFund(ID int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	// not yet using category relation, cause this step doesnt Belong to Many
	err := r.db.Where("fund_id=?", ID).Where("status='success'").Preload("UserFund").Preload("UserDonate").Preload("Fund.User").Find(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransactionByFundPending(ID int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	// not yet using category relation, cause this step doesnt Belong to Many
	err := r.db.Where("fund_id=?", ID).Where("status!='success'").Preload("UserFund").Preload("UserDonate").Preload("Fund.User").Find(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransactionByUser(ID int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	// not yet using category relation, cause this step doesnt Belong to Many
	err := r.db.Where("user_donate_id=?", ID).Preload("UserDonate").Preload("UserFund").Preload("Fund.User").Find(&transaction).Error

	return transaction, err
}

// func (r *repository) UpdateTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {
// 	err := r.db.Save(&transaction).Error
// 	return transaction, err
// }

func (r *repository) DeleteTransaction(ID int, transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Exec("SET FOREIGN_KEY_CHECKS=0;").Raw("DELETE FROM products WHERE id=?", ID).Scan(&transaction).Error
	return transaction, err
}

// Create UpdateTransaction method here ...
func (r *repository) UpdateTransaction(status string, ID string) error {
	var transaction models.Transaction
	r.db.Preload("User").First(&transaction, ID)

	// new status : pending
	// status : pending

	transaction.Status = status

	err := r.db.Save(&transaction).Error

	return err
}
