package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/MauricioMilano/stock_app/config"
	chat_model "github.com/MauricioMilano/stock_app/models/chat"
	chatroom_model "github.com/MauricioMilano/stock_app/models/chatroom"
	error_model "github.com/MauricioMilano/stock_app/models/error"
	user_model "github.com/MauricioMilano/stock_app/models/user"
	error_utils "github.com/MauricioMilano/stock_app/utils/error"

	"github.com/joho/godotenv"
)

type emptyOk struct {
	Message string
}

type JWTProps string

func ParseByteArray(r []byte, x interface{}) error {
	if err := json.Unmarshal(r, x); err != nil {
		return err
	}
	return nil
}

func ParseBody(r *http.Request, x interface{}) error {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return err
		}
	}
	return nil
}
func ErrResponse(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	errCode := codeFrom(err)
	w.WriteHeader(errCode)
	res := error_model.ErrorResponse{Message: err.Error(), Status: false, Code: errCode}
	data, err := json.Marshal(res)
	error_utils.ErrorCheck(err)
	w.Write(data)
}
func codeFrom(err error) int {
	switch err {
	case error_utils.ErrInvalidCredentials:
		return http.StatusBadRequest
	case error_utils.ErrDuplicateEmail:
		return http.StatusBadRequest
	case error_utils.ErrInRequestMarshaling:
		return http.StatusBadRequest
	case error_utils.ErrInRequestMarshaling:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func Ok(res []byte, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func OkEmpty(message string, w http.ResponseWriter) {
	m := emptyOk{message}
	res, err := json.Marshal(m)
	error_utils.ErrorCheck(err)
	Ok(res, w)
}

func TestHelper() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database := config.ConfigOpts{}
	database.ConnectDB()
	db := config.GetDB()
	err = db.AutoMigrate(&user_model.User{}, &chatroom_model.ChatRoom{}, &chat_model.Chat{})
	error_utils.ErrorCheck(err)
}
