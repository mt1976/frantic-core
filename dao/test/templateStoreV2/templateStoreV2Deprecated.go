package templateStoreV2

// Deprecated wrappers
//
// These exist for templateStore parity. Prefer the field-driven APIs:
// - GetBy(Fields.ID, id)
// - GetBy(Fields.Key, key)
// - DeleteBy(ctx, Fields.Key, key, note)

func GetById(id int) (TemplateStore, error) {
	panic("deprecated")
}

func GetByKey(key string) (TemplateStore, error) {
	panic("deprecated")
}

func DeleteByKey(key any) error {
	panic("deprecated")
}
