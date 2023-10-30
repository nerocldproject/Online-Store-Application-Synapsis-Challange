package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"osa.synapsis.chalange/modules/user/model"
	"osa.synapsis.chalange/modules/user/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return UserService{userRepo}
}

func (u UserService) Register(user model.User) (err error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return
	}

	user.Password = string(hashPass)
	err = u.userRepo.InsertUser(user)
	return
}

func (u UserService) Login(user model.User) (token string, err error) {
	data, err := u.userRepo.GetUserByUsername(user.Password)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(user.Password))
	if err != nil {
		return
	}

	claims := model.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
		},
		Username: user.UserName,
	}

	tokeNB := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokeNB.SignedString([]byte(viper.GetString("SECRET_KEY")))

	return
}