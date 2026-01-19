package templateStoreV2

// Deprecated wrappers
//
// These exist for templateStore parity. Prefer the field-driven APIs:
// - GetBy(Fields.ID, id)
// - GetBy(Fields.Key, key)
// - DeleteBy(ctx, Fields.Key, key, note)

// GetById is deprecated; prefer GetBy(Fields.ID, id).
func GetById(id int) (TemplateStore, error) {
	panic("deprecated")
}

// GetByKey is deprecated; prefer GetBy(Fields.Key, key).
func GetByKey(key string) (TemplateStore, error) {
	panic("deprecated")
}

// DeleteByKey is deprecated; prefer DeleteBy(ctx, Fields.Key, key, note).
func DeleteByKey(key any) error {
	panic("deprecated")
}
