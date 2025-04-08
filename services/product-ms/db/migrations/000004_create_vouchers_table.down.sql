-- Drop indexes
DROP INDEX IF EXISTS idx_voucher_categories_category_id;
DROP INDEX IF EXISTS idx_voucher_categories_voucher_id;
DROP INDEX IF EXISTS idx_voucher_products_product_id;
DROP INDEX IF EXISTS idx_voucher_products_voucher_id;

-- Drop triggers
DROP TRIGGER IF EXISTS update_vouchers_updated_at ON vouchers;

-- Drop tables
DROP TABLE IF EXISTS voucher_categories;
DROP TABLE IF EXISTS voucher_products;
DROP TABLE IF EXISTS vouchers;

-- Drop types
DROP TYPE IF EXISTS voucher_status;
DROP TYPE IF EXISTS voucher_type;
