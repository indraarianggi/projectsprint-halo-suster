package user

type RegisterITRequest struct {
	NIP      int    `json:"nip" validate:"required,nip=it"`
	Name     string `json:"name" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=5,max=33"`
}
