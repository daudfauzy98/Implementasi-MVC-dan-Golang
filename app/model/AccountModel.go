package model

import (
	"fmt"

	"github.com/daudfauzy98/Implementasi-MVC-dan-Golang/app/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Account untuk model tabel pada database
type Account struct {
	ID            int    `gorm:"primary_key" json:"-"`
	IDAccount     string `gorm:"id_account, omitempty"`
	Name          string `gorm:"name"`
	Password      string `gorm:"password, omitempty"`
	AccountNumber int    `gorm:"account_number, omitempty"`
	Saldo         int    `gorm:"saldo"`
}

type AccountModel struct {
	DB *gorm.DB
}

func (model AccountModel) InsertNewAccount(account Account) (bool, error) {
	account.AccountNumber = utils.RangeInt(1000, 10000)
	account.Saldo = 0
	account.IDAccount = fmt.Sprintf("id:-%d", utils.RangeInt(10, 5000))

	result := model.DB.Create(&account)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (model AccountModel) GetAccountDetail(idAccount int) (bool, error, []Transaction, Account) {
	var transaction []Transaction
	var account Account

	result := model.DB.Model(&Transaction{}).Where("sender = ? OR recipient = ? ", idAccount, idAccount).Find(&transaction)
	fmt.Println(idAccount)
	fmt.Println(transaction)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, errors.Errorf("Account not found!"), []Transaction{}, Account{}
		}
		return false, result.Error, []Transaction{}, Account{}
	}

	result = model.DB.Where(&Account{AccountNumber: idAccount}).Find(&account)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, errors.Errorf("Account not found!"), []Transaction{}, Account{}
		}
		return false, result.Error, []Transaction{}, Account{}
	}

	return true, nil, transaction, account
}
