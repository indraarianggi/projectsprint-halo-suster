package entity

import "time"

type UserClaims struct {
	ID        string    `json:"id"`
	NIP       int64     `json:"nip"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	ExpiredAt time.Time `json:"expiredAt"`
}
