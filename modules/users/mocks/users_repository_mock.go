package mocks

import (
	"bank/modules/entities"
	"errors"
	"time"
)

type UserRepositoryMock struct {
	user map[string]entities.UserResponse
}

func NewUserRepositoryMock() *UserRepositoryMock {

	return &UserRepositoryMock{user: make(map[string]entities.UserResponse)}
}

func (m *UserRepositoryMock) GetUsers() ([]entities.UserResponse, error) {
	// กำหนดเวลาสำหรับ mock data
	now := time.Now()
	userSince := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	lastVisit := time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC)

	// สร้างข้อมูล mock 3 รายการ
	users := []entities.UserResponse{
		{
			UserId:    1,
			Username:  "john_doe",
			Password:  "password123",
			Email:     "john@example.com",
			UserSince: userSince,
			LastVisit: lastVisit,
		},
		{
			UserId:    2,
			Username:  "jane_doe",
			Password:  "password456",
			Email:     "jane@example.com",
			UserSince: userSince,
			LastVisit: lastVisit,
		},
		{
			UserId:    3,
			Username:  "alice_smith",
			Password:  "password789",
			Email:     "alice@example.com",
			UserSince: userSince,
			LastVisit: now, // ใช้เวลาปัจจุบันสำหรับ last visit
		},
	}

	return users, nil
}

func (m UserRepositoryMock) GetUserByID(id int) (*entities.UserResponse, error) {

	for _, user := range m.user {
		if user.UserId == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}
