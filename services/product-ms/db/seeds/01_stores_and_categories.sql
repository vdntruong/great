-- Insert sample stores with various statuses and settings
INSERT INTO stores (
    id, name, slug, description, logo_url, cover_url,
    status, is_verified, owner_id, contact_email, contact_phone,
    address, settings, created_at, updated_at, deleted_at
) VALUES
    -- Active and verified store
    ('550e8400-e29b-41d4-a716-446655440000', 'Tech Haven', 'tech-haven',
     'Your one-stop shop for all things tech', 'https://example.com/tech-haven-logo.png',
     'https://example.com/tech-haven-cover.png', 'active', true,
     '550e8400-e29b-41d4-a716-446655440001', 'contact@techhaven.com', '+1234567890',
     '123 Tech Street, Silicon Valley',
     '{
        "theme": "dark",
        "currency": "USD",
        "tax_settings": {
            "enabled": true,
            "rate": 0.08,
            "inclusive": false
        },
        "shipping": {
            "free_threshold": 100,
            "rates": [
                {"region": "US", "rate": 5.99},
                {"region": "EU", "rate": 9.99}
            ]
        },
        "payment_methods": ["credit_card", "paypal", "apple_pay"],
        "notifications": {
            "email": true,
            "sms": false
        }
     }'::jsonb,
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00', NULL),

    -- Pending store
    ('550e8400-e29b-41d4-a716-446655440002', 'Fashion Forward', 'fashion-forward',
     'Trendy fashion for the modern individual', 'https://example.com/fashion-forward-logo.png',
     'https://example.com/fashion-forward-cover.png', 'pending', false,
     '550e8400-e29b-41d4-a716-446655440003', 'hello@fashionforward.com', '+1987654321',
     '456 Style Avenue, Fashion District',
     '{
        "theme": "light",
        "currency": "USD",
        "tax_settings": {
            "enabled": true,
            "rate": 0.06
        }
     }'::jsonb,
     '2024-02-01 00:00:00+00', '2024-02-01 00:00:00+00', NULL),

    -- Suspended store
    ('550e8400-e29b-41d4-a716-446655440004', 'Home Decor', 'home-decor',
     'Beautiful home decor items', 'https://example.com/home-decor-logo.png',
     'https://example.com/home-decor-cover.png', 'suspended', true,
     '550e8400-e29b-41d4-a716-446655440005', 'support@homedecor.com', '+1122334455',
     '789 Design Street, Art District',
     '{
        "theme": "light",
        "currency": "USD",
        "tax_settings": {
            "enabled": true,
            "rate": 0.07
        }
     }'::jsonb,
     '2024-03-01 00:00:00+00', '2024-03-15 00:00:00+00', NULL),

    -- Closed store (with deleted_at)
    ('550e8400-e29b-41d4-a716-446655440006', 'Book Nook', 'book-nook',
     'Your favorite bookstore', 'https://example.com/book-nook-logo.png',
     'https://example.com/book-nook-cover.png', 'closed', true,
     '550e8400-e29b-41d4-a716-446655440007', 'info@booknook.com', '+1555666777',
     '321 Reading Lane, Book District',
     '{
        "theme": "classic",
        "currency": "USD"
     }'::jsonb,
     '2023-01-01 00:00:00+00', '2023-12-31 00:00:00+00', '2023-12-31 00:00:00+00');

-- Insert sample store categories
INSERT INTO store_categories (
    id, store_id, name, slug, description, parent_id,
    created_at, updated_at
) VALUES
    -- Tech Haven categories
    ('550e8400-e29b-41d4-a716-446655440008', '550e8400-e29b-41d4-a716-446655440000',
     'Laptops', 'laptops', 'All laptop computers', NULL,
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440009', '550e8400-e29b-41d4-a716-446655440000',
     'Gaming Laptops', 'gaming-laptops', 'High-performance gaming laptops',
     '550e8400-e29b-41d4-a716-446655440008',
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440010', '550e8400-e29b-41d4-a716-446655440000',
     'Business Laptops', 'business-laptops', 'Professional laptops for business',
     '550e8400-e29b-41d4-a716-446655440008',
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440011', '550e8400-e29b-41d4-a716-446655440000',
     'Smartphones', 'smartphones', 'Mobile phones and accessories', NULL,
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440012', '550e8400-e29b-41d4-a716-446655440000',
     'Gaming Accessories', 'gaming-accessories', 'Gaming peripherals and accessories', NULL,
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    -- Fashion Forward categories
    ('550e8400-e29b-41d4-a716-446655440013', '550e8400-e29b-41d4-a716-446655440002',
     'Men''s Clothing', 'mens-clothing', 'Clothing for men', NULL,
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440014', '550e8400-e29b-41d4-a716-446655440002',
     'Women''s Clothing', 'womens-clothing', 'Clothing for women', NULL,
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440015', '550e8400-e29b-41d4-a716-446655440002',
     'Formal Wear', 'formal-wear', 'Formal clothing for special occasions', NULL,
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    -- Home Decor categories
    ('550e8400-e29b-41d4-a716-446655440016', '550e8400-e29b-41d4-a716-446655440004',
     'Furniture', 'furniture', 'Home furniture and furnishings', NULL,
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440017', '550e8400-e29b-41d4-a716-446655440004',
     'Decor', 'decor', 'Home decor and accessories', NULL,
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00'),

    ('550e8400-e29b-41d4-a716-446655440018', '550e8400-e29b-41d4-a716-446655440004',
     'Living Room', 'living-room', 'Living room furniture and decor',
     '550e8400-e29b-41d4-a716-446655440016',
     '2024-01-01 00:00:00+00', '2024-01-01 00:00:00+00');