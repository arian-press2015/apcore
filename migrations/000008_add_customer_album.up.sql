CREATE TABLE customer_albums (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP,
    name VARCHAR(30) NOT NULL,
    owner_id UUID NOT NULL,
    address VARCHAR(255) NOT NULL,
    CONSTRAINT unique_name_owner UNIQUE (name, owner_id)
);