package controller

import (
	"net/http"

	"github.com/daudfauzy98/Implementasi-MVC-dan-Golang/app/model"
	"github.com/daudfauzy98/Implementasi-MVC-dan-Golang/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TransactionController tipe struct
type TransactionController struct {
	DB *gorm.DB
}

func (ctrl TransactionController) Transfer(ctx *gin.Context) {
	transactionModel := model.TransactionModel{
		DB: ctrl.DB,
	}

	err := ctx.Bind(&transactionModel)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err := transactionModel.Transfer()
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WrapAPIError(ctx, "Unknows failed to transfer", http.StatusInternalServerError)
		return
	}

	utils.WrapAPISuccess(ctx, "Success!", http.StatusOK)
	return
}

func (ctrl TransactionController) Withdraw(ctx *gin.Context) {
	transactionModel := model.TransactionModel{
		DB: ctrl.DB,
	}

	err := ctx.Bind(&transactionModel)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err := transactionModel.Withdraw()
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WrapAPIError(ctx, "Unknows failed to transfer", http.StatusInternalServerError)
		return
	}

	utils.WrapAPISuccess(ctx, "Success!", http.StatusOK)
	return
}

func (ctrl TransactionController) Deposit(ctx *gin.Context) {
	transactionModel := model.TransactionModel{
		DB: ctrl.DB,
	}

	err := ctx.Bind(&transactionModel)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err := transactionModel.Deposit()
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WrapAPIError(ctx, "Unknows failed to transfer", http.StatusInternalServerError)
		return
	}

	utils.WrapAPISuccess(ctx, "Success!", http.StatusOK)
	return
}
