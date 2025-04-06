-- Insert sample discounts with various types and scopes
INSERT INTO discounts (
    id, store_id, name, code, type, value,
    scope, start_date, end_date, min_purchase_amount,
    max_discount_amount, usage_limit, usage_count, is_active,
    created_at, updated_at
) VALUES
    -- Tech Haven discounts
    ('550e8400-e29b-41d4-a716-446655440034', '550e8400-e29b-41d4-a716-446655440000',
     'Summer Sale', 'SUMMER20', 'percentage', 20.00,
     'all_products', '2024-06-01 00:00:00+00', '2024-08-31 23:59:59+00',
     100.00, 200.00, 1000, 250, true,
     '2024-05-01 00:00:00+00', '2024-06-15 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440035', '550e8400-e29b-41d4-a716-446655440000',
     'Gaming Laptops Special', 'GAMING50', 'fixed_amount', 50.00,
     'specific_products', '2024-07-01 00:00:00+00', '2024-07-31 23:59:59+00',
     500.00, 100.00, 100, 45, true,
     '2024-06-15 00:00:00+00', '2024-06-15 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440036', '550e8400-e29b-41d4-a716-446655440000',
     'New Year Sale', 'NEWYEAR25', 'percentage', 25.00,
     'specific_categories', '2024-12-26 00:00:00+00', '2025-01-15 23:59:59+00',
     200.00, 500.00, 500, 0, true,
     '2024-12-01 00:00:00+00', '2024-12-01 00:00:00+00'),

    -- Inactive discount
    ('550e8400-e29b-41d4-a716-446655440037', '550e8400-e29b-41d4-a716-446655440000',
     'Expired Sale', 'EXPIRED10', 'percentage', 10.00,
     'all_products', '2023-12-01 00:00:00+00', '2023-12-31 23:59:59+00',
     50.00, 100.00, 1000, 1000, false,
     '2023-11-15 00:00:00+00', '2024-01-01 00:00:00+00'),

    -- Fashion Forward discounts
    ('550e8400-e29b-41d4-a716-446655440038', '550e8400-e29b-41d4-a716-446655440002',
     'Spring Collection', 'SPRING15', 'percentage', 15.00,
     'all_products', '2024-03-01 00:00:00+00', '2024-05-31 23:59:59+00',
     75.00, 150.00, 500, 120, true,
     '2024-02-15 00:00:00+00', '2024-03-15 00:00:00+00');

-- Associate discounts with products
INSERT INTO discount_products (discount_id, product_id) VALUES
    ('550e8400-e29b-41d4-a716-446655440035', '550e8400-e29b-41d4-a716-446655440017'); -- Gaming Laptops Special -> Gaming Laptop Pro

-- Associate discounts with categories
INSERT INTO discount_categories (discount_id, category_id) VALUES
    ('550e8400-e29b-41d4-a716-446655440036', '550e8400-e29b-41d4-a716-446655440008'), -- New Year Sale -> Laptops
    ('550e8400-e29b-41d4-a716-446655440036', '550e8400-e29b-41d4-a716-446655440011'); -- New Year Sale -> Smartphones

-- Insert sample vouchers with various types and statuses
INSERT INTO vouchers (
    id, store_id, code, type, value, min_purchase_amount,
    max_discount_amount, start_date, end_date, usage_limit,
    usage_count, status, created_at, updated_at
) VALUES
    -- Tech Haven vouchers
    ('550e8400-e29b-41d4-a716-446655440039', '550e8400-e29b-41d4-a716-446655440000',
     'WELCOME10', 'percentage', 10.00, 50.00,
     100.00, '2024-01-01 00:00:00+00', '2024-12-31 23:59:59+00',
     1000, 450, 'active',
     '2023-12-15 00:00:00+00', '2024-06-15 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440040', '550e8400-e29b-41d4-a716-446655440000',
     'FREESHIP', 'free_shipping', NULL, 100.00,
     NULL, '2024-01-01 00:00:00+00', '2024-12-31 23:59:59+00',
     500, 120, 'active',
     '2023-12-15 00:00:00+00', '2024-06-15 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440041', '550e8400-e29b-41d4-a716-446655440000',
     'FIRSTORDER', 'fixed_amount', 20.00, 100.00,
     20.00, '2024-01-01 00:00:00+00', '2024-12-31 23:59:59+00',
     1000, 800, 'active',
     '2023-12-15 00:00:00+00', '2024-06-15 00:00:00+00'),

    -- Expired voucher
    ('550e8400-e29b-41d4-a716-446655440042', '550e8400-e29b-41d4-a716-446655440000',
     'HOLIDAY2023', 'percentage', 15.00, 75.00,
     150.00, '2023-12-01 00:00:00+00', '2023-12-31 23:59:59+00',
     500, 500, 'expired',
     '2023-11-15 00:00:00+00', '2024-01-01 00:00:00+00'),

    -- Inactive voucher
    ('550e8400-e29b-41d4-a716-446655440043', '550e8400-e29b-41d4-a716-446655440000',
     'TESTCODE', 'fixed_amount', 5.00, 25.00,
     5.00, '2024-01-01 00:00:00+00', '2024-12-31 23:59:59+00',
     100, 0, 'inactive',
     '2023-12-15 00:00:00+00', '2023-12-15 00:00:00+00'),

    -- Fashion Forward vouchers
    ('550e8400-e29b-41d4-a716-446655440044', '550e8400-e29b-41d4-a716-446655440002',
     'FASHION20', 'percentage', 20.00, 100.00,
     200.00, '2024-01-01 00:00:00+00', '2024-12-31 23:59:59+00',
     1000, 300, 'active',
     '2023-12-15 00:00:00+00', '2024-06-15 00:00:00+00');
