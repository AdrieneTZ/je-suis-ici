package repositories

import (
	"je-suis-ici/database"
	"je-suis-ici/models"
	"time"
)

type CheckInRepository struct {
	db *database.DB
}

func NewCheckInRepository(db *database.DB) *CheckInRepository {
	return &CheckInRepository{db: db}
}

// CreateCheckIn creates a new check-in record
func (r *CheckInRepository) CreateCheckIn(checkIn *models.CheckIn) (int, error) {
	query := `
		INSERT INTO checkins 
		(user_id, user_did, text, latitude, longitude, location_name, images, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`

	var id int
	now := time.Now().UTC()
	checkIn.CreatedAt = now
	checkIn.UpdatedAt = now

	err := r.db.QueryRowx(
		query,
		checkIn.UserID,
		checkIn.UserDID,
		checkIn.Text,
		checkIn.Latitude,
		checkIn.Longitude,
		checkIn.LocationName,
		checkIn.Images, // use Images.Value() from checkin model
		checkIn.CreatedAt,
		checkIn.UpdatedAt,
	).Scan(&id)

	return id, err
}

// GetCheckIn fetches a check-in by ID
func (r *CheckInRepository) GetCheckIn(id int) (*models.CheckIn, error) {
	query := `
		SELECT * FROM checkins WHERE id = $1
	`

	var checkIn models.CheckIn
	err := r.db.Get(&checkIn, query, id)
	if err != nil {
		return nil, err
	}

	return &checkIn, nil
}
