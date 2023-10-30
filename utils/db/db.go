package db

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnectDB() (*gorm.DB, error) {
	username := viper.GetString("USERNAME_DB")
	password := viper.GetString("PASSWORD_DB")
	database := viper.GetString("DATABASE_DB")
	host := viper.GetString("HOST_DB")
	port := viper.GetInt("PORT_DB")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
	host,
	username,
	password,
	database,
	port,)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Success connect to DB")

	return db, nil
}