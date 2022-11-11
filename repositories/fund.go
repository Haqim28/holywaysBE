package repositories

import (
	"holyways/models"

	"gorm.io/gorm"
)

type FundRepository interface {
	FindFunds() ([]models.Fund, error)
	CreateFund(fund models.Fund) (models.Fund, error)
	GetFund(ID int) (models.Fund, error)
	GetFundByUser(ID int) ([]models.Fund, error)
	UpdateFund(Fund models.Fund, ID int) (models.Fund, error)
}

func RepositoryFund(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindFunds() ([]models.Fund, error) {
	var funds []models.Fund
	err := r.db.Preload("User").Preload("UserDonate").Find(&funds).Error
	return funds, err
}

func (r *repository) CreateFund(fund models.Fund) (models.Fund, error) {
	err := r.db.Create(&fund).Error
	return fund, err
}

func (r *repository) GetFund(ID int) (models.Fund, error) {
	var fund models.Fund

	err := r.db.Preload("User").Preload("UserDonate").First(&fund, ID).Error

	return fund, err
}

func (r *repository) GetFundByUser(ID int) ([]models.Fund, error) {
	var funds []models.Fund

	err := r.db.Where("user_id=?", ID).Find(&funds).Error

	return funds, err
}

func (r *repository) UpdateFund(fund models.Fund, ID int) (models.Fund, error) {
	err := r.db.Save(&fund).Error

	return fund, err
}
