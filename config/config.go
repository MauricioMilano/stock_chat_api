package config

import (
	"fmt"
	"os"

	error_utils "github.com/MauricioMilano/stock_app/utils/error"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type ConfigOpts struct {
}

func (config *ConfigOpts) ConnectDB() {
	dbUserName := os.Getenv("APP_POSTGRES_USER")
	dbUserPassword := os.Getenv("APP_POSTGRES_PASS")
	dbHost := os.Getenv("APP_POSTGRES_HOST")
	dbPort := os.Getenv("APP_POSTGRES_PORT")
	dbName := os.Getenv("APP_POSTGRES_DB")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUserName, dbUserPassword, dbName)

	d, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	error_utils.ErrorCheck(err)

	db = d
}

func GetDB() *gorm.DB {
	return db
}
