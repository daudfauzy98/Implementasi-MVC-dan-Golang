package model

import (
	"time"

	"github.com/daudfauzy98/Implementasi-MVC-dan-Golang/app/constant"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Transaction struct {
	ID                     int    `gorm: "primary_key" json:"-"`
	TransactionType        int    `json:"transaction_type, omitempty"`
	TransactionDescription string `json:"transaction_description"`
	Sender                 int    `json:"sender"`
	Amount                 int    `json:"amount"`
	Recipient              int    `json:"recipient"`
	Timestamp              int64  `json:"timestamp, omitempty"`
}

type TransactionModel struct {
	DB *gorm.DB
}

func (model TransactionModel) Transfer(transaction Transaction) (bool, error) {
	err := model.DB.Transaction(func(tx *gorm.DB) error {
		var sender, recipient Account

		result := tx.Model(&Account{}).Where(&Account{AccountNumber: transaction.Sender}).First(&sender)
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

		result = tx.Model(&Account{}).Where(Account{
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

func (model TransactionModel) Withdraw(transaction Transaction) (bool, error) {
	err := model.DB.Transaction(func(tx *gorm.DB) error {
		var sender Account
		result := tx.Model(&Account{}).Where(&Account{
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

func (model TransactionModel) Deposit(transaction Transaction) (bool, error) {
	err := model.DB.Transaction(func(tx *gorm.DB) error {
		var sender Account
		result := tx.Model(&Account{}).Where(&Account{
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
