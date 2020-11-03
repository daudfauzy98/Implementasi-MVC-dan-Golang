package model

import (
	"fmt"

	"github.com/daudfauzy98/Implementasi-MVC-dan-Golang/app/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type AccountModel struct {
	DB            *gorm.DB
	ID            int    `gorm:"primary_key" json:"-"`
	IdAccount     string `gorm:"id_account, omitempty"`
	Name          string `gorm:"name"`
	Password      string `gorm:"password, omitempty"`
	AccountNumber int    `gorm:"account_number, omitempty"`
	Saldo         int    `gorm:"saldo"`
}

func (account AccountModel) InsertNewAccount() (bool, error) {
	account.AccountNumber = utils.RangeInt(1000, 10000)
	account.Saldo = 0
	account.IdAccount = fmt.Sprintf("id:-%d", utils.RangeInt(10, 5000))

	result := account.DB.Create(&account)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (account AccountModel) GetAccountDetail(idAccount int) (bool, error, []TransactionModel, AccountModel) {
	var transaction []TransactionModel

	result := account.DB.Where("sender = ? OR recipient = ? ", idAccount, idAccount).Find(&transaction)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, errors.Errorf("Account not found!"), []TransactionModel{}, AccountModel{}
		}
		return false, result.Error, []TransactionModel{}, AccountModel{}
	}

	result = account.DB.Where(&AccountModel{
		AccountNumber: idAccount,
	}).Find(&account)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, errors.Errorf("Account not found!"), []TransactionModel{}, AccountModel{}
		}
		return false, result.Error, []TransactionModel{}, AccountModel{}
	}
	return true, nil, transaction, account
}
