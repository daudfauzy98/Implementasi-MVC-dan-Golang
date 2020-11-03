package model

import (
	"time"

	"github.com/daudfauzy98/Implementasi-MVC-dan-Golang/app/constant"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type TransactionModel struct {
	DB                     *gorm.DB
	ID                     int    `gorm: "primary_key" json:"-"`
	TransactionType        int    `json:"transaction_type, omitempty"`
	TransactionDescription string `json:"transaction_description"`
	Sender                 int    `json:"sender"`
	Amount                 int    `json:"amount"`
	Recipient              int    `json:"recipient"`
	Timestamp              int64  `json:"timestamp, omitempty"`
}

func (transaction TransactionModel) Transfer() (bool, error) {
	err := transaction.DB.Transaction(func(tx *gorm.DB) error {
		var sender, recipient AccountModel

		result := tx.Model(&AccountModel{}).Where(&AccountModel{
			AccountNumber: transaction.Sender,
		}).First(&sender)
		if result.Error != nil {
			return result.Error
		}

		// Check balance first
		if sender.Saldo < transaction.Amount {
			return errors.Errorf("Insufficient saldo")
		}

		// Jika saldo pengirim mencukupi
		result = result.Update("saldo", sender.Saldo-transaction.Amount)
		if result.Error != nil {
			return result.Error
		}

		result = tx.Model(&AccountModel{}).Where(AccountModel{
			AccountNumber: transaction.Recipient,
		}).First(&recipient).Update("saldo", recipient.Saldo+transaction.Amount)

		if result.Error != nil {
			return result.Error
		}

		transaction.TransactionType = constant.TRANSFER
		transaction.Timestamp = time.Now().Unix()
		result = tx.Create(&transaction)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func (transaction TransactionModel) Withdraw() (bool, error) {
	err := transaction.DB.Transaction(func(tx *gorm.DB) error {
		var sender AccountModel
		result := tx.Model(&AccountModel{}).Where(&AccountModel{
			AccountNumber: transaction.Sender,
		}).First(&sender)
		if result.Error != nil {
			return result.Error
		}

		// Check balance first
		if sender.Saldo < transaction.Amount {
			return errors.Errorf("Insufficient saldo")
		}

		// Jika saldo pengirim mencukupi
		result = result.Update("saldo", sender.Saldo-transaction.Amount)
		if result.Error != nil {
			return result.Error
		}

		transaction.TransactionType = constant.WITHDRAW
		transaction.Timestamp = time.Now().Unix()
		result = tx.Create(&transaction)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func (transaction TransactionModel) Deposit() (bool, error) {
	err := transaction.DB.Transaction(func(tx *gorm.DB) error {
		var sender AccountModel
		result := tx.Model(&AccountModel{}).Where(&AccountModel{
			AccountNumber: transaction.Sender,
		}).First(&sender).Update("saldo", sender.Saldo+transaction.Amount)
		if result.Error != nil {
			return result.Error
		}

		transaction.TransactionType = constant.DEPOSIT
		transaction.Timestamp = time.Now().Unix()
		result = tx.Create(&transaction)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
