package models

import "time"

// User defines user model
type User struct {
	ID          int       `json:"-" db:"id"`                     // user table PK
	DID         string    `json:"did" db:"did"`                  // DID from ATProtocol
	Handle      string    `json:"handle" db:"handle"`            // user account
	DisplayName string    `json:"displayName" db:"display_name"` // display name
	Avatar      string    `json:"avatar,omitempty" db:"avatar"`
	Password    string    `json:"-" db:"password"` // hashed password
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

// TableName returns table name
func (u *User) TableName() string {
	return "users"
}
