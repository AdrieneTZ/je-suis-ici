package models

import "time"

/*
Define use model
*/
type User struct {
	ID          int       `json:"-" database:"id"`                     // user table PK
	DID         string    `json:"did" database:"did"`                  // DID from ATProtocol
	Handle      string    `json:"handle" database:"handle"`            // user name
	DisplayName string    `json:"displayName" database:"display_name"` // display name
	Avatar      string    `json:"avatar,omitempty" database:"avatar"`
	Password    string    `json:"-" database:"password"` // hashed password
	CreatedAt   time.Time `json:"createdAt" database:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" database:"updated_at"`
}
