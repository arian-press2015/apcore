-- Create users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP,
    full_name VARCHAR NOT NULL,
    phone VARCHAR(11) UNIQUE NOT NULL,
    profile_image VARCHAR,
    nid VARCHAR(10) UNIQUE,
    verified BOOLEAN DEFAULT false
);

-- Create roles table
CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP,
    name VARCHAR UNIQUE NOT NULL
);

-- Create user_roles table for many-to-many relationship between users and roles
CREATE TABLE user_roles (
    user_id UUID NOT NULL,
    role_id UUID NOT NULL,
    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
);

-- Create notifications table
CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP,
    recipient UUID NOT NULL,
    subject VARCHAR(20) NOT NULL,
    message TEXT NOT NULL,
    method VARCHAR(20) NOT NULL,
    is_read BOOLEAN DEFAULT false,
    FOREIGN KEY (recipient) REFERENCES users(id) ON DELETE CASCADE
);

-- Create notification_subscriptions table
CREATE TABLE notification_subscriptions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP,
    user_id UUID NOT NULL,
    method VARCHAR(20) NOT NULL,
    subject_type VARCHAR(20) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
