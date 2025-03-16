package services

import (
	"errors"
	"je-suis-ici/models"
	"je-suis-ici/repositories"
)

type CheckInService struct {
	checkInRepo *repositories.CheckInRepository
}

func NewCheckInService(
	checkInRepo *repositories.CheckInRepository,
) *CheckInService {
	return &CheckInService{
		checkInRepo: checkInRepo,
	}
}

// CreateCheckIn creates a new check-in
func (s *CheckInService) CreateCheckIn(checkIn *models.CheckIn) (int, error) {
	// Validate check-in data
	if checkIn.UserID == 0 {
		return 0, errors.New("user ID is required")
	}
	if checkIn.UserDID == "" {
		return 0, errors.New("user DID is required")
	}

	// Create the check-in
	return s.checkInRepo.CreateCheckIn(checkIn)
}

// GetCheckIn gets a check-in by ID
func (s *CheckInService) GetCheckIn(id int) (*models.CheckIn, error) {
	return s.checkInRepo.GetCheckIn(id)
}
