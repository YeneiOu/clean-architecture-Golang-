package usecase

import (
	"bank/modules/entities"
	"fmt"
)

type UserUse struct {
	UserRepo entities.UserRepository
}

func NewUserUsecase(userRepo entities.UserRepository) entities.UserUsecase {
	return &UserUse{
		UserRepo: userRepo,
	}
}

func (u *UserUse) GetUsers() ([]entities.UserResponse, error) {
	res, err := u.UserRepo.GetUsers()

	if err != nil {
		fmt.Println("res UserUse", err)
		return nil, err
	}

	return res, err
}

func (u *UserUse) GetUserByID(id int) (*entities.UserResponse, error) {
	res, err := u.UserRepo.GetUserByID(id)
	return res, err
}
