package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Username     string
	Email        string
	Type         UserType
	PasswordHash string
	Verified     bool
	Banned       bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (user User) IsAdmin() bool {
	return user.Type == USER_TYPE_ADMIN
}

func (user User) IsTeacher() bool {
	return user.Type == USER_TYPE_TEACHER
}

func (user User) IsStudent() bool {
	return user.Type == USER_TYPE_STUDENT
}