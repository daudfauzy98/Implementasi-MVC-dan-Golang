package controller

import (
	"net/http"

	"github.com/daudfauzy98/Implementasi-MVC-dan-Golang/app/model"
	"github.com/daudfauzy98/Implementasi-MVC-dan-Golang/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//AccountController tipe struct
type AccountController struct {
	DB *gorm.DB
}

func (ctrl AccountController) CreateAccount(ctx *gin.Context) {
	account := model.AccountModel{
		DB: ctrl.DB,
	}

	err := ctx.Bind(&account)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	hashPassword, err := utils.HashGenerator(account.Password)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	account.Password = hashPassword

	flag, err := account.InsertNewAccount()
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WrapAPIError(ctx, "Unknows failed to insert account", http.StatusInternalServerError)
		return
	}

	utils.WrapAPISuccess(ctx, "Success!", http.StatusOK)
}

func (ctrl AccountController) GetAccount(ctx *gin.Context) {
	idAccount := ctx.MustGet("account_number").(int)

	accountModel := model.AccountModel{
		DB: ctrl.DB,
	}

	flag, err, transaction, account := accountModel.GetAccountDetail(idAccount)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WrapAPIError(ctx, "Unknown error to get account", http.StatusInternalServerError)
		return
	}

	utils.WrapAPIData(ctx, map[string]interface{}{
		"account":     account,
		"transaction": transaction,
	}, http.StatusOK, "Success!")

	return
}

func (ctrl AccountController) Login(ctx *gin.Context) {
	authModel := model.AuthModel{
		DB: ctrl.DB,
	}

	err := ctx.Bind(&authModel)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err, token := authModel.Login()
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WrapAPIError(ctx, "Unknows failed to login", http.StatusInternalServerError)
		return
	}

	utils.WrapAPIData(ctx, map[string]interface{}{
		"token": token,
	}, http.StatusOK, "Success!")
}
