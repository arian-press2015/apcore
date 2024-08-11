CREATE TABLE ingredients (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    image_address VARCHAR(255) NOT NULL
);

CREATE TABLE product_ingredients (
    product_id UUID NOT NULL,
    ingredient_id UUID NOT NULL,
    PRIMARY KEY (product_id, ingredient_id),
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
    FOREIGN KEY (ingredient_id) REFERENCES ingredients(id) ON DELETE CASCADE
);

CREATE TABLE product_template_ingredients (
    product_template_id UUID NOT NULL,
    ingredient_id UUID NOT NULL,
    PRIMARY KEY (product_template_id, ingredient_id),
    FOREIGN KEY (product_template_id) REFERENCES product_templates(id) ON DELETE CASCADE,
    FOREIGN KEY (ingredient_id) REFERENCES ingredients(id) ON DELETE CASCADE
);