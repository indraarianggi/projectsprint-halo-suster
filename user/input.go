package user

type RegisterITRequest struct {
	NIP      int64  `json:"nip" validate:"required,nip=it"`
	Name     string `json:"name" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=5,max=33"`
}

type LoginRequest struct {
	NIP      int64  `json:"nip" validate:"required,nip"`
	Password string `json:"password" validate:"required,min=5,max=33"`
}

type RegisterNurseRequest struct {
	NIP              int64  `json:"nip" validate:"required,nip=nurse"`
	Name             string `json:"name" validate:"required,min=5,max=50"`
	IdentityImageUrl string `json:"identityCardScanImg" validate:"required,image_url"`
}

type NurseAccessRequest struct {
	ID       string `param:"id" validate:"required"`
	Password string `json:"password" validate:"required,min=5,max=33"`
}
