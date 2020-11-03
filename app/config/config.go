package config

import (
	"github.com/daudfauzy98/Implementasi-MVC-dan-Golang/app/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//DBinit untuk menginisialisasi koneksi ke database agar bisa digunakan
func DBinit() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:admin123@/digitalent_bank?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database!" + err.Error())
	}

	db.AutoMigrate(new(model.AccountModel), new(model.TransactionModel))

	return db
}
