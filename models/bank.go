package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Account struct {
	Number  string `json:"number"`
	Balance int    `json:"balance"`
}

func (a *Account) GetBalance(db *gorm.DB, owner string) (*Account, error) {
	err := db.Debug().Where("number = ?", owner).Take(&a).Error
	return a, err
}

func (a *Account) Deposit(db *gorm.DB, owner string) (*Account, error) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var model Account
		accountToUpdate, err := model.GetBalance(tx, owner)
		if err != nil {
			return err
		}
		currBal := accountToUpdate.Balance
		newBal := currBal + a.Balance

		newAcc := Account{
			Number:  a.Number,
			Balance: newBal,
		}

		err = tx.Debug().Model(&newAcc).Where("number = ?", newAcc.Number).Update("balance", newAcc.Balance).Error
		return err
	})

	if err != nil {
		return nil, err
	}
	return a.GetBalance(db, owner)
}

func (a *Account) Withdraw(db *gorm.DB, owner string) (*Account, error) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var model Account
		accountToUpdate, err := model.GetBalance(tx, owner)
		if err != nil {
			return err
		}
		currBal := accountToUpdate.Balance
		newBal := currBal - a.Balance
		if newBal < 0 {
			return fmt.Errorf("cannot withdraw submitted amount, insufficient funds")
		}

		newAcc := Account{
			Number:  a.Number,
			Balance: newBal,
		}

		err = tx.Debug().Model(&newAcc).Where("number = ?", newAcc.Number).Update("balance", newAcc.Balance).Error
		return err
	})

	if err != nil {
		return nil, err
	}
	return a.GetBalance(db, owner)
}
