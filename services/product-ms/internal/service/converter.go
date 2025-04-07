package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"product-ms/internal/models"
	"product-ms/internal/repository/dao"

	"github.com/google/uuid"
	"github.com/sqlc-dev/pqtype"
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

func nullInt32ToInt32Ptr(i sql.NullInt32) *int32 {
	if !i.Valid {
		return nil
	}
	return &i.Int32
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
func storeStatusToDAO(s models.StoreStatus) dao.StoreStatus {
	return dao.StoreStatus(s)
}

func daoStoreStatusToModel(s dao.StoreStatus) models.StoreStatus {
	return models.StoreStatus(s)
}

func productTypeToDAO(t models.ProductType) dao.ProductType {
	return dao.ProductType(t)
}

func daoProductTypeToModel(t dao.ProductType) models.ProductType {
	return models.ProductType(t)
}

func productStatusToDAO(s models.ProductStatus) dao.ProductStatus {
	return dao.ProductStatus(s)
}

func daoProductStatusToModel(s dao.ProductStatus) models.ProductStatus {
	return models.ProductStatus(s)
}

func discountTypeToDAO(t models.DiscountType) dao.DiscountType {
	return dao.DiscountType(t)
}

func daoDiscountTypeToModel(t dao.DiscountType) models.DiscountType {
	return models.DiscountType(t)
}

func discountScopeToDAO(s models.DiscountScope) dao.DiscountScope {
	return dao.DiscountScope(s)
}

func daoDiscountScopeToModel(s dao.DiscountScope) models.DiscountScope {
	return models.DiscountScope(s)
}

func voucherTypeToDAO(t models.VoucherType) dao.VoucherType {
	return dao.VoucherType(t)
}

func daoVoucherTypeToModel(t dao.VoucherType) models.VoucherType {
	return models.VoucherType(t)
}

func voucherStatusToDAO(s models.VoucherStatus) dao.VoucherStatus {
	return dao.VoucherStatus(s)
}

func daoVoucherStatusToModel(s dao.VoucherStatus) models.VoucherStatus {
	return models.VoucherStatus(s)
}
