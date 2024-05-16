CREATE TABLE medical_records (
    id VARCHAR(26) PRIMARY KEY,
    patient_id VARCHAR(26) REFERENCES patients(id) ON DELETE SET NULL,
    identity_number BIGINT,
    created_by_id VARCHAR(26) REFERENCES users(id) ON DELETE SET NULL,
    created_by_nip BIGINT,
    symptoms TEXT,
    medications TEXT,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL
);

CREATE INDEX idx_medical_records_identity_number ON medical_records(identity_number);
CREATE INDEX idx_medical_records_created_by_id ON medical_records(created_by_id);
CREATE INDEX idx_medical_records_created_by_nip ON medical_records(created_by_nip);