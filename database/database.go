package database

import (
	"atm-api/models"
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Printf("We are connected to the database")

	db, _ := DB.DB()
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	migrateTables(DB)
	return popTables(DB)

}

func migrateTables(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Account{},
		&models.User{},
	)
}

func popTables(db *gorm.DB) error {
	var accounts = []models.Account{{Owner: "1234", Balance: 5000}, {Owner: "4141", Balance: 350}}
	var users = []models.User{{ID: "1234", FirstName: "Peter", LastName: "Parker"}, {ID: "4141", FirstName: "Bruce", LastName: "Wayne"}}

	err := db.Create(&accounts).Error
	if err != nil {
		return err
	}
	err = db.Create(&users).Error
	if err != nil {
		return err
	}

	return nil
}
