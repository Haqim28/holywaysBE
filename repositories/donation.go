package repositories

import (
	"holyways/models"

	"gorm.io/gorm"
)

type DonationRepository interface {
	CreateDonation(order models.Donation) (models.Donation, error)
	GetDonation(ID int) (models.Donation, error)
	GetDonationByFund(ID int) ([]models.Donation, error)
	GetDonationByUser(ID int) ([]models.Donation, error)
	UpdateDonation(Donation models.Donation, ID int) (models.Donation, error)
	DeleteDonation(ID int, order models.Donation) (models.Donation, error)
}

func RepositoryDonation(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateDonation(donation models.Donation) (models.Donation, error) {
	err := r.db.Create(&donation).Error
	return donation, err
}
func (r *repository) GetDonation(ID int) (models.Donation, error) {
	var donation models.Donation

	err := r.db.Preload("Product").Preload("Cart").First(&donation, ID).Error

	return donation, err
}

func (r *repository) GetDonationByFund(ID int) ([]models.Donation, error) {
	var donation []models.Donation
	// not yet using category relation, cause this step doesnt Belong to Many
	err := r.db.Where("fund_id=?", ID).Preload("User").Preload("Fund.User").Find(&donation).Error

	return donation, err
}

func (r *repository) GetDonationByUser(ID int) ([]models.Donation, error) {
	var donation []models.Donation
	// not yet using category relation, cause this step doesnt Belong to Many
	err := r.db.Where("user_id=?", ID).Preload("User").Preload("Fund.User").Find(&donation).Error

	return donation, err
}

func (r *repository) UpdateDonation(donation models.Donation, ID int) (models.Donation, error) {
	err := r.db.Save(&donation).Error
	return donation, err
}

func (r *repository) DeleteDonation(ID int, donation models.Donation) (models.Donation, error) {
	err := r.db.Delete(&donation).Error
	return donation, err
}
