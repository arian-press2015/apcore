BEGIN;

-- Add unique constraint to categories (customer_id, slug)
ALTER TABLE categories
ADD COLUMN customer_id UUID NOT NULL,
ADD CONSTRAINT fk_category_customer FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE,
ADD CONSTRAINT category_unique_slug_per_customer UNIQUE (customer_id, slug);

-- Add unique constraint to category_templates (name)
ALTER TABLE category_templates
ADD COLUMN customer_id UUID NOT NULL,
ADD CONSTRAINT fk_category_template_customer FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE,
ADD CONSTRAINT category_template_unique_slug_per_customer UNIQUE (name);

-- Add customer_id column to products, foreign key, and unique constraint (customer_id, slug)
ALTER TABLE products
ADD COLUMN customer_id UUID NOT NULL,
ADD CONSTRAINT fk_product_customer FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE,
ADD CONSTRAINT product_unique_slug_per_customer UNIQUE (customer_id, slug);

-- Add customer_id column to product_templates, foreign key, and unique constraint (customer_id, slug)
ALTER TABLE product_templates
ADD COLUMN customer_id UUID NOT NULL,
ADD CONSTRAINT fk_template_product_customer FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE,
ADD CONSTRAINT product_template_unique_slug_per_customer UNIQUE (customer_id, slug);

COMMIT;
