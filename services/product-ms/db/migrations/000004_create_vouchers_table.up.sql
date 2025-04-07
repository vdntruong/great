-- Create enum types
CREATE TYPE voucher_type AS ENUM ('percentage', 'fixed_amount', 'free_shipping');
CREATE TYPE voucher_status AS ENUM ('active', 'inactive', 'expired');

-- Create vouchers table
CREATE TABLE vouchers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    store_id UUID NOT NULL REFERENCES stores(id) ON DELETE CASCADE,
    code VARCHAR(50) NOT NULL,
    type voucher_type NOT NULL,
    value DECIMAL(10,2),
    min_purchase_amount DECIMAL(10,2),
    max_discount_amount DECIMAL(10,2),
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE,
    usage_limit INTEGER,
    usage_count INTEGER DEFAULT 0,
    status voucher_status NOT NULL DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(store_id, code)
);

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
CREATE INDEX idx_vouchers_store_id ON vouchers(store_id);
CREATE INDEX idx_vouchers_code ON vouchers(code);
CREATE INDEX idx_vouchers_start_date ON vouchers(start_date);
CREATE INDEX idx_vouchers_end_date ON vouchers(end_date);
CREATE INDEX idx_voucher_products_voucher_id ON voucher_products(voucher_id);
CREATE INDEX idx_voucher_products_product_id ON voucher_products(product_id);
CREATE INDEX idx_voucher_categories_voucher_id ON voucher_categories(voucher_id);
CREATE INDEX idx_voucher_categories_category_id ON voucher_categories(category_id);

-- Create triggers for updated_at
CREATE TRIGGER update_vouchers_updated_at
  BEFORE UPDATE ON vouchers
  FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
