package login_model

import user_model "github.com/MauricioMilano/stock_app/models/user"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User     user_model.User `json:"user"`
	JwtToken string          `json:"token"`
}
