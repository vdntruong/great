-- Insert sample products with various statuses and types
INSERT INTO products (
    id, store_id, name, slug, description, type,
    status, price, compare_at_price, cost_price, sku,
    barcode, weight, weight_unit, is_taxable, is_featured,
    is_gift_card, requires_shipping, inventory_quantity,
    inventory_policy, inventory_tracking, seo_title,
    seo_description, metadata, created_at, updated_at, deleted_at
) VALUES
    -- Tech Haven products
    ('550e8400-e29b-41d4-a716-446655440017', '550e8400-e29b-41d4-a716-446655440000',
     'Gaming Laptop Pro', 'gaming-laptop-pro',
     'High-performance gaming laptop with RTX 4080', 'simple',
     'active', 1999.99, 2299.99, 1500.00, 'TH-GLP-001',
     '123456789012', 2.5, 'kg', true, true,
     false, true, 10, 'deny', true,
     'Gaming Laptop Pro - Ultimate Gaming Experience',
     'Experience gaming like never before with our Pro gaming laptop',
     '{"specs": {"cpu": "Intel i9-13900H", "gpu": "RTX 4080", "ram": "32GB", "storage": "2TB SSD"}, "features": ["RGB Keyboard", "240Hz Display", "Advanced Cooling"]}',
     '2024-01-15 00:00:00+00', '2024-01-15 00:00:00+00', NULL),

    ('550e8400-e29b-41d4-a716-446655440018', '550e8400-e29b-41d4-a716-446655440000',
     'Business Ultrabook', 'business-ultrabook',
     'Professional laptop for business users', 'simple',
     'active', 1499.99, 1699.99, 1000.00, 'TH-BU-001',
     '123456789013', 1.3, 'kg', true, false,
     false, true, 15, 'deny', true,
     'Business Ultrabook - Professional Performance',
     'Perfect for business professionals on the go',
     '{"specs": {"cpu": "Intel i7-1360P", "gpu": "Iris Xe", "ram": "16GB", "storage": "1TB SSD"}, "features": ["Fingerprint Reader", "Backlit Keyboard", "Military Grade Durability"]}',
     '2024-02-01 00:00:00+00', '2024-02-01 00:00:00+00', NULL),

    ('550e8400-e29b-41d4-a716-446655440019', '550e8400-e29b-41d4-a716-446655440000',
     'Smartphone X', 'smartphone-x',
     'Latest flagship smartphone', 'variable',
     'active', 999.99, 1199.99, 700.00, 'TH-SX-001',
     '123456789014', 0.2, 'kg', true, true,
     false, true, 20, 'deny', true,
     'Smartphone X - Next Generation Mobile',
     'Experience the future of mobile technology',
     '{"specs": {"processor": "Snapdragon 8 Gen 2", "camera": "50MP Triple Camera", "battery": "5000mAh"}, "features": ["5G", "Wireless Charging", "IP68 Rating"]}',
     '2024-03-01 00:00:00+00', '2024-03-01 00:00:00+00', NULL),

    -- Draft product
    ('550e8400-e29b-41d4-a716-446655440020', '550e8400-e29b-41d4-a716-446655440000',
     'Upcoming Gaming Monitor', 'upcoming-gaming-monitor',
     'Next-gen gaming monitor', 'simple',
     'draft', 799.99, 899.99, 600.00, 'TH-GM-001',
     '123456789015', 8.0, 'kg', true, false,
     false, true, 0, 'deny', true,
     'Upcoming Gaming Monitor - 4K Gaming',
     'Prepare for the next level of gaming',
     '{"specs": {"size": "32-inch", "resolution": "4K", "refresh_rate": "240Hz", "panel": "OLED"}, "features": ["HDR", "G-Sync", "USB Hub"]}',
     '2024-04-01 00:00:00+00', '2024-04-01 00:00:00+00', NULL),

    -- Inactive product
    ('550e8400-e29b-41d4-a716-446655440021', '550e8400-e29b-41d4-a716-446655440000',
     'Old Model Laptop', 'old-model-laptop',
     'Previous generation laptop', 'simple',
     'inactive', 899.99, 1299.99, 700.00, 'TH-OML-001',
     '123456789016', 2.2, 'kg', true, false,
     false, true, 5, 'deny', true,
     'Old Model Laptop - Previous Generation',
     'Previous generation laptop at a discounted price',
     '{"specs": {"cpu": "Intel i7-1165G7", "gpu": "MX450", "ram": "16GB", "storage": "512GB SSD"}, "features": ["USB-C", "HDMI", "SD Card Reader"]}',
     '2023-12-01 00:00:00+00', '2024-01-01 00:00:00+00', NULL),

    -- Archived product
    ('550e8400-e29b-41d4-a716-446655440022', '550e8400-e29b-41d4-a716-446655440000',
     'Discontinued Phone', 'discontinued-phone',
     'Outdated smartphone model', 'simple',
     'archived', 299.99, 499.99, 400.00, 'TH-DP-001',
     '123456789017', 0.18, 'kg', true, false,
     false, true, 0, 'deny', true,
     'Discontinued Phone - No Longer Available',
     'This model is no longer in production',
     '{"specs": {"processor": "Snapdragon 855", "camera": "12MP Dual Camera", "battery": "4000mAh"}, "features": ["4G", "Fingerprint Sensor"]}',
     '2023-06-01 00:00:00+00', '2023-12-01 00:00:00+00', NULL);

