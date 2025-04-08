-- Create enum types
CREATE TYPE product_status AS ENUM ('draft', 'active', 'inactive', 'archived');
CREATE TYPE product_type AS ENUM ('simple', 'variable');
CREATE TYPE image_type AS ENUM ('main', 'thumbnail', 'gallery');

-- Create products table
CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    store_id UUID NOT NULL REFERENCES stores(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL,
    description TEXT,
    type product_type NOT NULL DEFAULT 'simple',
    status product_status NOT NULL DEFAULT 'draft',
    price DECIMAL(10,2) NOT NULL,
    compare_at_price DECIMAL(10,2),
    cost_price DECIMAL(10,2),
    sku VARCHAR(100),
    barcode VARCHAR(100),
    weight DECIMAL(10,2),
    weight_unit VARCHAR(10),
    is_taxable BOOLEAN DEFAULT true,
    is_featured BOOLEAN DEFAULT false,
    is_gift_card BOOLEAN DEFAULT false,
    requires_shipping BOOLEAN DEFAULT true,
    inventory_quantity INTEGER DEFAULT 0,
    inventory_policy VARCHAR(20) DEFAULT 'deny',
    inventory_tracking BOOLEAN DEFAULT true,
    seo_title VARCHAR(255),
    seo_description TEXT,
    metadata JSONB DEFAULT '{}'::jsonb,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    UNIQUE(store_id, slug)
);

-- Create product_categories table
CREATE TABLE product_categories (
    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    category_id UUID NOT NULL REFERENCES store_categories(id) ON DELETE CASCADE,
    PRIMARY KEY (product_id, category_id)
);

-- Create product_variants table
CREATE TABLE product_variants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    sku VARCHAR(100),
    barcode VARCHAR(100),
    price DECIMAL(10,2) NOT NULL,
    compare_at_price DECIMAL(10,2),
    cost_price DECIMAL(10,2),
    weight DECIMAL(10,2),
    weight_unit VARCHAR(10),
    inventory_quantity INTEGER DEFAULT 0,
    inventory_policy VARCHAR(20) DEFAULT 'deny',
    inventory_tracking BOOLEAN DEFAULT true,
    option_values JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create product_images table
CREATE TABLE product_images (
      id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
      product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
      variant_id UUID REFERENCES product_variants(id) ON DELETE CASCADE,
      url VARCHAR(255) NOT NULL,
      alt_text VARCHAR(255),
      type image_type NOT NULL DEFAULT 'gallery',
      sort_order INTEGER DEFAULT 0,
      created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes
CREATE INDEX idx_products_store_id ON products(store_id);
CREATE INDEX idx_products_status ON products(status);
CREATE INDEX idx_product_variants_product_id ON product_variants(product_id);
CREATE INDEX idx_product_images_product_id ON product_images(product_id);
CREATE INDEX idx_product_images_variant_id ON product_images(variant_id);

-- Create triggers for updated_at
CREATE TRIGGER update_products_updated_at
    BEFORE UPDATE ON products
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_product_variants_updated_at
    BEFORE UPDATE ON product_variants
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_product_images_updated_at
    BEFORE UPDATE ON product_images
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
