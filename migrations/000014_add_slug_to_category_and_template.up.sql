ALTER TABLE categories
ADD COLUMN slug VARCHAR(30) NOT NULL;

ALTER TABLE category_templates
ADD COLUMN slug VARCHAR(30) NOT NULL;