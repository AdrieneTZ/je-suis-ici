package models

import "time"

/*
Define check in model
*/
type CheckIn struct {
	ID           int       `json:"-" database:"id"`             // check in table PK
	UserID       int       `json:"-" database:"user_id"`        // user table Id, FK
	UserDID      string    `json:"userDid" database:"user_did"` // user DID from ATProtocol
	Text         string    `json:"text,omitempty"`
	Latitude     float64   `json:"latitude" database:"latitude"`
	Longitude    float64   `json:"longitude" database:"longitude"`
	LocationName string    `json:"locationName" database:"location_name"`
	Images       []string  `json:"images,omitempty" database:"images"`
	CreatedAt    time.Time `json:"createdAt" database:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt" database:"indexed_at"`
}
