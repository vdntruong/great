-- Drop triggers
DROP TRIGGER IF EXISTS update_product_images_updated_at ON product_images;
DROP TRIGGER IF EXISTS update_product_variants_updated_at ON product_variants;
DROP TRIGGER IF EXISTS update_products_updated_at ON products;

-- Drop tables
DROP TABLE IF EXISTS product_images;
DROP TABLE IF EXISTS product_variants;
DROP TABLE IF EXISTS product_categories;
DROP TABLE IF EXISTS products;

-- Drop types
DROP TYPE IF EXISTS image_type;
DROP TYPE IF EXISTS product_type;
DROP TYPE IF EXISTS product_status;
