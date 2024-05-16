CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(26) PRIMARY KEY,
    nip BIGINT UNIQUE,
    name VARCHAR(50),
    role VARCHAR(5) CHECK (role IN ('it', 'nurse')),
    password VARCHAR(255),
    identity_image_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL
);

CREATE INDEX idx_users_id ON users(id);
CREATE INDEX idx_users_nip ON users(nip);
CREATE INDEX idx_users_name ON users(name);
CREATE INDEX idx_users_role ON users(role);