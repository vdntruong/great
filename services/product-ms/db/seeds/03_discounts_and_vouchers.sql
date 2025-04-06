-- Insert sample discounts
INSERT INTO discounts (
    id, store_id, name, code, type, value,
    scope, start_date, end_date, min_purchase_amount,
    max_discount_amount, usage_limit, is_active
) VALUES
    -- Tech Haven discounts
    ('550e8400-e29b-41d4-a716-446655440023', '550e8400-e29b-41d4-a716-446655440000',
     'Summer Sale', 'SUMMER20', 'percentage', 20.00,
     'all_products', '2024-06-01 00:00:00+00', '2024-08-31 23:59:59+00',
     100.00, 200.00, 1000, true),

    ('550e8400-e29b-41d4-a716-446655440024', '550e8400-e29b-41d4-a716-446655440000',
     'Gaming Laptops Special', 'GAMING50', 'fixed_amount', 50.00,
     'specific_products', '2024-07-01 00:00:00+00', '2024-07-31 23:59:59+00',
     500.00, 100.00, 100, true),

    ('550e8400-e29b-41d4-a716-446655440025', '550e8400-e29b-41d4-a716-446655440000',
     'New Year Sale', 'NEWYEAR25', 'percentage', 25.00,
     'specific_categories', '2024-12-26 00:00:00+00', '2025-01-15 23:59:59+00',
     200.00, 500.00, 500, true);

-- Associate discounts with products
INSERT INTO discount_products (discount_id, product_id) VALUES
    ('550e8400-e29b-41d4-a716-446655440024', '550e8400-e29b-41d4-a716-446655440010'); -- Gaming Laptops Special -> Gaming Laptop Pro

-- Associate discounts with categories
INSERT INTO discount_categories (discount_id, category_id) VALUES
    ('550e8400-e29b-41d4-a716-446655440025', '550e8400-e29b-41d4-a716-446655440004'), -- New Year Sale -> Laptops
    ('550e8400-e29b-41d4-a716-446655440025', '550e8400-e29b-41d4-a716-446655440005'); -- New Year Sale -> Smartphones

-- Insert sample vouchers
INSERT INTO vouchers (
    id, store_id, code, type, value, min_purchase_amount,
    max_discount_amount, start_date, end_date, usage_limit,
    status
) VALUES
    -- Tech Haven vouchers
    ('550e8400-e29b-41d4-a716-446655440026', '550e8400-e29b-41d4-a716-446655440000',
     'WELCOME10', 'percentage', 10.00, 50.00,
     100.00, '2024-01-01 00:00:00+00', '2024-12-31 23:59:59+00',
     1000, 'active'),

    ('550e8400-e29b-41d4-a716-446655440027', '550e8400-e29b-41d4-a716-446655440000',
     'FREESHIP', 'free_shipping', NULL, 100.00,
     NULL, '2024-01-01 00:00:00+00', '2024-12-31 23:59:59+00',
     500, 'active'),

    ('550e8400-e29b-41d4-a716-446655440028', '550e8400-e29b-41d4-a716-446655440000',
     'FIRSTORDER', 'fixed_amount', 20.00, 100.00,
     20.00, '2024-01-01 00:00:00+00', '2024-12-31 23:59:59+00',
     1000, 'active');
