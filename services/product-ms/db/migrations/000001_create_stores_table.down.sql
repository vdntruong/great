-- Drop triggers
DROP TRIGGER IF EXISTS update_stores_updated_at ON stores;

-- Drop tables
DROP TABLE IF EXISTS store_categories;
DROP TABLE IF EXISTS stores;

-- Drop types
DROP TYPE IF EXISTS store_status;
