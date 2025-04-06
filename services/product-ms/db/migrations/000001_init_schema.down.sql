-- Drop triggers
DROP TRIGGER IF EXISTS update_vouchers_updated_at ON vouchers;
DROP TRIGGER IF EXISTS update_discounts_updated_at ON discounts;
DROP TRIGGER IF EXISTS update_product_images_updated_at ON product_images;
DROP TRIGGER IF EXISTS update_product_variants_updated_at ON product_variants;
DROP TRIGGER IF EXISTS update_products_updated_at ON products;
DROP TRIGGER IF EXISTS update_stores_updated_at ON stores;

-- Drop function
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Drop tables
DROP TABLE IF EXISTS discount_categories;
DROP TABLE IF EXISTS discount_products;
DROP TABLE IF EXISTS vouchers;
DROP TABLE IF EXISTS discounts;
DROP TABLE IF EXISTS product_images;
DROP TABLE IF EXISTS product_variants;
DROP TABLE IF EXISTS product_categories;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS store_categories;
DROP TABLE IF EXISTS stores;

-- Drop types
DROP TYPE IF EXISTS voucher_status;
DROP TYPE IF EXISTS voucher_type;
DROP TYPE IF EXISTS discount_scope;
DROP TYPE IF EXISTS discount_type;
DROP TYPE IF EXISTS image_type;
DROP TYPE IF EXISTS product_type;
DROP TYPE IF EXISTS product_status;
DROP TYPE IF EXISTS store_status;
