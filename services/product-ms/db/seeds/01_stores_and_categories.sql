-- Insert sample stores
INSERT INTO stores (
    id, name, slug, description, logo_url, cover_url,
    status, is_verified, owner_id, contact_email, contact_phone,
    address, settings
) VALUES
    ('550e8400-e29b-41d4-a716-446655440000', 'Tech Haven', 'tech-haven',
     'Your one-stop shop for all things tech', 'https://example.com/tech-haven-logo.png',
     'https://example.com/tech-haven-cover.png', 'active', true,
     '550e8400-e29b-41d4-a716-446655440001', 'contact@techhaven.com', '+1234567890',
     '123 Tech Street, Silicon Valley', '{"theme": "dark", "currency": "USD"}'),

    ('550e8400-e29b-41d4-a716-446655440002', 'Fashion Forward', 'fashion-forward',
     'Trendy fashion for the modern individual', 'https://example.com/fashion-forward-logo.png',
     'https://example.com/fashion-forward-cover.png', 'active', true,
     '550e8400-e29b-41d4-a716-446655440003', 'hello@fashionforward.com', '+1987654321',
     '456 Style Avenue, Fashion District', '{"theme": "light", "currency": "USD"}');

-- Insert sample store categories
INSERT INTO store_categories (
    id, store_id, name, slug, description, parent_id, sort_order
) VALUES
    -- Tech Haven categories
    ('550e8400-e29b-41d4-a716-446655440004', '550e8400-e29b-41d4-a716-446655440000',
     'Laptops', 'laptops', 'High-performance laptops for work and play', NULL, 1),

    ('550e8400-e29b-41d4-a716-446655440005', '550e8400-e29b-41d4-a716-446655440000',
     'Smartphones', 'smartphones', 'Latest smartphones and accessories', NULL, 2),

    ('550e8400-e29b-41d4-a716-446655440006', '550e8400-e29b-41d4-a716-446655440000',
     'Gaming Laptops', 'gaming-laptops', 'Powerful laptops for gaming enthusiasts',
     '550e8400-e29b-41d4-a716-446655440004', 1),

    -- Fashion Forward categories
    ('550e8400-e29b-41d4-a716-446655440007', '550e8400-e29b-41d4-a716-446655440002',
     'Men''s Clothing', 'mens-clothing', 'Stylish clothing for men', NULL, 1),

    ('550e8400-e29b-41d4-a716-446655440008', '550e8400-e29b-41d4-a716-446655440002',
     'Women''s Clothing', 'womens-clothing', 'Fashionable clothing for women', NULL, 2),

    ('550e8400-e29b-41d4-a716-446655440009', '550e8400-e29b-41d4-a716-446655440002',
     'Accessories', 'accessories', 'Trendy accessories for all', NULL, 3);