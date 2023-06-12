package repository

import (
	"context"
	"encoding/json"
	"errors"
	"go-redis-mysql_userAuth/pkg/domain"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(body domain.User) (domain.UserDetails, error)
	UserLogin(body domain.User) (domain.User, error)
}
type userRepo struct {
	sDB *gorm.DB
	rDB *redis.Client
}

func NewUserRepo(sdb *gorm.DB, rdb *redis.Client) UserRepo {
	return &userRepo{sdb, rdb}
}
func (r *userRepo) CreateUser(body domain.User) (domain.UserDetails, error) {
	user := domain.User{
		UserName: body.UserName,
		Email:    body.Email,
		Password: body.Password,
	}
	result := r.sDB.Create(&user)

	if result.Error != nil {
		return domain.UserDetails{}, result.Error
	}
	log.Println("from create user",user)
	return domain.UserDetails{
		UserId:   user.ID,
		UserName: user.UserName,
		Email:    user.Email,
	}, nil
}
func (r *userRepo) UserLogin(body domain.User) (domain.User, error) {
	user := domain.User{}
	ctx := context.Background()

	result, _ := r.rDB.Get(ctx, body.Email).Result()
	if len(result) > 0 {
		if err := json.Unmarshal([]byte(result), &user); err != nil {
			return user, err
		}
		if user.ID == 0 {
			log.Println("Calling from redis")
			return user, errors.New("user not found")
		}
		log.Println("Calling from redis",user)
		return user, nil
	}

	res := r.sDB.Model(domain.User{}).Select("id","user_name","email","password").Where("email", body.Email).Find(&user)
	if res.Error != nil || user.ID == 0 {
		log.Println("Calling from mysql", res.Error)
		return user, errors.New("user not found")
	}
	jsonBytes, err := json.Marshal(&user)
	if err != nil {
		return user, err
	}
	jsonString := string(jsonBytes)
	log.Println(jsonString)
	err = r.rDB.Set(ctx, body.Email, jsonString, 7*24*time.Hour).Err()
	if err != nil {
		return user, err
	}
	log.Println("Calling from mysql", res.Error)
	return user, nil
}
