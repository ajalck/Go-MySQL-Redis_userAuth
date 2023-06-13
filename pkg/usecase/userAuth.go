package usecase

import (
	"crypto/md5"
	"errors"
	"fmt"
	"go-redis-mysql_userAuth/pkg/domain"
	"go-redis-mysql_userAuth/pkg/repository"
)

type UserUseCase interface {
	CreateUser(body domain.User) (domain.UserDetails, error)
	UserLogin(body domain.User) (domain.UserDetails, error)
}
type userUseCase struct {
	repo repository.UserRepo
}

func NewUserUseCase(repo repository.UserRepo) UserUseCase {
	return &userUseCase{repo}
}

func (u *userUseCase) CreateUser(body domain.User) (domain.UserDetails, error) {

	bytes := fmt.Sprintf("%x", md5.Sum([]byte(body.Password)))
	body.Password = string(bytes)

	user, err := u.repo.CreateUser(body)
	if err != nil {
		return user, err
	}
	return user, nil

}
func (u *userUseCase) UserLogin(body domain.User) (domain.UserDetails, error) {
	user, err := u.repo.UserLogin(body)
	if err != nil {
		return domain.UserDetails{}, err
	}
	bytes := fmt.Sprintf("%x", md5.Sum([]byte(body.Password)))
	reqPassword := string(bytes)
	if user.Password != reqPassword {
		return domain.UserDetails{}, errors.New("invalid username of password")
	}
	return domain.UserDetails{
		UserId:   user.ID,
		UserName: user.UserName,
		Email:    user.Email,
	}, nil

}
