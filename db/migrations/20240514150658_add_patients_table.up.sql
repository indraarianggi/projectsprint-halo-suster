CREATE TABLE patients (
    id VARCHAR(26) PRIMARY KEY,
    identity_number BIGINT,
    name VARCHAR(30),
    phone_number VARCHAR(16),
    birth_date TIMESTAMP,
    gender VARCHAR(6) CHECK (gender IN ('male', 'female')),
    identity_image_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL
);

CREATE INDEX idx_patients_identity_number ON patients(identity_number);
CREATE INDEX idx_patients_name ON patients(name);
CREATE INDEX idx_patients_phone_number ON patients(phone_number);