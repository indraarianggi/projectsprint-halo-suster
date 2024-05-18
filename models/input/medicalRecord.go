package input

type AddMedicalRecordRequest struct {
	IdentityNumber int64  `json:"identityNumber" validate:"required,numlen=16"`
	Symptoms       string `json:"symptoms" validate:"required,min=1,max=2000"`
	Medications    string `json:"medications" validate:"required,min=1,max=2000"`
	CreatedByID    string // value from user claims
	CreatedByNIP   int64  // value from user claims
}

type GetListMedicalRecordRequest struct {
	IdentityNumber string `param:"identityDetail.identityNumber" validate:"omitempty,number"`
	CreatedByID    string `param:"createdBy.userId"`
	CreatedByNIP   string `param:"createdBy.nip"`
	CreatedAt      string `param:"createdAt"`
	Limit          string `param:"limit" validate:"omitempty,number"`
	Offset         string `param:"offset" validate:"omitempty,number"`
}
