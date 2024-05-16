package models

type UserWithToken struct {
	ID    string `json:"userId"`
	NIP   int    `json:"nip"`
	Name  string `json:"name"`
	Token string `json:"accessToken"`
}
