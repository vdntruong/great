// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: product_image.sql

package dao

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const CreateProductImage = `-- name: CreateProductImage :one
INSERT INTO product_images (
    product_id,
    variant_id,
    url,
    alt_text,
    type,
    sort_order
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING id, product_id, variant_id, url, alt_text, type, sort_order, created_at, updated_at
`

type CreateProductImageParams struct {
	ProductID uuid.UUID      `db:"product_id" json:"product_id"`
	VariantID uuid.NullUUID  `db:"variant_id" json:"variant_id"`
	Url       string         `db:"url" json:"url"`
	AltText   sql.NullString `db:"alt_text" json:"alt_text"`
	Type      ImageType      `db:"type" json:"type"`
	SortOrder sql.NullInt32  `db:"sort_order" json:"sort_order"`
}

func (q *Queries) CreateProductImage(ctx context.Context, arg *CreateProductImageParams) (*ProductImage, error) {
	row := q.queryRow(ctx, q.createProductImageStmt, CreateProductImage,
		arg.ProductID,
		arg.VariantID,
		arg.Url,
		arg.AltText,
		arg.Type,
		arg.SortOrder,
	)
	var i ProductImage
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.VariantID,
		&i.Url,
		&i.AltText,
		&i.Type,
		&i.SortOrder,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const DeleteProductImage = `-- name: DeleteProductImage :exec
DELETE FROM product_images
WHERE id = $1
`

func (q *Queries) DeleteProductImage(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteProductImageStmt, DeleteProductImage, id)
	return err
}

const DeleteProductImages = `-- name: DeleteProductImages :exec
DELETE FROM product_images
WHERE product_id = $1
`

func (q *Queries) DeleteProductImages(ctx context.Context, productID uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteProductImagesStmt, DeleteProductImages, productID)
	return err
}

const DeleteVariantImages = `-- name: DeleteVariantImages :exec
DELETE FROM product_images
WHERE variant_id = $1
`

func (q *Queries) DeleteVariantImages(ctx context.Context, variantID uuid.NullUUID) error {
	_, err := q.exec(ctx, q.deleteVariantImagesStmt, DeleteVariantImages, variantID)
	return err
}

const GetGalleryProductImages = `-- name: GetGalleryProductImages :many
SELECT id, product_id, variant_id, url, alt_text, type, sort_order, created_at, updated_at FROM product_images
WHERE product_id = $1 AND type = 'gallery'
ORDER BY sort_order ASC
`

