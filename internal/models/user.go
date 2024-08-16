package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID            string    `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Age           int       `json:"age"`
	RecordingDate int64     `json:"recording_date" gorm:"-"`
	RecordingTime time.Time `gorm:"type:timestamp;not null;default:current_timestamp" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}

	u.RecordingTime = time.Now().UTC()
	u.RecordingDate = u.RecordingTime.Unix()
	return
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	u.RecordingDate = u.RecordingTime.Unix()
	return
}
