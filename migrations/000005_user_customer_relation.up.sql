CREATE TABLE user_customers (
    user_id UUID NOT NULL,
    customer_id UUID NOT NULL,
    PRIMARY KEY (user_id, customer_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE
);