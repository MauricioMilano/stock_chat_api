package services

import (
	"os"
	"strconv"
	"time"

	login_model "github.com/MauricioMilano/stock_app/models/login"
	signup_model "github.com/MauricioMilano/stock_app/models/signup"
	token_model "github.com/MauricioMilano/stock_app/models/token"
	user_model "github.com/MauricioMilano/stock_app/models/user"
	error_utils "github.com/MauricioMilano/stock_app/utils/error"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Auth interface {
	Login(userName, password string) (login_model.LoginResponse, error)
	SignUp(email, userName, password string) (signup_model.SignUpResponse, error)
}

type auth struct{}

func NewAuthService() *auth {
	return &auth{}
}

func (a *auth) Login(email, password string) (login_model.LoginResponse, error) {
	var user user_model.User
	err := user.GetUserByEmail(email)
	if user.ID == 0 || err.Error != nil {
		return login_model.LoginResponse{}, error_utils.ErrInvalidCredentials
	}
	expiresAt := getExpirestAt()
	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return login_model.LoginResponse{}, error_utils.ErrInvalidCredentials
	}

	tk := token_model.NewToken(user, expiresAt)

	tokenString, err2 := tk.ToTokenString()
	error_utils.ErrorCheck(err2)

	return login_model.LoginResponse{User: user, JwtToken: tokenString}, nil
}

func (a *auth) SignUp(email, userName, password string) (signup_model.SignUpResponse, error) {
	var userCheck user_model.User
	userCheck.GetUserByEmail(email)
	if userCheck.ID > 0 {
		return signup_model.SignUpResponse{}, error_utils.ErrDuplicateEmail
	}
	expiresAt, hashPass := encryptPassword(password)

	hPS := string(hashPass)
	user := user_model.User{
		Password: hPS,
		Email:    email,
		UserName: userName,
	}
	err3 := user.SaveNew()
	error_utils.DBErrorCheck(err3)
	user.Password = ""

	tk := &token_model.Token{
		UserID:   user.ID,
		UserName: user.UserName,
		Email:    user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	tokenString, err := tk.ToTokenString()
	error_utils.ErrorCheck(err)
	return signup_model.SignUpResponse{User: user, JwtToken: tokenString}, nil

}

func encryptPassword(password string) (int64, []byte) {
	expiresAt := getExpirestAt()

	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	error_utils.ErrorCheck(err)
	return expiresAt, hashPass
}

func getExpirestAt() int64 {
	jwtTTL, err := strconv.Atoi(os.Getenv("JWT_TTL"))
	error_utils.ErrorCheck(err)
	expiresAt := time.Now().Add(time.Hour * time.Duration(jwtTTL)).Unix()
	return expiresAt
}
