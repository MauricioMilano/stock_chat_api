package controllers

import (
	"encoding/json"
	"net/http"

	login_model "github.com/MauricioMilano/stock_app/models/login"
	signup_model "github.com/MauricioMilano/stock_app/models/signup"
	"github.com/MauricioMilano/stock_app/services"
	"github.com/MauricioMilano/stock_app/utils"
	error_utils "github.com/MauricioMilano/stock_app/utils/error"
)

type AuthController struct {
	authServ services.Auth
}

func (a *AuthController) RegisterService(s services.Auth) {
	a.authServ = s
}

func (a *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodPost {
		lP := login_model.LoginRequest{}
		err := utils.ParseBody(r, &lP)
		if err != nil {
			utils.ErrResponse(error_utils.ErrInRequestMarshaling, w)
			return
		}

		res, err := a.authServ.Login(lP.Email, lP.Password)
		if err != nil {
			utils.ErrResponse(err, w)
			return
		}

		res.User.Password = ""
		data, err := json.Marshal(res)
		error_utils.ErrorCheck(err)

		utils.Ok(data, w)
	}

}

func (a *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodPost {

		lP := signup_model.SignUpRequest{}
		err := utils.ParseBody(r, &lP)
		if err != nil {
			utils.ErrResponse(error_utils.ErrInRequestMarshaling, w)
			return
		}

		res, err := a.authServ.SignUp(lP.Email, lP.UserName, lP.Password)
		if err != nil {
			utils.ErrResponse(err, w)
			return
		}

		res.User.Password = ""
		data, err := json.Marshal(res)
		error_utils.ErrorCheck(err)

		utils.Ok(data, w)
	}
}
