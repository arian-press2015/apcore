CREATE TABLE comments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP,
    title VARCHAR NOT NULL,
    body TEXT NOT NULL,
    rate INT NOT NULL,
    product_id UUID NOT NULL,
    CONSTRAINT fk_product
        FOREIGN KEY (product_id) 
        REFERENCES products(id)
        ON DELETE CASCADE
);