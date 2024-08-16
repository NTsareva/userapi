package server_client

import (
	"time"

	"userapi/internal/models"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(user *models.User) error {
	user.RecordingDate = time.Now().Unix()
	return r.db.Create(user).Error
}

func (r *Repository) GetUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *Repository) GetUsersByCriteria(startDate, endDate time.Time, minAge, maxAge int) ([]models.User, int64, error) {
	var users []models.User
	query := r.db.Model(&models.User{})

	if !startDate.IsZero() {
		startTimestamp := startDate.Unix()
		query = query.Where("extract(epoch from recording_time) >= ?", startTimestamp)
	}

	if !endDate.IsZero() {
		endTimestamp := endDate.Unix()
		query = query.Where("extract(epoch from recording_time) <= ?", endTimestamp)
	}

	if minAge > 0 {
		query = query.Where("age >= ?", minAge)
	}

	if maxAge > 0 {
		query = query.Where("age <= ?", maxAge)
	}

	var count int64
	err := query.Find(&users).Count(&count).Error
	return users, count, err
}
