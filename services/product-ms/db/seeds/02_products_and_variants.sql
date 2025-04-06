-- Insert sample products
INSERT INTO products (
    id, store_id, name, slug, description, type,
    status, price, compare_at_price, cost_price, sku,
    barcode, weight, weight_unit, is_taxable, is_featured,
    is_gift_card, requires_shipping, inventory_quantity,
    inventory_policy, inventory_tracking, seo_title,
    seo_description, metadata
) VALUES
    -- Tech Haven products
    ('550e8400-e29b-41d4-a716-446655440010', '550e8400-e29b-41d4-a716-446655440000',
     'Gaming Laptop Pro', 'gaming-laptop-pro',
     'High-performance gaming laptop with RTX 3080', 'simple',
     'active', 1999.99, 2299.99, 1500.00, 'GLP-001',
     '123456789012', 2.5, 'kg', true, true,
     false, true, 10, 'deny', true,
     'Gaming Laptop Pro - Ultimate Gaming Experience',
     'Experience gaming like never before with our Pro gaming laptop',
     '{"brand": "TechMaster", "warranty": "2 years"}'),

    ('550e8400-e29b-41d4-a716-446655440011', '550e8400-e29b-41d4-a716-446655440000',
     'Business Ultrabook', 'business-ultrabook',
     'Sleek and powerful laptop for professionals', 'simple',
     'active', 1299.99, 1499.99, 1000.00, 'BU-001',
     '123456789013', 1.8, 'kg', true, false,
     false, true, 15, 'deny', true,
     'Business Ultrabook - Professional Performance',
     'Perfect for business professionals on the go',
     '{"brand": "TechMaster", "warranty": "3 years"}'),

    ('550e8400-e29b-41d4-a716-446655440012', '550e8400-e29b-41d4-a716-446655440000',
     'Smartphone X', 'smartphone-x',
     'Latest smartphone with advanced features', 'variable',
     'active', 999.99, 1199.99, 700.00, 'SPX-001',
     '123456789014', 0.2, 'kg', true, true,
     false, true, 20, 'deny', true,
     'Smartphone X - Next Generation Mobile',
     'Experience the future of mobile technology',
     '{"brand": "TechMaster", "warranty": "1 year"}');

-- Associate products with categories
INSERT INTO product_categories (product_id, category_id) VALUES
    ('550e8400-e29b-41d4-a716-446655440010', '550e8400-e29b-41d4-a716-446655440006'), -- Gaming Laptop Pro -> Gaming Laptops
    ('550e8400-e29b-41d4-a716-446655440011', '550e8400-e29b-41d4-a716-446655440004'), -- Business Ultrabook -> Laptops
    ('550e8400-e29b-41d4-a716-446655440012', '550e8400-e29b-41d4-a716-446655440005'); -- Smartphone X -> Smartphones

-- Insert variants for Smartphone X
INSERT INTO product_variants (
    id, product_id, name, sku, barcode, price,
    compare_at_price, cost_price, weight, weight_unit,
    inventory_quantity, inventory_policy, inventory_tracking,
    option_values
) VALUES
    ('550e8400-e29b-41d4-a716-446655440013', '550e8400-e29b-41d4-a716-446655440012',
     '128GB Black', 'SPX-001-BLK-128', '123456789015',
     999.99, 1199.99, 700.00, 0.2, 'kg',
     10, 'deny', true,
     '{"color": "black", "storage": "128GB"}'),

    ('550e8400-e29b-41d4-a716-446655440014', '550e8400-e29b-41d4-a716-446655440012',
     '256GB Black', 'SPX-001-BLK-256', '123456789016',
     1099.99, 1299.99, 750.00, 0.2, 'kg',
     8, 'deny', true,
     '{"color": "black", "storage": "256GB"}'),

    ('550e8400-e29b-41d4-a716-446655440015', '550e8400-e29b-41d4-a716-446655440012',
     '128GB White', 'SPX-001-WHT-128', '123456789017',
     999.99, 1199.99, 700.00, 0.2, 'kg',
     12, 'deny', true,
     '{"color": "white", "storage": "128GB"}');

-- Insert product images
INSERT INTO product_images (
    id, product_id, variant_id, url, alt_text,
    type, sort_order
) VALUES
    -- Gaming Laptop Pro images
    ('550e8400-e29b-41d4-a716-446655440016', '550e8400-e29b-41d4-a716-446655440010', NULL,
     'https://example.com/gaming-laptop-1.jpg', 'Gaming Laptop Pro Front View',
     'main', 1),

    ('550e8400-e29b-41d4-a716-446655440017', '550e8400-e29b-41d4-a716-446655440010', NULL,
     'https://example.com/gaming-laptop-2.jpg', 'Gaming Laptop Pro Back View',
     'gallery', 2),

    -- Business Ultrabook images
    ('550e8400-e29b-41d4-a716-446655440018', '550e8400-e29b-41d4-a716-446655440011', NULL,
     'https://example.com/business-laptop-1.jpg', 'Business Ultrabook Front View',
     'main', 1),

    -- Smartphone X images
    ('550e8400-e29b-41d4-a716-446655440019', '550e8400-e29b-41d4-a716-446655440012', NULL,
     'https://example.com/smartphone-1.jpg', 'Smartphone X Front View',
     'main', 1),

    ('550e8400-e29b-41d4-a716-446655440020', '550e8400-e29b-41d4-a716-446655440012', NULL,
     'https://example.com/smartphone-2.jpg', 'Smartphone X Back View',
     'gallery', 2),

    -- Variant-specific images
    ('550e8400-e29b-41d4-a716-446655440021', '550e8400-e29b-41d4-a716-446655440012', '550e8400-e29b-41d4-a716-446655440013',
     'https://example.com/smartphone-black-128.jpg', 'Smartphone X Black 128GB',
     'thumbnail', 1),

    ('550e8400-e29b-41d4-a716-446655440022', '550e8400-e29b-41d4-a716-446655440012', '550e8400-e29b-41d4-a716-446655440015',
     'https://example.com/smartphone-white-128.jpg', 'Smartphone X White 128GB',
     'thumbnail', 1);
