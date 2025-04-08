package config

// ProductConfig holds configuration for product-related operations
type ProductConfig struct {
	DefaultPageSize int `mapstructure:"default_page_size" default:"10"`
	MaxPageSize     int `mapstructure:"max_page_size" default:"100"`
}

// NewProductConfig creates a new ProductConfig with default values
func NewProductConfig() *ProductConfig {
	return &ProductConfig{
		DefaultPageSize: 10,
		MaxPageSize:     100,
	}
}
