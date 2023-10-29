package repository

import (
	"fmt"

	"gorm.io/gorm"
	"osa.synapsis.chalange/modules/user/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository (db *gorm.DB) UserRepository {
	return UserRepository{db}
}

func (u UserRepository) InsertUser(user model.User) (err error) {
	tx := u.db.Begin()
	func(){
		if r := recover(); r != nil {
			tx.Rollback()
			err = r.(error)
		}
	}()

	var count int64
	result := tx.Table("customer").Where("user_name = ? and deleted_at is null", user.UserName).Count(&count)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	if count > 0 {
		tx.Rollback()
		return fmt.Errorf("maaf, username sudah digunakan")
	}

	result = tx.Table("customer").Omit("user_id").Create(&user)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	
	return tx.Commit().Error
}

func (u UserRepository) GetUserByUsername(username string) (user model.User, err error) {
	result := u.db.Table("customer").Where("user_name = ? and deleted_at is null", username).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	if result.RowsAffected == 0 {
		return user, fmt.Errorf("maaf, data tidak ditemukan")
	}

	return
}