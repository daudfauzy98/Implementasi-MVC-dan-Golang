package model

import (
	"errors"

	"github.com/daudfauzy98/Implementasi-MVC-dan-Golang/app/utils"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type AuthModel struct {
	DB       *gorm.DB
	Name     string `jason: "name"`
	Password string `jason: "password"`
}

func (auth AuthModel) Login() (bool, error, string) {
	var account AccountModel

	result := auth.DB.Where(&AccountModel{Name: auth.Name}).First(&account)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, errors.New("Account not found!"), ""
		}
		return false, result.Error, ""
	}

	err := utils.HashComparator([]byte(account.Password), []byte(auth.Password))
	if err != nil {
		return false, errors.New("Incorrect password"), ""
	}

	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":            auth.Name,
		"account_account": account.AccountNumber,
	})

	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		return false, err, ""
	}

	return true, nil, token
}
