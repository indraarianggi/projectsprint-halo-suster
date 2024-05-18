package input

type AddPatientRequest struct {
	IdentityNumber   int64  `json:"identityNumber" validate:"required,numlen=16"`
	PhoneNumber      string `json:"phoneNumber" validate:"required,min=10,max=15,startswith=+62"`
	Name             string `json:"name" validate:"required,min=5,max=50"`
	BirthDate        string `json:"birthDate" validate:"required,iso8601_date"`
	Gender           string `json:"gender" validate:"required,oneof=male female"`
	IdentityImageUrl string `json:"identityCardScanImg" validate:"required,image_url"`
}

type GetListPatientRequest struct {
	IdentityNumber string `param:"identityNumber" validate:"omitempty,number"`
	PhoneNumber    string `param:"phoneNumber"`
	Name           string `param:"name"`
	CreatedAt      string `param:"createdAt"`
	Limit          string `param:"limit" validate:"omitempty,number"`
	Offset         string `param:"offset" validate:"omitempty,number"`
}
