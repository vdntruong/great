-- Drop triggers
DROP TRIGGER IF EXISTS update_order_items_updated_at ON order_items;
DROP TRIGGER IF EXISTS update_orders_updated_at ON orders;

-- Drop tables
DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;
