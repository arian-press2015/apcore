-- Create admins table
CREATE TABLE admins (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP,
    username VARCHAR UNIQUE NOT NULL,
    phone VARCHAR(11) UNIQUE NOT NULL,
    profile_image VARCHAR,
    password VARCHAR NOT NULL,
    totp_secret VARCHAR
);

-- Create customers table
CREATE TABLE customers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP,
    name VARCHAR NOT NULL,
    details TEXT,
    phone VARCHAR UNIQUE NOT NULL,
    logo VARCHAR,
    is_active BOOLEAN DEFAULT false,
    is_disabled BOOLEAN DEFAULT false,
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL
);

-- Create features table
CREATE TABLE features (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP,
    name VARCHAR UNIQUE NOT NULL,
    description TEXT NOT NULL,
    enabled BOOLEAN NOT NULL
);
