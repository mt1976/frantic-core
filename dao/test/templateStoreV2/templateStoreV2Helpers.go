package templateStoreV2

import (
	"context"
	"fmt"

	ce "github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao/entities"
)

// upgradeProcessing performs any one-time upgrade or migration logic on the record.
func (record *TemplateStore) upgradeProcessing() error {
	return nil
}

// defaultProcessing applies any default values prior to validation and persistence.
func (record *TemplateStore) defaultProcessing() error {
	return nil
}

// validationProcessing validates the record and returns an error if it is invalid.
func (record *TemplateStore) validationProcessing() error {
	return nil
}

// postGetProcessing runs any post-load processing after a record is retrieved.
func (h *TemplateStore) postGetProcessing() error {
	return nil
}

// preDeleteProcessing runs any checks or actions required before delete.
func (record *TemplateStore) preDeleteProcessing() error {
	return nil
}

// templateClone contains the package's clone logic.
func templateClone(ctx context.Context, source TemplateStore) (TemplateStore, error) {
	_ = ctx
	_ = source
	panic("Not Implemented")
}

// assertTemplateStore asserts that an `any` returned by lower layers is a *TemplateStore.
func assertTemplateStore(result any, field entities.Field, value any) (*TemplateStore, error) {
	x, ok := result.(*TemplateStore)
	if !ok {
		return nil, ce.ErrDAOAssertWrapper(tableName, field.String(), value,
			ce.ErrInvalidTypeWrapper(field.String(), fmt.Sprintf("%T", result), "*TemplateStore"))
	}
	return x, nil
}
