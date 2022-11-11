package repositories

import (
	"holyways/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(order models.Transaction) (models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	GetTransactionByFund(ID int) ([]models.Transaction, error)
	GetTransactionByUser(ID int) ([]models.Transaction, error)
	UpdateTransaction(Transaction models.Transaction, ID int) (models.Transaction, error)
	DeleteTransaction(ID int, order models.Transaction) (models.Transaction, error)
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

	err := r.db.Preload("Product").Preload("Cart").First(&transaction, ID).Error

	return transaction, err
}

func (r *repository) GetTransactionByFund(ID int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	// not yet using category relation, cause this step doesnt Belong to Many
	err := r.db.Where("fund_id=?", ID).Preload("UserFund").Preload("UserDonate").Preload("Fund.User").Find(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransactionByUser(ID int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	// not yet using category relation, cause this step doesnt Belong to Many
	err := r.db.Where("user_donate_id=?", ID).Preload("UserDonate").Preload("UserFund").Preload("Fund.User").Find(&transaction).Error

	return transaction, err
}

func (r *repository) UpdateTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Save(&transaction).Error
	return transaction, err
}

func (r *repository) DeleteTransaction(ID int, transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Error
	return transaction, err
}
