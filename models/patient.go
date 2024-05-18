package models

import (
	"encoding/json"
	"time"
)

type Patient struct {
	ID               string    `json:"id" db:"id"`
	IdentityNumber   int64     `json:"identityNumber,omitempty" db:"identity_number"`
	Name             string    `json:"name,omitempty" db:"name"`
	PhoneNumber      string    `json:"phoneNumber,omitempty" db:"phone_number"`
	BirthDate        time.Time `json:"birthDate,omitempty" db:"birth_date"`
	Gender           string    `json:"gender,omitempty" db:"gender"`
	IdentityImageUrl string    `json:"identityCardScanImg,omitempty" db:"identity_image_url"`
	CreatedAt        time.Time `json:"createdAt,omitempty" db:"created_at"`
	UpdatedAt        time.Time `json:"updatedAt,omitempty" db:"updated_at"`
}

// makes "omitempty" in CreatedAt and UpdatedAt works
func (u Patient) MarshalJSON() ([]byte, error) {
	type Alias Patient
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
