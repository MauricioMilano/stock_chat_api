package signup_model

import user_model "github.com/MauricioMilano/stock_app/models/user"

type SignUpRequest struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type SignUpResponse struct {
	User     user_model.User `json:"user"`
	JwtToken string          `json:"token"`
}
