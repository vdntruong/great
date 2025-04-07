package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	dao2 "product-ms/db/dao"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/sqlc-dev/pqtype"
	"product-ms/internal/models"
)

// Common type conversions
func stringToFloat64(s string) float64 {
	if s == "" {
		return 0
	}
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func float64ToString(f float64) string {
	return fmt.Sprintf("%.2f", f)
}

func stringToFloat64Ptr(s sql.NullString) *float64 {
	if !s.Valid {
		return nil
	}
	f, _ := strconv.ParseFloat(s.String, 64)
	return &f
}

func float64PtrToString(f *float64) sql.NullString {
	if f == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: fmt.Sprintf("%.2f", *f), Valid: true}
}

func stringToInt32(s string) int32 {
	if s == "" {
		return 0
	}
	i, _ := strconv.ParseInt(s, 10, 32)
	return int32(i)
}

func int32ToString(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

func stringToInt32Ptr(s sql.NullInt32) *int32 {
	if !s.Valid {
		return nil
	}
	return &s.Int32
}

func int32PtrToNullInt32(i *int32) sql.NullInt32 {
	if i == nil {
		return sql.NullInt32{Valid: false}
	}
	return sql.NullInt32{Int32: *i, Valid: true}
}

func nullInt32ToInt32(i sql.NullInt32) int32 {
	if !i.Valid {
		return 0
	}
	return i.Int32
}

func stringToBool(s string) bool {
	return s == "true"
}

func boolToString(b bool) string {
	return strconv.FormatBool(b)
}

func stringToBoolPtr(s sql.NullBool) *bool {
	if !s.Valid {
		return nil
	}
	return &s.Bool
}

func boolPtrToNullBool(b *bool) sql.NullBool {
	if b == nil {
		return sql.NullBool{Valid: false}
	}
	return sql.NullBool{Bool: *b, Valid: true}
}

func timeToNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{Time: t, Valid: true}
}

func timePtrToNullTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{Time: *t, Valid: true}
}

func nullTimeToTimePtr(t sql.NullTime) *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Time
}

