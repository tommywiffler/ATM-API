package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID        string "json:id"
	FirstName string "json:firstName"
	LastName  string "json:lastName"
	AccNumber string "json:accnumber"
}

type Login struct {
	ID string "json:id"
}

func (u *User) UserLogin(db *gorm.DB, id string) (*User, error) {
	err := db.Debug().Where("id = ? ", id).Take(&u).Error
	return u, err
}
