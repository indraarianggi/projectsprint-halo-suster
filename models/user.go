package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID               string         `json:"id" db:"id"`
	NIP              int            `json:"nip" db:"nip"`
	Name             string         `json:"name" db:"name"`
	Role             string         `json:"role" db:"role"`
	Password         sql.NullString `json:"-" db:"password"`
	IdentityImageUrl string         `json:"identityCardScanImg" db:"identity_image_url"`
	CreatedAt        time.Time      `json:"createdAt" db:"created_at"`
	UpdatedAt        time.Time      `json:"updatedAt" db:"updated_at"`
}
