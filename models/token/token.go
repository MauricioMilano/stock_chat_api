package token_model

import (
	"os"

	user_model "github.com/MauricioMilano/stock_app/models/user"

	jwt "github.com/golang-jwt/jwt"
)

type Token struct {
	UserID   uint
	UserName string
	Email    string
	*jwt.StandardClaims
}

func (tk *Token) ToTokenString() (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	jwtSecret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func NewToken(user user_model.User, expiresAt int64) *Token {
	tk := &Token{
		UserID:   user.ID,
		UserName: user.UserName,
		Email:    user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	return tk
}
