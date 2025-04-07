-- Drop triggers
DROP TRIGGER IF EXISTS update_discounts_updated_at ON discounts;

-- Drop tables
DROP TABLE IF EXISTS discount_categories;
DROP TABLE IF EXISTS discount_products;
DROP TABLE IF EXISTS discounts;

-- Drop types
DROP TYPE IF EXISTS discount_scope;
DROP TYPE IF EXISTS discount_type;