func stringToUUID(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

func uuidToString(id uuid.UUID) string {
	return id.String()
}

func jsonToMap(j pqtype.NullRawMessage) (map[string]interface{}, error) {
	if !j.Valid {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(j.RawMessage, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func mapToJSON(m map[string]interface{}) (pqtype.NullRawMessage, error) {
	if m == nil {
		return pqtype.NullRawMessage{Valid: false}, nil
	}
	b, err := json.Marshal(m)
	if err != nil {
		return pqtype.NullRawMessage{Valid: false}, err
	}
	return pqtype.NullRawMessage{RawMessage: b, Valid: true}, nil
}

// Enum type conversions
func storeStatusToDAO(s models.StoreStatus) dao2.StoreStatus {
	return dao2.StoreStatus(s)
}

func daoStoreStatusToModel(s dao2.StoreStatus) models.StoreStatus {
	return models.StoreStatus(s)
}

func productTypeToDAO(t models.ProductType) dao2.ProductType {
	return dao2.ProductType(t)
}

func daoProductTypeToModel(t dao2.ProductType) models.ProductType {
	return models.ProductType(t)
}

func productStatusToDAO(s models.ProductStatus) dao2.ProductStatus {
	return dao2.ProductStatus(s)
}

func daoProductStatusToModel(s dao2.ProductStatus) models.ProductStatus {
	return models.ProductStatus(s)
}

func discountTypeToDAO(t models.DiscountType) dao2.DiscountType {
	return dao2.DiscountType(t)
}

func daoDiscountTypeToModel(t dao2.DiscountType) models.DiscountType {
	return models.DiscountType(t)
}

func discountScopeToDAO(s models.DiscountScope) dao2.DiscountScope {
	return dao2.DiscountScope(s)
}

func daoDiscountScopeToModel(s dao2.DiscountScope) models.DiscountScope {
	return models.DiscountScope(s)
}

func voucherTypeToDAO(t models.VoucherType) dao2.VoucherType {
	return dao2.VoucherType(t)
}

func daoVoucherTypeToModel(t dao2.VoucherType) models.VoucherType {
	return models.VoucherType(t)
}

func voucherStatusToDAO(s models.VoucherStatus) dao2.VoucherStatus {
	return dao2.VoucherStatus(s)
}

func daoVoucherStatusToModel(s dao2.VoucherStatus) models.VoucherStatus {
	return models.VoucherStatus(s)
}

// ConvertDiscountToModel converts a DAO discount to a model discount
func ConvertDiscountToModel(d *dao2.Discount) *models.Discount {
	if d == nil {
		return nil
	}

	return &models.Discount{
		ID:                d.ID,
		StoreID:           d.StoreID,
		Name:              d.Name,
		Code:              d.Code,
		Type:              models.DiscountType(d.Type),
		Value:             stringToFloat64(d.Value),
		Scope:             models.DiscountScope(d.Scope),
		StartDate:         d.StartDate,
		EndDate:           nullTimeToTimePtr(d.EndDate),
		MinPurchaseAmount: stringToFloat64Ptr(d.MinPurchaseAmount),
		MaxDiscountAmount: stringToFloat64Ptr(d.MaxDiscountAmount),
		UsageLimit:        nullInt32ToInt32Ptr(d.UsageLimit),
		UsageCount:        nullInt32ToInt32(d.UsageCount),
		IsActive:          d.IsActive.Bool,
		CreatedAt:         d.CreatedAt.Time,
		UpdatedAt:         d.UpdatedAt.Time,
	}
}

// ConvertDiscountListToModel converts a DAO discount list to a model discount list
func ConvertDiscountListToModel(discounts []*dao2.Discount, total int64, page, limit int) *models.DiscountList {
	items := make([]models.Discount, len(discounts))
	for i, d := range discounts {
		items[i] = *ConvertDiscountToModel(d)
	}

	return &models.DiscountList{
		Discounts:  items,
		TotalCount: total,
		Page:       page,
		Limit:      limit,
	}
}

// ConvertCreateDiscountParamsToDAO converts model create discount params to DAO params
func ConvertCreateDiscountParamsToDAO(params models.CreateDiscountParams) dao2.CreateDiscountParams {
	return dao2.CreateDiscountParams{
		StoreID:           params.StoreID,
		Name:              params.Name,
		Code:              params.Code,
		Type:              dao2.DiscountType(params.Type),
		Value:             float64ToString(params.Value),
		Scope:             dao2.DiscountScope(params.Scope),
		StartDate:         params.StartDate,
		EndDate:           timePtrToNullTime(params.EndDate),
		MinPurchaseAmount: float64PtrToString(params.MinPurchaseAmount),
		MaxDiscountAmount: float64PtrToString(params.MaxDiscountAmount),
		UsageLimit:        int32PtrToNullInt32(params.UsageLimit),
		IsActive:          sql.NullBool{Bool: params.IsActive, Valid: true},
	}
}

// ConvertUpdateDiscountParamsToDAO converts model update discount params to DAO params
func ConvertUpdateDiscountParamsToDAO(params models.UpdateDiscountParams) dao2.UpdateDiscountParams {
	result := dao2.UpdateDiscountParams{
		ID: params.ID,
	}

	if params.Name != nil {
		result.Name = *params.Name
	}
	if params.Code != nil {
		result.Code = *params.Code
	}
	if params.Type != nil {
		result.Type = dao2.DiscountType(*params.Type)
	}
	if params.Value != nil {
		result.Value = float64ToString(*params.Value)
	}
	if params.Scope != nil {
		result.Scope = dao2.DiscountScope(*params.Scope)
	}
	if params.StartDate != nil {
		result.StartDate = *params.StartDate
	}
	if params.EndDate != nil {
		result.EndDate = timePtrToNullTime(params.EndDate)
	}
	if params.MinPurchaseAmount != nil {
		result.MinPurchaseAmount = float64PtrToString(params.MinPurchaseAmount)
	}
	if params.MaxDiscountAmount != nil {
		result.MaxDiscountAmount = float64PtrToString(params.MaxDiscountAmount)
	}
	if params.UsageLimit != nil {
		result.UsageLimit = int32PtrToNullInt32(params.UsageLimit)
	}
	if params.IsActive != nil {
		result.IsActive = boolPtrToNullBool(params.IsActive)
	}

	return result
}

// ConvertListDiscountsParamsToDAO converts model list discount params to DAO params
func ConvertListDiscountsParamsToDAO(params models.ListDiscountsParams) dao2.ListDiscountsParams {
	return dao2.ListDiscountsParams{
		StoreID: params.StoreID,
		Limit:   params.Limit,
		Offset:  params.Offset,
	}
}

// Helper functions for type conversions
func nullStringToStringPtr(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func stringPtrToNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: *s, Valid: true}
}

func nullRawMessageToMap(nrm pqtype.NullRawMessage) map[string]interface{} {
	if !nrm.Valid {
		return nil
	}
	var result map[string]interface{}
	if err := json.Unmarshal(nrm.RawMessage, &result); err != nil {
		return nil
	}
	return result
}

func mapToNullRawMessage(m map[string]interface{}) pqtype.NullRawMessage {
	if m == nil {
		return pqtype.NullRawMessage{Valid: false}
	}
	raw, err := json.Marshal(m)
	if err != nil {
		return pqtype.NullRawMessage{Valid: false}
	}
	return pqtype.NullRawMessage{RawMessage: raw, Valid: true}
}

func nullInt32ToInt32Ptr(ni sql.NullInt32) *int32 {
	if !ni.Valid {
		return nil
	}
	return &ni.Int32
}

func float64PtrToNullString(f *float64) sql.NullString {
	if f == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: fmt.Sprintf("%.2f", *f), Valid: true}
}

func nullBoolToBool(nb sql.NullBool) bool {
	if !nb.Valid {
		return false
	}
	return nb.Bool
}

func boolToNullBool(b bool) sql.NullBool {
	return sql.NullBool{Bool: b, Valid: true}
}

func nullBoolToString(nb sql.NullBool) string {
	if !nb.Valid {
		return "disabled"
	}
	if nb.Bool {
		return "enabled"
	}
	return "disabled"
}

func stringToNullBool(s string) sql.NullBool {
	return sql.NullBool{Bool: s == "enabled", Valid: true}
}

// ConvertStoreToModel converts a DAO store to a model store
func ConvertStoreToModel(s *dao2.Store) *models.Store {
	if s == nil {
		return nil
	}

	return &models.Store{
		ID:           s.ID,
		Name:         s.Name,
		Slug:         s.Slug,
		Description:  nullStringToStringPtr(s.Description),
		LogoURL:      nullStringToStringPtr(s.LogoUrl),
		CoverURL:     nullStringToStringPtr(s.CoverUrl),
		Status:       string(s.Status),
		IsVerified:   s.IsVerified,
		OwnerID:      s.OwnerID,
		ContactEmail: nullStringToStringPtr(s.ContactEmail),
		ContactPhone: nullStringToStringPtr(s.ContactPhone),
		Address:      nullStringToStringPtr(s.Address),
		Settings:     nullRawMessageToMap(s.Settings),
		CreatedAt:    s.CreatedAt.Time,
		UpdatedAt:    s.UpdatedAt.Time,
	}
}

// ConvertStoreListToModel converts a DAO store list to a model store list
func ConvertStoreListToModel(stores []*dao2.Store, total int64, page, limit int) *models.StoreList {
	items := make([]models.Store, len(stores))
	for i, s := range stores {
		items[i] = *ConvertStoreToModel(s)
	}

	return &models.StoreList{
		Stores:     items,
		TotalCount: total,
		Page:       page,
		Limit:      limit,
	}
}

// ConvertCreateStoreParamsToDAO converts model create store params to DAO params
func ConvertCreateStoreParamsToDAO(params models.CreateStoreParams) dao2.CreateStoreParams {
	return dao2.CreateStoreParams{
		Name:         params.Name,
		Slug:         params.Slug,
		Description:  stringPtrToNullString(&params.Description),
		LogoUrl:      stringPtrToNullString(&params.LogoURL),
		CoverUrl:     stringPtrToNullString(&params.CoverURL),
		Status:       dao2.StoreStatus(params.Status),
		IsVerified:   params.IsVerified,
		OwnerID:      params.OwnerID,
		ContactEmail: stringPtrToNullString(&params.ContactEmail),
		ContactPhone: stringPtrToNullString(&params.ContactPhone),
		Address:      stringPtrToNullString(&params.Address),
		Settings:     mapToNullRawMessage(params.Settings),
	}
}

// ConvertUpdateStoreParamsToDAO converts model update store params to DAO params
func ConvertUpdateStoreParamsToDAO(id uuid.UUID, params models.UpdateStoreParams) dao2.UpdateStoreParams {
	result := dao2.UpdateStoreParams{
		ID: id,
	}

	if params.Name != nil {
		result.Name = *params.Name
	}
	if params.Slug != nil {
		result.Slug = *params.Slug
	}
	if params.Description != nil {
		result.Description = stringPtrToNullString(params.Description)
	}
	if params.LogoURL != nil {
		result.LogoUrl = stringPtrToNullString(params.LogoURL)
	}
	if params.CoverURL != nil {
		result.CoverUrl = stringPtrToNullString(params.CoverURL)
	}
	if params.Status != nil {
		result.Status = dao2.StoreStatus(*params.Status)
	}
	if params.IsVerified != nil {
		result.IsVerified = *params.IsVerified
	}
	if params.ContactEmail != nil {
		result.ContactEmail = stringPtrToNullString(params.ContactEmail)
	}
	if params.ContactPhone != nil {
		result.ContactPhone = stringPtrToNullString(params.ContactPhone)
	}
	if params.Address != nil {
		result.Address = stringPtrToNullString(params.Address)
	}
	if params.Settings != nil {
		result.Settings = mapToNullRawMessage(*params.Settings)
	}

	return result
}

// ConvertListStoresParamsToDAO converts model list store params to DAO params
func ConvertListStoresParamsToDAO(params models.ListStoresParams) dao2.ListStoresParams {
	return dao2.ListStoresParams{
		Limit:  params.Limit,
		Offset: params.Offset,
	}
}

// ConvertVoucherToModel converts a DAO voucher to a model voucher
func ConvertVoucherToModel(v *dao2.Voucher) *models.Voucher {
	if v == nil {
		return nil
	}

	return &models.Voucher{
		ID:                v.ID,
		StoreID:           v.StoreID,
		Code:              v.Code,
		Type:              models.VoucherType(v.Type),
		Value:             stringToFloat64Ptr(v.Value),
		MinPurchaseAmount: stringToFloat64Ptr(v.MinPurchaseAmount),
		MaxDiscountAmount: stringToFloat64Ptr(v.MaxDiscountAmount),
		StartDate:         v.StartDate,
		EndDate:           nullTimeToTimePtr(v.EndDate),
		UsageLimit:        nullInt32ToInt32Ptr(v.UsageLimit),
		UsageCount:        nullInt32ToInt32(v.UsageCount),
		Status:            models.VoucherStatus(v.Status),
		CreatedAt:         v.CreatedAt.Time,
		UpdatedAt:         v.UpdatedAt.Time,
	}
}

// ConvertVoucherListToModel converts a DAO voucher list to a model voucher list
func ConvertVoucherListToModel(vouchers []*dao2.Voucher, total int64, page, limit int) *models.VoucherList {
	items := make([]models.Voucher, len(vouchers))
	for i, v := range vouchers {
		items[i] = *ConvertVoucherToModel(v)
	}

	return &models.VoucherList{
		Vouchers:   items,
		TotalCount: total,
		Page:       page,
		Limit:      limit,
	}
}

// ConvertCreateVoucherParamsToDAO converts model create voucher params to DAO params
func ConvertCreateVoucherParamsToDAO(params models.CreateVoucherParams) dao2.CreateVoucherParams {
	return dao2.CreateVoucherParams{
		StoreID:           params.StoreID,
		Code:              params.Code,
		Type:              dao2.VoucherType(params.Type),
		Value:             float64PtrToNullString(params.Value),
		MinPurchaseAmount: float64PtrToNullString(params.MinPurchaseAmount),
		MaxDiscountAmount: float64PtrToNullString(params.MaxDiscountAmount),
		StartDate:         params.StartDate,
		EndDate:           timePtrToNullTime(params.EndDate),
		UsageLimit:        int32PtrToNullInt32(params.UsageLimit),
		Status:            dao2.VoucherStatus(params.Status),
	}
}

// ConvertUpdateVoucherParamsToDAO converts model update voucher params to DAO params
func ConvertUpdateVoucherParamsToDAO(params models.UpdateVoucherParams) dao2.UpdateVoucherParams {
	result := dao2.UpdateVoucherParams{
		ID: params.ID,
	}

	if params.Code != nil {
		result.Code = *params.Code
	}
	if params.Type != nil {
		result.Type = dao2.VoucherType(*params.Type)
	}
	if params.Value != nil {
		result.Value = float64PtrToNullString(params.Value)
	}
	if params.MinPurchaseAmount != nil {
		result.MinPurchaseAmount = float64PtrToNullString(params.MinPurchaseAmount)
	}
	if params.MaxDiscountAmount != nil {
		result.MaxDiscountAmount = float64PtrToNullString(params.MaxDiscountAmount)
	}
	if params.StartDate != nil {
		result.StartDate = *params.StartDate
	}
	if params.EndDate != nil {
		result.EndDate = timePtrToNullTime(params.EndDate)
	}
	if params.UsageLimit != nil {
		result.UsageLimit = int32PtrToNullInt32(params.UsageLimit)
	}
	if params.Status != nil {
		result.Status = dao2.VoucherStatus(*params.Status)
	}

	return result
}

// ConvertListVouchersParamsToDAO converts model list voucher params to DAO params
func ConvertListVouchersParamsToDAO(params models.ListVouchersParams) dao2.ListVouchersParams {
	return dao2.ListVouchersParams{
		StoreID: params.StoreID,
		Limit:   params.Limit,
		Offset:  params.Offset,
	}
}

// ConvertProductToModel converts a DAO product to a model product
func ConvertProductToModel(p *dao2.Product) *models.Product {
	if p == nil {
		return nil
	}

	return &models.Product{
		ID:                p.ID,
		StoreID:           p.StoreID,
		Name:              p.Name,
		Slug:              p.Slug,
		Description:       nullStringToStringPtr(p.Description),
		Type:              string(p.Type),
		Status:            string(p.Status),
		Price:             stringToFloat64(p.Price),
		CompareAtPrice:    stringToFloat64(p.CompareAtPrice.String),
		CostPrice:         stringToFloat64(p.CostPrice.String),
		SKU:               nullStringToStringPtr(p.Sku),
		Barcode:           nullStringToStringPtr(p.Barcode),
		Weight:            stringToFloat64(p.Weight.String),
		WeightUnit:        nullStringToStringPtr(p.WeightUnit),
		IsTaxable:         nullBoolToBool(p.IsTaxable),
		IsFeatured:        nullBoolToBool(p.IsFeatured),
		IsGiftCard:        nullBoolToBool(p.IsGiftCard),
		RequiresShipping:  nullBoolToBool(p.RequiresShipping),
		InventoryPolicy:   nullStringToStringPtr(p.InventoryPolicy),
		InventoryTracking: nullBoolToString(p.InventoryTracking),
		SEOTitle:          nullStringToStringPtr(p.SeoTitle),
		SEODescription:    nullStringToStringPtr(p.SeoDescription),
		Metadata:          nullRawMessageToMap(p.Metadata),
		CreatedAt:         p.CreatedAt.Time,
		UpdatedAt:         p.UpdatedAt.Time,
	}
}

// ConvertCreateProductParamsToDAO converts model create product params to DAO params
func ConvertCreateProductParamsToDAO(params models.CreateProductParams) (*dao2.CreateProductParams, error) {
	return &dao2.CreateProductParams{
		StoreID:           params.StoreID,
		Name:              params.Name,
		Slug:              params.Slug,
		Description:       stringPtrToNullString(&params.Description),
		Type:              dao2.ProductType(params.Type),
		Status:            dao2.ProductStatus(params.Status),
		Price:             float64ToString(params.Price),
		CompareAtPrice:    float64PtrToNullString(&params.CompareAtPrice),
		CostPrice:         float64PtrToNullString(&params.CostPrice),
		Sku:               stringPtrToNullString(&params.SKU),
		Barcode:           stringPtrToNullString(&params.Barcode),
		Weight:            float64PtrToNullString(&params.Weight),
		WeightUnit:        stringPtrToNullString(&params.WeightUnit),
		IsTaxable:         boolToNullBool(params.IsTaxable),
		IsFeatured:        boolToNullBool(params.IsFeatured),
		IsGiftCard:        boolToNullBool(params.IsGiftCard),
		RequiresShipping:  boolToNullBool(params.RequiresShipping),
		InventoryPolicy:   stringPtrToNullString(&params.InventoryPolicy),
		InventoryTracking: stringToNullBool(params.InventoryTracking),
		SeoTitle:          stringPtrToNullString(&params.SEOTitle),
		SeoDescription:    stringPtrToNullString(&params.SEODescription),
		Metadata:          mapToNullRawMessage(params.Metadata),
	}, nil
}

// ConvertUpdateProductParamsToDAO converts model update product params to DAO params
func ConvertUpdateProductParamsToDAO(params models.UpdateProductParams) (*dao2.UpdateProductParams, error) {
	result := &dao2.UpdateProductParams{
		ID: params.ID,
	}

	if params.Name != nil {
		result.Name = *params.Name
	}
	if params.Slug != nil {
		result.Slug = *params.Slug
	}
	if params.Description != nil {
		result.Description = stringPtrToNullString(params.Description)
	}
	if params.Type != nil {
		result.Type = dao2.ProductType(*params.Type)
	}
	if params.Status != nil {
		result.Status = dao2.ProductStatus(*params.Status)
	}
	if params.Price != nil {
		result.Price = float64ToString(*params.Price)
	}
	if params.CompareAtPrice != nil {
		result.CompareAtPrice = float64PtrToNullString(params.CompareAtPrice)
	}
	if params.CostPrice != nil {
		result.CostPrice = float64PtrToNullString(params.CostPrice)
	}
	if params.SKU != nil {
		result.Sku = stringPtrToNullString(params.SKU)
	}
	if params.Barcode != nil {
		result.Barcode = stringPtrToNullString(params.Barcode)
	}
	if params.Weight != nil {
		result.Weight = float64PtrToNullString(params.Weight)
	}
	if params.WeightUnit != nil {
		result.WeightUnit = stringPtrToNullString(params.WeightUnit)
	}
	if params.IsTaxable != nil {
		result.IsTaxable = boolPtrToNullBool(params.IsTaxable)
	}
	if params.IsFeatured != nil {
		result.IsFeatured = boolPtrToNullBool(params.IsFeatured)
	}
	if params.IsGiftCard != nil {
		result.IsGiftCard = boolPtrToNullBool(params.IsGiftCard)
	}
	if params.RequiresShipping != nil {
		result.RequiresShipping = boolPtrToNullBool(params.RequiresShipping)
	}
	if params.InventoryPolicy != nil {
		result.InventoryPolicy = stringPtrToNullString(params.InventoryPolicy)
	}
	if params.InventoryTracking != nil {
		result.InventoryTracking = stringToNullBool(*params.InventoryTracking)
	}
	if params.SEOTitle != nil {
		result.SeoTitle = stringPtrToNullString(params.SEOTitle)
	}
	if params.SEODescription != nil {
		result.SeoDescription = stringPtrToNullString(params.SEODescription)
	}
	if params.Metadata != nil {
		result.Metadata = mapToNullRawMessage(*params.Metadata)
	}

	return result, nil
}

// ConvertListProductsParamsToDAO converts model list product params to DAO params
func ConvertListProductsParamsToDAO(params models.ListProductsParams) *dao2.ListProductsParams {
	return &dao2.ListProductsParams{
		StoreID: params.StoreID,
		Limit:   params.Limit,
		Offset:  params.Offset,
	}
}
