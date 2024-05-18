package entity

import (
	"encoding/json"
	"time"
)

type MedicalRecord struct {
	ID             string    `json:"id,omitempty" db:"id"`
	PatientID      string    `json:"patientId,omitempty" db:"patient_id"`
	IdentityNumber int64     `json:"identityNumber,omitempty" db:"identity_number"`
	Symptoms       string    `json:"symptoms,omitempty" db:"symptoms"`
	Medications    string    `json:"medications,omitempty" db:"medications"`
	CreatedByID    string    `json:"createdById,omitempty" db:"created_by_id"`
	CreatedByNIP   int64     `json:"createdByNIP,omitempty" db:"created_by_nip"`
	CreatedAt      time.Time `json:"createdAt,omitempty" db:"created_at"`
	UpdatedAt      time.Time `json:"updatedAt,omitempty" db:"updated_at"`
}

// makes "omitempty" in CreatedAt and UpdatedAt works
func (u MedicalRecord) MarshalJSON() ([]byte, error) {
	type Alias MedicalRecord
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

type MedicalRecordResponse struct {
	ID                      string    `json:"id,omitempty" db:"medical_record_id"`
	PatientIdentityNumber   int64     `json:"identityDetail.identityNumber" db:"identity_number"`
	PatientPhoneNumber      string    `json:"identityDetail.phoneNumber" db:"phone_number"`
	PatientName             string    `json:"identityDetail.name" db:"name"`
	PatientBirthDate        string    `json:"identityDetail.birthDate" db:"birth_date"`
	PatientGender           string    `json:"identityDetail.gender" db:"gender"`
	PatientIdentityImageUrl string    `json:"identityDetail.identityCardScanImg" db:"identity_image_url"`
	Symptoms                string    `json:"symptoms,omitempty" db:"symptoms"`
	Medications             string    `json:"medications,omitempty" db:"medications"`
	CreatedByID             string    `json:"createdBy.userId" db:"created_by_id"`
	CreatedByNIP            int64     `json:"createdBy.nip" db:"created_by_nip"`
	CreatedByName           string    `json:"createdBy.name" db:"created_by_name"`
	CreatedAt               time.Time `json:"createdAt,omitempty" db:"created_at"`
	UpdatedAt               time.Time `json:"updatedAt,omitempty" db:"updated_at"`
}

// makes "omitempty" in CreatedAt and UpdatedAt works
func (u MedicalRecordResponse) MarshalJSON() ([]byte, error) {
	type Alias MedicalRecordResponse
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
