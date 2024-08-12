BEGIN;

-- Remove unique constraint from categories (customer_id, slug)
ALTER TABLE categories
DROP CONSTRAINT IF EXISTS fk_category_customer,
DROP CONSTRAINT IF EXISTS category_unique_slug_per_customer;
DROP COLUMN IF EXISTS customer_id;

-- Remove unique constraint from category_templates (name)
ALTER TABLE category_templates
DROP CONSTRAINT IF EXISTS fk_category_template_customer,
DROP CONSTRAINT IF EXISTS category_template_unique_slug_per_customer;
DROP COLUMN IF EXISTS customer_id;

-- Remove foreign key, unique constraint, and customer_id column from products
ALTER TABLE products
DROP CONSTRAINT IF EXISTS fk_product_customer,
DROP CONSTRAINT IF EXISTS product_unique_slug_per_customer,
DROP COLUMN IF EXISTS customer_id;

-- Remove foreign key, unique constraint, and customer_id column from product_templates
ALTER TABLE product_templates
DROP CONSTRAINT IF EXISTS fk_template_product_customer,
DROP CONSTRAINT IF EXISTS product_template_unique_slug_per_customer,
DROP COLUMN IF EXISTS customer_id;

COMMIT;
