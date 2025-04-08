-- Insert sample products for Tech Haven
INSERT INTO products (
    id, store_id, name, slug, description, type, status,
    price, compare_at_price, cost_price, sku, barcode,
    weight, weight_unit, is_taxable, is_featured, is_gift_card,
    requires_shipping, inventory_quantity, inventory_policy,
    inventory_tracking, seo_title, seo_description,
    created_at, updated_at, deleted_at
) VALUES
    -- Tech Haven products
    ('550e8400-e29b-41d4-a716-446655440022', '550e8400-e29b-41d4-a716-446655440000',
     'MacBook Pro M3', 'macbook-pro-m3',
     'The most powerful MacBook Pro ever with M3 chip', 'simple',
     'active', 1999.99, 2199.99, 1500.00, 'MBP-M3-001', '123456789012',
     2.0, 'kg', true, true, false, true, 50, 'deny',
     true, 'MacBook Pro M3 - Tech Haven', 'Get the new MacBook Pro M3 with amazing performance',
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00', NULL),

    ('550e8400-e29b-41d4-a716-446655440023', '550e8400-e29b-41d4-a716-446655440000',
     'iPhone 15 Pro', 'iphone-15-pro',
     'The most advanced iPhone with A17 Pro chip', 'simple',
     'active', 999.99, 1099.99, 800.00, 'IP15P-001', '234567890123',
     0.2, 'kg', true, true, false, true, 100, 'deny',
     true, 'iPhone 15 Pro - Tech Haven', 'Experience the new iPhone 15 Pro with revolutionary features',
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00', NULL),

    ('550e8400-e29b-41d4-a716-446655440024', '550e8400-e29b-41d4-a716-446655440000',
     'Gaming Laptop', 'gaming-laptop',
     'High-performance gaming laptop with RTX 4090', 'variable',
     'active', 2499.99, 2699.99, 2000.00, 'GL-001', '345678901234',
     2.5, 'kg', true, true, false, true, 30, 'deny',
     true, 'Gaming Laptop - Tech Haven', 'Ultimate gaming experience with RTX 4090',
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00', NULL),

    -- Fashion Forward products
    ('550e8400-e29b-41d4-a716-446655440025', '550e8400-e29b-41d4-a716-446655440002',
     'Leather Jacket', 'leather-jacket',
     'Premium quality leather jacket', 'simple',
     'active', 299.99, 399.99, 200.00, 'LJ-001', '456789012345',
     1.0, 'kg', true, true, false, true, 20, 'deny',
     true, 'Leather Jacket - Fashion Forward', 'Stylish leather jacket for any occasion',
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00', NULL),

    ('550e8400-e29b-41d4-a716-446655440026', '550e8400-e29b-41d4-a716-446655440002',
     'Summer Dress', 'summer-dress',
     'Light and comfortable summer dress', 'variable',
     'active', 79.99, 99.99, 50.00, 'SD-001', '567890123456',
     0.3, 'kg', true, true, false, true, 50, 'deny',
     true, 'Summer Dress - Fashion Forward', 'Perfect summer dress for warm days',
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00', NULL),

    -- Home Decor products
    ('550e8400-e29b-41d4-a716-446655440027', '550e8400-e29b-41d4-a716-446655440004',
     'Modern Sofa', 'modern-sofa',
     'Contemporary design sofa with premium fabric', 'simple',
     'active', 899.99, 999.99, 600.00, 'MS-001', '678901234567',
     30.0, 'kg', true, true, false, true, 10, 'deny',
     true, 'Modern Sofa - Home Decor', 'Stylish and comfortable modern sofa',
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00', NULL),

    ('550e8400-e29b-41d4-a716-446655440028', '550e8400-e29b-41d4-a716-446655440004',
     'Wall Art Set', 'wall-art-set',
     'Set of 3 modern wall art pieces', 'simple',
     'active', 149.99, 199.99, 100.00, 'WAS-001', '789012345678',
     2.0, 'kg', true, true, false, true, 25, 'deny',
     true, 'Wall Art Set - Home Decor', 'Beautiful wall art set for your home',
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00', NULL),

    -- Beauty Haven products
    ('550e8400-e29b-41d4-a716-446655440029', '550e8400-e29b-41d4-a716-446655440006',
     'Vitamin C Serum', 'vitamin-c-serum',
     'Brightening and anti-aging serum', 'simple',
     'active', 49.99, 59.99, 30.00, 'VCS-001', '890123456789',
     0.1, 'kg', true, true, false, true, 100, 'deny',
     true, 'Vitamin C Serum - Beauty Haven', 'Premium vitamin C serum for glowing skin',
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00', NULL),

    ('550e8400-e29b-41d4-a716-446655440030', '550e8400-e29b-41d4-a716-446655440006',
     'Makeup Palette', 'makeup-palette',
     'Professional makeup palette with 12 shades', 'simple',
     'active', 39.99, 49.99, 25.00, 'MP-001', '901234567890',
     0.2, 'kg', true, true, false, true, 75, 'deny',
     true, 'Makeup Palette - Beauty Haven', 'Versatile makeup palette for all occasions',
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00', NULL);

-- Insert product variants
INSERT INTO product_variants (
    id, product_id, name, sku, barcode, price, compare_at_price,
    cost_price, weight, weight_unit, is_taxable, requires_shipping,
    inventory_quantity, inventory_policy, inventory_tracking,
    created_at, updated_at
) VALUES
    -- Gaming Laptop variants
    ('550e8400-e29b-41d4-a716-446655440031', '550e8400-e29b-41d4-a716-446655440024',
     '16GB RAM, 1TB SSD', 'GL-001-16-1TB', '345678901234-1',
     2499.99, 2699.99, 2000.00, 2.5, 'kg', true, true,
     15, 'deny', true,
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440032', '550e8400-e29b-41d4-a716-446655440024',
     '32GB RAM, 2TB SSD', 'GL-001-32-2TB', '345678901234-2',
     2999.99, 3199.99, 2400.00, 2.5, 'kg', true, true,
     15, 'deny', true,
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    -- Summer Dress variants
    ('550e8400-e29b-41d4-a716-446655440033', '550e8400-e29b-41d4-a716-446655440026',
     'Small', 'SD-001-S', '567890123456-S',
     79.99, 99.99, 50.00, 0.3, 'kg', true, true,
     10, 'deny', true,
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440034', '550e8400-e29b-41d4-a716-446655440026',
     'Medium', 'SD-001-M', '567890123456-M',
     79.99, 99.99, 50.00, 0.3, 'kg', true, true,
     20, 'deny', true,
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440035', '550e8400-e29b-41d4-a716-446655440026',
     'Large', 'SD-001-L', '567890123456-L',
     79.99, 99.99, 50.00, 0.3, 'kg', true, true,
     20, 'deny', true,
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00');

-- Insert product images
INSERT INTO product_images (
    id, product_id, url, alt_text, position, is_main,
    created_at, updated_at
) VALUES
    -- MacBook Pro M3 images
    ('550e8400-e29b-41d4-a716-446655440036', '550e8400-e29b-41d4-a716-446655440022',
     'https://example.com/macbook-pro-m3-1.jpg', 'MacBook Pro M3 front view',
     1, true, '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440037', '550e8400-e29b-41d4-a716-446655440022',
     'https://example.com/macbook-pro-m3-2.jpg', 'MacBook Pro M3 side view',
     2, false, '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    -- iPhone 15 Pro images
    ('550e8400-e29b-41d4-a716-446655440038', '550e8400-e29b-41d4-a716-446655440023',
     'https://example.com/iphone-15-pro-1.jpg', 'iPhone 15 Pro front view',
     1, true, '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440039', '550e8400-e29b-41d4-a716-446655440023',
     'https://example.com/iphone-15-pro-2.jpg', 'iPhone 15 Pro back view',
     2, false, '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    -- Leather Jacket images
    ('550e8400-e29b-41d4-a716-446655440040', '550e8400-e29b-41d4-a716-446655440025',
     'https://example.com/leather-jacket-1.jpg', 'Leather Jacket front view',
     1, true, '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440041', '550e8400-e29b-41d4-a716-446655440025',
     'https://example.com/leather-jacket-2.jpg', 'Leather Jacket back view',
     2, false, '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    -- Summer Dress images
    ('550e8400-e29b-41d4-a716-446655440042', '550e8400-e29b-41d4-a716-446655440026',
     'https://example.com/summer-dress-1.jpg', 'Summer Dress front view',
     1, true, '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440043', '550e8400-e29b-41d4-a716-446655440026',
     'https://example.com/summer-dress-2.jpg', 'Summer Dress back view',
     2, false, '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    -- Modern Sofa images
    ('550e8400-e29b-41d4-a716-446655440044', '550e8400-e29b-41d4-a716-446655440027',
     'https://example.com/modern-sofa-1.jpg', 'Modern Sofa front view',
     1, true, '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440045', '550e8400-e29b-41d4-a716-446655440027',
     'https://example.com/modern-sofa-2.jpg', 'Modern Sofa side view',
     2, false, '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    -- Wall Art Set images
    ('550e8400-e29b-41d4-a716-446655440046', '550e8400-e29b-41d4-a716-446655440028',
     'https://example.com/wall-art-set-1.jpg', 'Wall Art Set complete view',
     1, true, '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440047', '550e8400-e29b-41d4-a716-446655440028',
     'https://example.com/wall-art-set-2.jpg', 'Wall Art Set individual pieces',
     2, false, '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    -- Vitamin C Serum images
    ('550e8400-e29b-41d4-a716-446655440048', '550e8400-e29b-41d4-a716-446655440029',
     'https://example.com/vitamin-c-serum-1.jpg', 'Vitamin C Serum bottle',
     1, true, '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440049', '550e8400-e29b-41d4-a716-446655440029',
     'https://example.com/vitamin-c-serum-2.jpg', 'Vitamin C Serum ingredients',
     2, false, '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    -- Makeup Palette images
    ('550e8400-e29b-41d4-a716-446655440050', '550e8400-e29b-41d4-a716-446655440030',
     'https://example.com/makeup-palette-1.jpg', 'Makeup Palette closed',
     1, true, '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440051', '550e8400-e29b-41d4-a716-446655440030',
     'https://example.com/makeup-palette-2.jpg', 'Makeup Palette open',
     2, false, '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00');

-- Insert product category associations
INSERT INTO product_categories (
    product_id, category_id, created_at, updated_at
) VALUES
    -- Tech Haven products
    ('550e8400-e29b-41d4-a716-446655440022', '550e8400-e29b-41d4-a716-446655440008',
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),
    ('550e8400-e29b-41d4-a716-446655440022', '550e8400-e29b-41d4-a716-446655440010',
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),
    ('550e8400-e29b-41d4-a716-446655440023', '550e8400-e29b-41d4-a716-446655440011',
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),
    ('550e8400-e29b-41d4-a716-446655440024', '550e8400-e29b-41d4-a716-446655440009',
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    -- Fashion Forward products
    ('550e8400-e29b-41d4-a716-446655440025', '550e8400-e29b-41d4-a716-446655440013',
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),
    ('550e8400-e29b-41d4-a716-446655440026', '550e8400-e29b-41d4-a716-446655440014',
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    -- Home Decor products
    ('550e8400-e29b-41d4-a716-446655440027', '550e8400-e29b-41d4-a716-446655440016',
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),
    ('550e8400-e29b-41d4-a716-446655440027', '550e8400-e29b-41d4-a716-446655440018',
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),
    ('550e8400-e29b-41d4-a716-446655440028', '550e8400-e29b-41d4-a716-446655440017',
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    -- Beauty Haven products
    ('550e8400-e29b-41d4-a716-446655440029', '550e8400-e29b-41d4-a716-446655440019',
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),
    ('550e8400-e29b-41d4-a716-446655440030', '550e8400-e29b-41d4-a716-446655440020',
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00');
