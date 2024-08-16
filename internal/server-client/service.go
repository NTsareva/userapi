package server_client

import (
	"time"

	"userapi/internal/interfaces"
	"userapi/internal/models"
)

type Service struct {
	repo interfaces.UserRepository
}

func NewService(r interfaces.UserRepository) *Service {
	return &Service{repo: r}
}

func (s *Service) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

func (s *Service) GetUsers() ([]models.User, error) {
	return s.repo.GetUsers()
}

func (s *Service) GenerateReport(startDate, endDate time.Time, minAge, maxAge int) ([]models.User, int64, error) {
	return s.repo.GetUsersBy(startDate, endDate, minAge, maxAge)
}
