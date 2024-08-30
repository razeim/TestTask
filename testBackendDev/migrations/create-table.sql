CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),  
    email VARCHAR(255) UNIQUE NOT NULL,             
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP  
);

CREATE TABLE IF NOT EXISTS tokens (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    refresh_token TEXT NOT NULL,
    user_ip TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMPTZ NOT NULL
);

INSERT INTO users (id, email) VALUES
('e4b7b1f8-f4d4-4e1a-a7cb-65e0a2d59c92', '1@gmail.com'),
('a9b3c1d7-4e1b-45d7-8c52-6f3d4a6c9e10', '2@gmail.com'),
('f74e6a3d-9c9e-4c4e-b8de-04e3f94f5d88', '3@gmail.com');
('d8e1e0b5-f9d6-4b0e-b6f5-2f7f9c8e1b40', '4@gmail.com');
('9c8a0d1b-2e4e-45cf-9f62-d9d89a9e6d6b', '5@gmail.com');