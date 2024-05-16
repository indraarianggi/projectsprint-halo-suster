package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

type User struct {
	ID               string         `json:"id,omitempty" db:"id"`
	NIP              int64          `json:"nip,omitempty" db:"nip"`
	Name             string         `json:"name,omitempty" db:"name"`
	Role             string         `json:"role,omitempty" db:"role"`
	Password         sql.NullString `json:"-" db:"password"`
	IdentityImageUrl string         `json:"identityCardScanImg,omitempty" db:"identity_image_url"`
	CreatedAt        time.Time      `json:"createdAt,omitempty" db:"created_at"`
	UpdatedAt        time.Time      `json:"updatedAt,omitempty" db:"updated_at"`
	DeletedAt        sql.NullTime   `json:"-" db:"deleted_at"`
}

type UserWithToken struct {
	ID    string `json:"userId"`
	NIP   int64  `json:"nip"`
	Name  string `json:"name"`
	Token string `json:"accessToken"`
}

// makes "omitempty" in CreatedAt and UpdatedAt works (thanks chat gpt 😆)
func (u User) MarshalJSON() ([]byte, error) {
	type Alias User
	aux := &struct {
		*Alias
		CreatedAt interface{} `json:"createdAt,omitempty"`
		UpdatedAt interface{} `json:"updatedAt,omitempty"`
	}{
		Alias: (*Alias)(&u),
	}

	if u.CreatedAt.IsZero() {
		aux.CreatedAt = nil
	} else {
		aux.CreatedAt = u.CreatedAt
	}

	return json.Marshal(aux)
}
