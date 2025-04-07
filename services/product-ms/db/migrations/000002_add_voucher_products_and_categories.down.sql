-- Drop indexes
DROP INDEX IF EXISTS idx_voucher_categories_category_id;
DROP INDEX IF EXISTS idx_voucher_categories_voucher_id;
DROP INDEX IF EXISTS idx_voucher_products_product_id;
DROP INDEX IF EXISTS idx_voucher_products_voucher_id;

-- Drop tables
DROP TABLE IF EXISTS voucher_categories;
DROP TABLE IF EXISTS voucher_products;
