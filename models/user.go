package models

import "time"

/*
Define use model
*/
type User struct {
	ID          string    `json:"id"`          // DID
	Handle      string    `json:"handle"`      // user name
	DisplayName string    `json:"displayName"` // display name
	Avatar      string    `json:"avatar,omitempty"`
	Password    string    `json:"-"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
