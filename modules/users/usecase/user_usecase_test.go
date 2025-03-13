package usecase

import (
	"bank/modules/users/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserUsecase_GetUsers(t *testing.T) {
	// สร้าง Mock UserRepository
	userRepoMock := mocks.NewUserRepositoryMock()

	// สร้าง UserUsecase โดยใช้ Mock
	userUsecase := NewUserUsecase(userRepoMock)

	// เรียกใช้ GetUsers()
	users, err := userUsecase.GetUsers()

	// ตรวจสอบผลลัพธ์
	assert.NoError(t, err)
	assert.Equal(t, 3, len(users)) // ตรวจสอบจำนวนข้อมูลที่ return

	// ตรวจสอบข้อมูล mock รายการแรก
	assert.Equal(t, 1, users[0].UserId)
	assert.Equal(t, "john_doe", users[0].Username)
	assert.Equal(t, "password123", users[0].Password)
	assert.Equal(t, "john@example.com", users[0].Email)
	assert.Equal(t, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), users[0].UserSince)
	assert.Equal(t, time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC), users[0].LastVisit)

	// ตรวจสอบข้อมูล mock รายการที่สอง
	assert.Equal(t, 2, users[1].UserId)
	assert.Equal(t, "jane_doe", users[1].Username)
	assert.Equal(t, "password456", users[1].Password)
	assert.Equal(t, "jane@example.com", users[1].Email)
	assert.Equal(t, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), users[1].UserSince)
	assert.Equal(t, time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC), users[1].LastVisit)

	// ตรวจสอบข้อมูล mock รายการที่สาม
	assert.Equal(t, 3, users[2].UserId)
	assert.Equal(t, "alice_smith", users[2].Username)
	assert.Equal(t, "password789", users[2].Password)
	assert.Equal(t, "alice@example.com", users[2].Email)
	assert.Equal(t, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), users[2].UserSince)
	assert.WithinDuration(t, time.Now(), users[2].LastVisit, time.Second) // ตรวจสอบ last visit ว่าอยู่ในช่วงเวลาปัจจุบัน
}
