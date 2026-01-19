package templateStoreV2

import (
	"context"
	"fmt"

	ce "github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao/entities"
)

func (record *TemplateStore) upgradeProcessing() error {
	return nil
}

func (record *TemplateStore) defaultProcessing() error {
	return nil
}

func (record *TemplateStore) validationProcessing() error {
	return nil
}

func (h *TemplateStore) postGetProcessing() error {
	return nil
}

func (record *TemplateStore) preDeleteProcessing() error {
	return nil
}

func templateClone(ctx context.Context, source TemplateStore) (TemplateStore, error) {
	_ = ctx
	_ = source
	panic("Not Implemented")
}

func assertTemplateStore(result any, field entities.Field, value any) (*TemplateStore, error) {
	x, ok := result.(*TemplateStore)
	if !ok {
		return nil, ce.ErrDAOAssertWrapper(tableName, field.String(), value,
			ce.ErrInvalidTypeWrapper(field.String(), fmt.Sprintf("%T", result), "*TemplateStore"))
	}
	return x, nil
}
