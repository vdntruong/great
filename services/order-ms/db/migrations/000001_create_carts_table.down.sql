-- Drop triggers
DROP TRIGGER IF EXISTS update_cart_items_updated_at ON cart_items;
DROP TRIGGER IF EXISTS update_carts_updated_at ON carts;

-- Drop tables
DROP TABLE IF EXISTS cart_items;
DROP TABLE IF EXISTS carts;
