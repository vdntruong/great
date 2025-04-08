-- Insert test users (using fixed UUIDs for consistency)
INSERT INTO carts (id, user_id) VALUES
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'e0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11'),
    ('b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'f0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12');

-- Insert cart items
INSERT INTO cart_items (cart_id, product_id, quantity) VALUES
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 2),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'd0eebc99-9c0b-4ef8-bb6d-6bb9bd380a14', 1),
    ('b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 3);

-- Insert orders with different statuses
INSERT INTO orders (id, user_id, status, total_amount) VALUES
    ('c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 'e0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'pending', 150.00),
    ('d0eebc99-9c0b-4ef8-bb6d-6bb9bd380a14', 'e0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'paid', 75.50),
    ('e0eebc99-9c0b-4ef8-bb6d-6bb9bd380a15', 'f0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'delivered', 200.00);

-- Insert order items
INSERT INTO order_items (order_id, product_id, quantity, price, subtotal) VALUES
    ('c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 'c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 2, 50.00, 100.00),
    ('c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 'd0eebc99-9c0b-4ef8-bb6d-6bb9bd380a14', 1, 50.00, 50.00),
    ('d0eebc99-9c0b-4ef8-bb6d-6bb9bd380a14', 'c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 1, 75.50, 75.50),
    ('e0eebc99-9c0b-4ef8-bb6d-6bb9bd380a15', 'd0eebc99-9c0b-4ef8-bb6d-6bb9bd380a14', 4, 50.00, 200.00);
