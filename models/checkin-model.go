package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// Define check in model
type CheckIn struct {
	ID           int       `json:"-" db:"id"`             // check in table PK
	UserID       int       `json:"userId" db:"user_id"`   // user table Id, FK
	UserDID      string    `json:"userDid" db:"user_did"` // user DID from ATProtocol
	Text         string    `json:"text,omitempty" db:"text"`
	Latitude     float64   `json:"latitude" db:"latitude"`
	Longitude    float64   `json:"longitude" db:"longitude"`
	LocationName string    `json:"locationName" db:"location_name"`
	Images       Images    `json:"images,omitempty" db:"images"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt" db:"indexed_at"`
}

// TableName returns table name
func (ch *CheckIn) TableName() string {
	return "check_ins"
}

// Images is a custom type for handling JSON arrays in PostgreSQL
type Images []string

// Value implements the driver.Valuer interface for Images
// converts []string to JSONB when saving to the database
func (i Images) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}

	return json.Marshal(i)
}

// Scan implements the sql.Scanner interface for Images
// converts JSONB from the database to []string
func (i *Images) Scan(value interface{}) error {
	if value == nil {
		*i = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan Images: value is not []byte")
	}

	return json.Unmarshal(bytes, i)
}
