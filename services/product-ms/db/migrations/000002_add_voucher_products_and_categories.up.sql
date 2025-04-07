-- Create voucher_products table
CREATE TABLE voucher_products (
    voucher_id UUID NOT NULL REFERENCES vouchers(id) ON DELETE CASCADE,
    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (voucher_id, product_id)
);

-- Create voucher_categories table
CREATE TABLE voucher_categories (
    voucher_id UUID NOT NULL REFERENCES vouchers(id) ON DELETE CASCADE,
    category_id UUID NOT NULL REFERENCES store_categories(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (voucher_id, category_id)
);

-- Create indexes
CREATE INDEX idx_voucher_products_voucher_id ON voucher_products(voucher_id);
CREATE INDEX idx_voucher_products_product_id ON voucher_products(product_id);
CREATE INDEX idx_voucher_categories_voucher_id ON voucher_categories(voucher_id);
CREATE INDEX idx_voucher_categories_category_id ON voucher_categories(category_id);
