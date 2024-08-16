package interfaces

import (
	"time"

	"userapi/internal/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUsers() ([]models.User, error)
	GetUsersBy(startDate, endDate time.Time, minAge, maxAge int) ([]models.User, int64, error)
}