func (q *Queries) GetGalleryProductImages(ctx context.Context, productID uuid.UUID) ([]*ProductImage, error) {
	rows, err := q.query(ctx, q.getGalleryProductImagesStmt, GetGalleryProductImages, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*ProductImage{}
	for rows.Next() {
		var i ProductImage
		if err := rows.Scan(
			&i.ID,
			&i.ProductID,
			&i.VariantID,
			&i.Url,
			&i.AltText,
			&i.Type,
			&i.SortOrder,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const GetMainProductImage = `-- name: GetMainProductImage :one
SELECT id, product_id, variant_id, url, alt_text, type, sort_order, created_at, updated_at FROM product_images
WHERE product_id = $1 AND type = 'main'
LIMIT 1
`

func (q *Queries) GetMainProductImage(ctx context.Context, productID uuid.UUID) (*ProductImage, error) {
	row := q.queryRow(ctx, q.getMainProductImageStmt, GetMainProductImage, productID)
	var i ProductImage
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.VariantID,
		&i.Url,
		&i.AltText,
		&i.Type,
		&i.SortOrder,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const GetProductImage = `-- name: GetProductImage :one
SELECT id, product_id, variant_id, url, alt_text, type, sort_order, created_at, updated_at FROM product_images
WHERE id = $1
`

func (q *Queries) GetProductImage(ctx context.Context, id uuid.UUID) (*ProductImage, error) {
	row := q.queryRow(ctx, q.getProductImageStmt, GetProductImage, id)
	var i ProductImage
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.VariantID,
		&i.Url,
		&i.AltText,
		&i.Type,
		&i.SortOrder,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const GetThumbnailProductImage = `-- name: GetThumbnailProductImage :one
SELECT id, product_id, variant_id, url, alt_text, type, sort_order, created_at, updated_at FROM product_images
WHERE product_id = $1 AND type = 'thumbnail'
LIMIT 1
`

func (q *Queries) GetThumbnailProductImage(ctx context.Context, productID uuid.UUID) (*ProductImage, error) {
	row := q.queryRow(ctx, q.getThumbnailProductImageStmt, GetThumbnailProductImage, productID)
	var i ProductImage
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.VariantID,
		&i.Url,
		&i.AltText,
		&i.Type,
		&i.SortOrder,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const ListProductImages = `-- name: ListProductImages :many
SELECT id, product_id, variant_id, url, alt_text, type, sort_order, created_at, updated_at FROM product_images
WHERE product_id = $1
ORDER BY sort_order ASC, created_at DESC
`

func (q *Queries) ListProductImages(ctx context.Context, productID uuid.UUID) ([]*ProductImage, error) {
	rows, err := q.query(ctx, q.listProductImagesStmt, ListProductImages, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*ProductImage{}
	for rows.Next() {
		var i ProductImage
		if err := rows.Scan(
			&i.ID,
			&i.ProductID,
			&i.VariantID,
			&i.Url,
			&i.AltText,
			&i.Type,
			&i.SortOrder,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const ListVariantImages = `-- name: ListVariantImages :many
SELECT id, product_id, variant_id, url, alt_text, type, sort_order, created_at, updated_at FROM product_images
WHERE variant_id = $1
ORDER BY sort_order ASC, created_at DESC
`

func (q *Queries) ListVariantImages(ctx context.Context, variantID uuid.NullUUID) ([]*ProductImage, error) {
	rows, err := q.query(ctx, q.listVariantImagesStmt, ListVariantImages, variantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*ProductImage{}
	for rows.Next() {
		var i ProductImage
		if err := rows.Scan(
			&i.ID,
			&i.ProductID,
			&i.VariantID,
			&i.Url,
			&i.AltText,
			&i.Type,
			&i.SortOrder,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const UpdateImageSortOrder = `-- name: UpdateImageSortOrder :exec
UPDATE product_images
SET sort_order = $2
WHERE id = $1
`

type UpdateImageSortOrderParams struct {
	ID        uuid.UUID     `db:"id" json:"id"`
	SortOrder sql.NullInt32 `db:"sort_order" json:"sort_order"`
}

func (q *Queries) UpdateImageSortOrder(ctx context.Context, arg *UpdateImageSortOrderParams) error {
	_, err := q.exec(ctx, q.updateImageSortOrderStmt, UpdateImageSortOrder, arg.ID, arg.SortOrder)
	return err
}

const UpdateProductImage = `-- name: UpdateProductImage :one
UPDATE product_images
SET
    url = COALESCE($2, url),
    alt_text = COALESCE($3, alt_text),
    type = COALESCE($4, type),
    sort_order = COALESCE($5, sort_order)
WHERE id = $1
RETURNING id, product_id, variant_id, url, alt_text, type, sort_order, created_at, updated_at
`

type UpdateProductImageParams struct {
	ID        uuid.UUID      `db:"id" json:"id"`
	Url       string         `db:"url" json:"url"`
	AltText   sql.NullString `db:"alt_text" json:"alt_text"`
	Type      ImageType      `db:"type" json:"type"`
	SortOrder sql.NullInt32  `db:"sort_order" json:"sort_order"`
}

func (q *Queries) UpdateProductImage(ctx context.Context, arg *UpdateProductImageParams) (*ProductImage, error) {
	row := q.queryRow(ctx, q.updateProductImageStmt, UpdateProductImage,
		arg.ID,
		arg.Url,
		arg.AltText,
		arg.Type,
		arg.SortOrder,
	)
	var i ProductImage
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.VariantID,
		&i.Url,
		&i.AltText,
		&i.Type,
		&i.SortOrder,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}
