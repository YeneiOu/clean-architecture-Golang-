package service

import (
	"bank/modules/entities"
	"fmt"
	"log"
)

type userService struct {
	UserRepo entities.UserRepository
}

func NewUserService(userRepo entities.UserRepository) *userService {
	return &userService{
		UserRepo: userRepo,
	}
}

func (s userService) GetUsers() ([]entities.UserResponseService, error) {
	user, err := s.UserRepo.GetUsers()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	uesrResponse := []entities.UserResponseService{}
	for _, user := range user {
		userResponse := entities.UserResponseService{
			UserId:   user.UserId,
			Username: user.Username,
			Email:    user.Email,
		}
		uesrResponse = append(uesrResponse, userResponse)
	}
	fmt.Println("User uesrResponse =", uesrResponse)
	return uesrResponse, nil
}
