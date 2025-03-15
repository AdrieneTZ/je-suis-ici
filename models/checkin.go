package models

import "time"

/*
Define location model
*/
type Location struct {
	Latitude     float64 `json:"latitude" db:"latitude"`
	Longitude    float64 `json:"longitude" db:"longitude"`
	LocationName string  `json:"locationName" db:"location_name"`
}

/*
Define check in record model
*/
type CheckIn struct {
	ID        int       `json:"-" db:"id"`             // check in record table PK
	UserID    int       `json:"-" db:"user_id"`        // user table Id, FK
	UserDID   string    `json:"userDid" db:"user_did"` // user DID from ATProtocol
	Text      string    `json:"text,omitempty"`
	Location  Location  `json:"location" db:"location"`
	Images    []string  `json:"images,omitempty" db:"images"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"indexed_at"`
}
