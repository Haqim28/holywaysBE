package database

import (
	"holyways/models"
	"holyways/pkg/mysql"

	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Fund{},
		&models.Donation{},
		&models.Transaction{},
		// 	&models.Order{},
		//
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