-- Associate products with categories
INSERT INTO product_categories (product_id, category_id) VALUES
    ('550e8400-e29b-41d4-a716-446655440017', '550e8400-e29b-41d4-a716-446655440008'), -- Laptops
    ('550e8400-e29b-41d4-a716-446655440017', '550e8400-e29b-41d4-a716-446655440009'), -- Gaming Laptops
    ('550e8400-e29b-41d4-a716-446655440018', '550e8400-e29b-41d4-a716-446655440008'), -- Laptops
    ('550e8400-e29b-41d4-a716-446655440018', '550e8400-e29b-41d4-a716-446655440010'), -- Business Laptops
    ('550e8400-e29b-41d4-a716-446655440019', '550e8400-e29b-41d4-a716-446655440011'), -- Smartphones
    ('550e8400-e29b-41d4-a716-446655440020', '550e8400-e29b-41d4-a716-446655440012'), -- Gaming Accessories
    ('550e8400-e29b-41d4-a716-446655440021', '550e8400-e29b-41d4-a716-446655440008'), -- Laptops
    ('550e8400-e29b-41d4-a716-446655440022', '550e8400-e29b-41d4-a716-446655440011'); -- Smartphones

-- Insert variants for Smartphone X
INSERT INTO product_variants (
    id, product_id, name, sku, barcode, price,
    compare_at_price, cost_price, weight, weight_unit,
    inventory_quantity, inventory_policy, inventory_tracking,
    option_values, created_at, updated_at
) VALUES
    ('550e8400-e29b-41d4-a716-446655440023', '550e8400-e29b-41d4-a716-446655440019',
     '128GB Black', 'SPX-001-BLK-128', '123456789018',
     999.99, 1199.99, 700.00, 0.2, 'kg',
     10, 'deny', true,
     '{
        "color": "black",
        "storage": "128GB",
        "features": ["5G", "Wireless Charging"]
     }'::jsonb,
     '2024-01-03 00:00:00+00', '2024-01-03 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440024', '550e8400-e29b-41d4-a716-446655440019',
     '256GB Black', 'SPX-001-BLK-256', '123456789019',
     1099.99, 1299.99, 750.00, 0.2, 'kg',
     8, 'deny', true,
     '{
        "color": "black",
        "storage": "256GB",
        "features": ["5G", "Wireless Charging"]
     }'::jsonb,
     '2024-01-03 00:00:00+00', '2024-01-03 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440025', '550e8400-e29b-41d4-a716-446655440019',
     '128GB White', 'SPX-001-WHT-128', '123456789020',
     999.99, 1199.99, 700.00, 0.2, 'kg',
     12, 'deny', true,
     '{
        "color": "white",
        "storage": "128GB",
        "features": ["5G", "Wireless Charging"]
     }'::jsonb,
     '2024-01-03 00:00:00+00', '2024-01-03 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440026', '550e8400-e29b-41d4-a716-446655440019',
     '256GB White', 'SPX-001-WHT-256', '123456789021',
     1099.99, 1299.99, 750.00, 0.2, 'kg',
     6, 'deny', true,
     '{
        "color": "white",
        "storage": "256GB",
        "features": ["5G", "Wireless Charging"]
     }'::jsonb,
     '2024-01-03 00:00:00+00', '2024-01-03 00:00:00+00');

-- Insert product images
INSERT INTO product_images (
    id, product_id, variant_id, url, alt_text,
    type, sort_order, created_at, updated_at
) VALUES
    -- Gaming Laptop Pro images
    ('550e8400-e29b-41d4-a716-446655440027', '550e8400-e29b-41d4-a716-446655440017', NULL,
     'https://example.com/gaming-laptop-1.jpg', 'Gaming Laptop Pro Front View',
     'main', 1, '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440028', '550e8400-e29b-41d4-a716-446655440017', NULL,
     'https://example.com/gaming-laptop-2.jpg', 'Gaming Laptop Pro Back View',
     'gallery', 2, '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    -- Business Ultrabook images
    ('550e8400-e29b-41d4-a716-446655440029', '550e8400-e29b-41d4-a716-446655440018', NULL,
     'https://example.com/business-laptop-1.jpg', 'Business Ultrabook Front View',
     'main', 1, '2024-01-02 00:00:00+00', '2024-01-02 00:00:00+00'),

    -- Smartphone X images
    ('550e8400-e29b-41d4-a716-446655440030', '550e8400-e29b-41d4-a716-446655440019', NULL,
     'https://example.com/smartphone-1.jpg', 'Smartphone X Front View',
     'main', 1, '2024-01-03 00:00:00+00', '2024-01-03 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440031', '550e8400-e29b-41d4-a716-446655440019', NULL,
     'https://example.com/smartphone-2.jpg', 'Smartphone X Back View',
     'gallery', 2, '2024-01-03 00:00:00+00', '2024-01-03 00:00:00+00'),

    -- Variant-specific images
    ('550e8400-e29b-41d4-a716-446655440032', '550e8400-e29b-41d4-a716-446655440019', '550e8400-e29b-41d4-a716-446655440023',
     'https://example.com/smartphone-black-128.jpg', 'Smartphone X Black 128GB',
     'thumbnail', 1, '2024-01-03 00:00:00+00', '2024-01-03 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440033', '550e8400-e29b-41d4-a716-446655440019', '550e8400-e29b-41d4-a716-446655440025',
     'https://example.com/smartphone-white-128.jpg', 'Smartphone X White 128GB',
     'thumbnail', 1, '2024-01-03 00:00:00+00', '2024-01-03 00:00:00+00');
